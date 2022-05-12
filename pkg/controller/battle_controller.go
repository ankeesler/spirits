package controller

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
	"github.com/ankeesler/spirits/internal/battlerunner"
	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
	"github.com/go-logr/logr"
)

type battleContext struct {
	context.Context
	cancelFunc  context.CancelFunc
	spiritsRefs []corev1.LocalObjectReference

	cancelReasons []string
}

func newBattleContext(spiritsRefs []corev1.LocalObjectReference) *battleContext {
	ctx, cancel := context.WithCancel(context.Background())
	return &battleContext{
		Context:     ctx,
		cancelFunc:  cancel,
		spiritsRefs: spiritsRefs,
	}
}

func (bc *battleContext) Err() error {
	if bc.Context.Err() != nil {
		return fmt.Errorf("%w (%s)", bc.Context.Err(), strings.Join(bc.cancelReasons, ","))
	}
	return nil
}

func (bc *battleContext) cancel(reason string) {
	bc.cancelReasons = append(bc.cancelReasons, reason)
	bc.cancelFunc()
}

// BattleReconciler reconciles a Battle object
type BattleReconciler struct {
	client.Client

	ActionSource ActionSource
	BattleCache  sync.Map // map from spiritsv1alpha1.Battle types.NamespacedName.String() to *battleContext
}

// SetupWithManager sets up the controller with the Manager.
func (r *BattleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&spiritsv1alpha1.Battle{}).
		Owns(&spiritsv1alpha1.Spirit{}).
		Named("battle").
		Complete(r)
}

func (r *BattleReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var battle spiritsv1alpha1.Battle
	return reconcile(ctx, req, &reconcileHelper[*spiritsv1alpha1.Battle]{
		Client:   r.Client,
		Object:   &battle,
		OnDelete: r.onDelete,
		OnUpsert: r.onUpsert,
	})
}

func (r *BattleReconciler) onDelete(ctx context.Context, log logr.Logger, req ctrl.Request) error {
	if battleContextAny, ok := r.BattleCache.LoadAndDelete(req.NamespacedName.String()); ok {
		battleContextAny.(*battleContext).cancel("battle deleted")
	}
	return nil
}

func (r *BattleReconciler) onUpsert(
	ctx context.Context,
	log logr.Logger,
	req ctrl.Request,
	battle *spiritsv1alpha1.Battle,
) error {
	// Update conditions on current battle status
	err := r.progressBattle(ctx, log, battle)
	battle.Status.Conditions = []metav1.Condition{
		newCondition(battle, progressingCondition, err),
	}

	// Force the battle phase to be error, if there is one
	// Otherwise the battle phase will get updated by the battle callback
	if !meta.IsStatusConditionTrue(battle.Status.Conditions, progressingCondition) {
		log.V(2).Info("setting the phase as error")
		battle.Status.Phase = spiritsv1alpha1.BattlePhaseError
		battle.Status.Message = err.Error()
	}

	return nil
}

func (r *BattleReconciler) progressBattle(
	ctx context.Context,
	log logr.Logger,
	battle *spiritsv1alpha1.Battle,
) error {
	// Get the spirits that are used in this battle
	inBattleSpirits, err := r.getInBattleSpirits(ctx, log, battle)
	if err != nil {
		return fmt.Errorf("get in battle spirits: %w", err)
	}
	log.V(2).Info("get in battle spirits", "battle", battle, "inBattleSpirits", inBattleSpirits)

	// Go ahead and create a new context for this battle, it will be canceled if
	// not used below
	doStartNewBattle := false
	maybeNewBattleContext := newBattleContext(getSpiritsRefs(inBattleSpirits))
	defer func() {
		if !doStartNewBattle {
			maybeNewBattleContext.cancel("context not needed")
		}
	}()

	// Try to get the internal battle for this external battle
	currentBattleContextAny, battleExists := r.BattleCache.LoadOrStore(client.ObjectKeyFromObject(battle).String(), maybeNewBattleContext)
	currentBattleContext := currentBattleContextAny.(*battleContext)

	// If no battle exists, we will need to start a new battle
	if !battleExists {
		log.V(1).Info("starting new battle: one does not exist")
		doStartNewBattle = true
	}

	// If there is an existing battle, and the existing battle is running with different spirits than we expect,
	// then we cancel the old battle context and use the new one
	if battleExists && !equality.Semantic.DeepEqual(maybeNewBattleContext.spiritsRefs, currentBattleContext.spiritsRefs) {
		log.V(1).Info(
			"starting new battle: desired in battle spirits do not match actual spirits",
			"desired", maybeNewBattleContext.spiritsRefs,
			"actual", currentBattleContext.spiritsRefs,
		)

		currentBattleContext.cancel(fmt.Sprintf(
			"desired in battle spirits %s do not match actual in battle spirits %s",
			maybeNewBattleContext.spiritsRefs,
			currentBattleContext.spiritsRefs,
		))
		currentBattleContext = maybeNewBattleContext
		r.BattleCache.Store(client.ObjectKeyFromObject(battle).String(), currentBattleContext)

		doStartNewBattle = true
	}

	// Really start the battle, if we need to
	if doStartNewBattle {
		internalBattle, internalInBattleSpirits, err := r.convertToInternalBattle(battle, inBattleSpirits)
		if err != nil {
			return fmt.Errorf("convert to internal battle: %w", err)
		}
		go r.runBattle(currentBattleContext, internalBattle, internalInBattleSpirits)
	}

	// Update the external battle in-battle spirits
	// Other battle status fields are updated during the battle callback and the lazy action func
	battle.Status.InBattleSpirits = currentBattleContext.spiritsRefs

	return nil
}

func (r *BattleReconciler) getInBattleSpirits(
	ctx context.Context,
	log logr.Logger,
	battle *spiritsv1alpha1.Battle,
) ([]*spiritsv1alpha1.Spirit, error) {
	spirits, err := r.getSpirits(ctx, log, battle)
	if err != nil {
		return nil, fmt.Errorf("get spirits: %w", err)
	}
	log.V(2).Info("get spirits", "spirits", spirits)

	var inBattleSpirits []*spiritsv1alpha1.Spirit
	for _, spirit := range spirits {
		// Get in-battle external spirit
		inBattleSpirit := spiritsv1alpha1.Spirit{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: battle.Namespace,
				Name:      fmt.Sprintf("%s-%d-%s-%d", battle.Name, battle.Generation, spirit.Name, spirit.Generation),
				Labels: map[string]string{
					inBattleSpiritBattleNameLabel:       battle.Name,
					inBattleSpiritBattleGenerationLabel: fmt.Sprintf("%d", battle.Generation),
					inBattleSpiritSpiritNameLabel:       spirit.Name,
					inBattleSpiritSpiritGenerationLabel: fmt.Sprintf("%d", spirit.Generation),
				},
				OwnerReferences: []metav1.OwnerReference{
					*metav1.NewControllerRef(battle, battle.GroupVersionKind()),
				},
			},
			Spec: spirit.Spec,
		}
		if err := r.Create(ctx, &inBattleSpirit); err != nil && !k8serrors.IsAlreadyExists(err) {
			return nil, fmt.Errorf("create in battle spirit: %w", err)
		}

		inBattleSpirits = append(inBattleSpirits, &inBattleSpirit)
	}

	return inBattleSpirits, nil
}

func (r *BattleReconciler) getSpirits(
	ctx context.Context,
	log logr.Logger,
	battle *spiritsv1alpha1.Battle,
) ([]*spiritsv1alpha1.Spirit, error) {
	var spirits []*spiritsv1alpha1.Spirit
	for _, spiritRef := range battle.Spec.Spirits {
		spirit := spiritsv1alpha1.Spirit{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: battle.Namespace,
				Name:      spiritRef.Name,
			},
		}
		if err := r.Get(ctx, client.ObjectKeyFromObject(&spirit), &spirit); err != nil {
			return nil, fmt.Errorf("get spirit: %w", err)
		}
		log.V(2).Info("get spirit", "spirit", spirit, "name", spiritRef)

		if !meta.IsStatusConditionTrue(spirit.Status.Conditions, readyCondition) {
			return nil, fmt.Errorf("spirit %s not ready", client.ObjectKeyFromObject(&spirit))
		}

		spirits = append(spirits, &spirit)
	}
	return spirits, nil
}

func (r *BattleReconciler) runBattle(
	ctx context.Context,
	battle *spiritsinternal.Battle,
	inBattleSpirits []*spiritsinternal.Spirit,
) {
	ctrl.Log.V(1).Info("battle starting", "battle", battle, "inBattleSpirits", spiritsString(inBattleSpirits))

	// Run the battle
	battlerunner.Run(ctx, battle, inBattleSpirits, r.battleCallback)

	ctrl.Log.V(1).Info("battle finished", "battle", battle, "inBattleSpirits", spiritsString(inBattleSpirits))

	// After the battle is over, update the status
	// Don't clear it from the cache, because we will want to remember that we don't have to run this battle again
	if err := r.convertAndCreateOrPatch(ctx, battle, &spiritsv1alpha1.Battle{}, func() error {
		battle.Status.Phase = spiritsinternal.BattlePhaseFinished
		return nil
	}); err != nil {
		ctrl.Log.Error(err, "run battle: convert and create or patch")
	}
	if battleContextAny, ok := r.BattleCache.Load(client.ObjectKeyFromObject(battle).String()); ok {
		battleContextAny.(*battleContext).cancel("battle finished")
	}
}

func (r *BattleReconciler) battleCallback(
	battle *spiritsinternal.Battle,
	inBattleSpirits []*spiritsinternal.Spirit,
	err error,
) {
	ctrl.Log.V(1).Info("battle callback", "battle", battle, "inBattleSpirits", spiritsString(inBattleSpirits), "err", err)

	// Set a really long timeout, just in case
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	// Update the battle
	if createOrPatchErr := r.convertAndCreateOrPatch(ctx, battle, &spiritsv1alpha1.Battle{}, func() error {
		battle.Status.Conditions = []metav1.Condition{
			newCondition(battle, progressingCondition, err),
		}

		battle.Status.Phase = spiritsinternal.BattlePhaseRunning
		battle.Status.ActingSpirit = corev1.LocalObjectReference{}
		battle.Status.Message = ""
		if err != nil {
			battle.Status.Phase = spiritsinternal.BattlePhaseError
			battle.Status.Message = err.Error()
		}

		return nil
	}); createOrPatchErr != nil {
		ctrl.Log.Error(createOrPatchErr, "create or patch battle")
	}

	// Update the spirits
	for _, inBattleSpirit := range inBattleSpirits {
		internalSpirit := inBattleSpirit.DeepCopy()
		if err := r.convertAndCreateOrPatch(ctx, internalSpirit, &spiritsv1alpha1.Spirit{}, func() error {
			internalSpirit.Spec = inBattleSpirit.Spec
			return nil
		}); err != nil {
			ctrl.Log.Error(err, "create or patch spirit")
		}
	}
}

func (r *BattleReconciler) convertToInternalBattle(
	battle *spiritsv1alpha1.Battle,
	spirits []*spiritsv1alpha1.Spirit,
) (*spiritsinternal.Battle, []*spiritsinternal.Spirit, error) {
	var internalBattle spiritsinternal.Battle
	if err := r.Client.Scheme().Convert(battle, &internalBattle, nil); err != nil {
		return nil, nil, fmt.Errorf("convert external battle to internal battle: %w", err)
	}
	ctrl.Log.V(2).Info("convert external battle to internal battle", "external battle", battle, "internal battle", internalBattle)

	internalSpirits := []*spiritsinternal.Spirit{}
	for _, spirit := range spirits {
		var internalSpirit spiritsinternal.Spirit
		if err := r.Client.Scheme().Convert(spirit, &internalSpirit, nil); err != nil {
			return nil, nil, fmt.Errorf("convert external spirit to internal spirit: %w", err)
		}
		ctrl.Log.V(2).Info("convert external spirit to internal spirit", "external spirit", spirit, "internal battle", internalSpirit)

		var err error
		internalSpirit.Spec.Internal.Action, err = getAction(
			&spirit.Spec.Action,
			r.getLazyActionFunc(battle, spirit),
			r.Client.Scheme(),
		)
		if err != nil {
			return nil, nil, fmt.Errorf("get action: %w", err)
		}

		internalSpirits = append(internalSpirits, &internalSpirit)
	}

	return &internalBattle, internalSpirits, nil
}

func (r *BattleReconciler) getLazyActionFunc(
	battle *spiritsv1alpha1.Battle,
	inBattleSpirit *spiritsv1alpha1.Spirit,
) func(ctx context.Context) (spiritsinternal.Action, error) {
	return func(ctx context.Context) (spiritsinternal.Action, error) {
		ctrl.Log.V(1).Info("lazy action func", "battle", battle, "inBattleSpirit", inBattleSpirit)

		if _, err := controllerutil.CreateOrPatch(ctx, r.Client, battle, func() error {
			battle.Status.ActingSpirit = corev1.LocalObjectReference{Name: inBattleSpirit.Name}
			battle.Status.Phase = spiritsv1alpha1.BattlePhaseAwaitingAction
			return nil
		}); err != nil {
			return nil, fmt.Errorf("create or patch battle during lazy action: %w", err)
		}

		battleName, ok := inBattleSpirit.Labels[inBattleSpiritBattleNameLabel]
		if !ok {
			return nil, fmt.Errorf("missing battle name from spirit %q", client.ObjectKeyFromObject(inBattleSpirit).String())
		}

		spiritName, ok := inBattleSpirit.Labels[inBattleSpiritSpiritNameLabel]
		if !ok {
			return nil, fmt.Errorf("missing spirit name from spirit %q", client.ObjectKeyFromObject(inBattleSpirit).String())
		}

		if battle.Name != battleName {
			return nil, fmt.Errorf("battle name from battle %q does not battle name from spirit %q", battle.Name, battleName)
		}

		actionName, err := r.ActionSource.Pend(
			ctx,
			battle.Namespace,
			battleName,
			spiritName,
		)
		if err != nil {
			return nil, fmt.Errorf("actions queue pend: %w", err)
		}

		// TODO: this would be easier if we had a map instead of a list here...use internal type? And move to internal go package?
		var action *spiritsv1alpha1.SpiritAction
		for _, namedAction := range inBattleSpirit.Spec.Action.Choices.Actions {
			if namedAction.Name == actionName {
				action = &namedAction.Action
				break
			}
		}
		if action == nil {
			return nil, fmt.Errorf("actions queue pend: unknown action for name: %q", actionName)
		}

		internalAction, err := getAction(
			action,
			r.getLazyActionFunc(battle, inBattleSpirit),
			r.Client.Scheme(),
		)
		if err != nil {
			return nil, fmt.Errorf("get action: %w", err)
		}

		return internalAction, nil
	}
}

func (r *BattleReconciler) convertAndCreateOrPatch(
	ctx context.Context,
	internalObj, externalObj client.Object,
	mutateFunc func() error,
) error {
	externalObj.SetNamespace(internalObj.GetNamespace())
	externalObj.SetName(internalObj.GetName())
	externalObj.SetLabels(internalObj.GetLabels())
	externalObj.SetOwnerReferences(internalObj.GetOwnerReferences())
	if _, err := controllerutil.CreateOrPatch(ctx, r.Client, externalObj, func() error {
		ctrl.Log.V(2).Info("convert and create or patch: pre-external-to-internal-convert", "internal object", internalObj, "external object", externalObj)

		if err := r.Client.Scheme().Convert(externalObj, internalObj, nil); err != nil {
			return fmt.Errorf("convert external object to internal object: %w", err)
		}

		ctrl.Log.V(2).Info("convert and create or patch: pre-mutate", "internal object", internalObj, "external object", externalObj)

		if err := mutateFunc(); err != nil {
			return err
		}

		ctrl.Log.V(2).Info("convert and create or patch: post-mutate", "internal object", internalObj, "external object", externalObj)

		if err := r.Client.Scheme().Convert(internalObj, externalObj, nil); err != nil {
			return fmt.Errorf("convert internal object to external object: %w", err)
		}

		ctrl.Log.V(2).Info("convert and create or patch: post-internal-to-external-convert", "internal object", internalObj, "external object", externalObj)

		return nil
	}); err != nil {
		return err
	}
	return nil
}

func getSpiritsRefs(spirits []*spiritsv1alpha1.Spirit) []corev1.LocalObjectReference {
	var refs []corev1.LocalObjectReference
	for _, spirit := range spirits {
		refs = append(refs, corev1.LocalObjectReference{Name: spirit.Name})
	}
	return refs
}

func spiritsString(spirits []*spiritsinternal.Spirit) string {
	s := strings.Builder{}
	for _, spirit := range spirits {
		s.WriteString(fmt.Sprintf("%s@%d ", spirit.Name, spirit.Spec.Attributes.Stats.Health))
	}
	return s.String()
}
