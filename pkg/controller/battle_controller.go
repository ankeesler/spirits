package controller

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
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

// BattleReconciler reconciles a Battle object
type BattleReconciler struct {
	client.Client

	ActionSource      ActionSource
	BattleCancelFuncs sync.Map
}

// SetupWithManager sets up the controller with the Manager.
func (r *BattleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&spiritsv1alpha1.Battle{}).
		Owns(&spiritsv1alpha1.Spirit{}).
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
	if cancel, ok := r.BattleCancelFuncs.LoadAndDelete(req.NamespacedName.String()); ok {
		cancel.(context.CancelFunc)()
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
	log.V(2).Info("get in battle spirits", "in battle spirits", inBattleSpirits)

	// Go ahead and create a context for the battle, it will be canceled if
	// not used by the battle
	ctx, cancel := context.WithCancel(context.Background())
	cancelAny, exists := r.BattleCancelFuncs.LoadOrStore(client.ObjectKeyFromObject(battle).String(), cancel)

	// If the battle exists...
	if exists {
		// ...and it is running with the expected spirits, then we are done
		if matchingSpirits(inBattleSpirits, battle.Status.InBattleSpirits) {
			// The context we created is not used by the battle, so trash it
			log.Info("in battle spirits match expected")
			cancel()
			return nil
		}

		// Otherwise, cancel the current battle
		cancelAny.(context.CancelFunc)()
	}

	// Start the battle
	internalBattle, internalInBattleSpirits, err := r.convertToInternalBattle(battle, inBattleSpirits)
	if err != nil {
		return fmt.Errorf("convert to internal battle: %w", err)
	}
	go battlerunner.Run(ctx, internalBattle, internalInBattleSpirits, r.battleCallback)

	// Update the spirits that are running in this battle
	battle.Status.InBattleSpirits = []corev1.LocalObjectReference{}
	for _, inBattleSpirit := range inBattleSpirits {
		battle.Status.InBattleSpirits = append(battle.Status.InBattleSpirits, corev1.LocalObjectReference{
			Name: inBattleSpirit.Name,
		})
	}

	// TODO: does this work with multiple replicas?
	// TODO: what happens if multiple goroutines are running?

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

func (r *BattleReconciler) battleCallback(
	battle *spiritsinternal.Battle,
	inBattleSpirits []*spiritsinternal.Spirit,
	done bool,
	err error,
) {
	ctrl.Log.V(1).Info("battle callback", "battle", battle, "inBattleSpirits", inBattleSpirits, "err", err)

	// Set a really long timeout, just in case
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	// Update the battle
	if err := r.convertAndCreateOrPatch(ctx, battle, &spiritsv1alpha1.Battle{}, func() error {
		battle.Status.Conditions = []metav1.Condition{
			newCondition(battle, progressingCondition, err),
		}

		if err != nil {
			battle.Status.Phase = spiritsinternal.BattlePhaseError
			battle.Status.Message = err.Error()
		} else if done {
			battle.Status.Phase = spiritsinternal.BattlePhaseFinished
		} else {
			battle.Status.Phase = spiritsinternal.BattlePhaseRunning
		}
		battle.Status.ActingSpirit = corev1.LocalObjectReference{}

		return nil
	}); err != nil {
		ctrl.Log.Error(err, "create or patch battle")
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
			spirit.Spec.Actions,
			spirit.Spec.Intelligence,
			r.getLazyActionFunc(battle.DeepCopy(), spirit.DeepCopy()),
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
			return nil, fmt.Errorf("create or patch: %w", err)
		}

		battleName, ok := inBattleSpirit.Labels[inBattleSpiritBattleNameLabel]
		if !ok {
			return nil, errors.New("unknown battle name")
		}

		battleGeneration, ok := inBattleSpirit.Labels[inBattleSpiritBattleGenerationLabel]
		if !ok {
			return nil, errors.New("unknown battle name")
		}

		spiritName, ok := inBattleSpirit.Labels[inBattleSpiritSpiritGenerationLabel]
		if !ok {
			return nil, errors.New("unknown spirit name")
		}

		spiritGeneration, ok := inBattleSpirit.Labels[inBattleSpiritSpiritGenerationLabel]
		if !ok {
			return nil, errors.New("unknown spirit generation")
		}

		actionName, err := r.ActionSource.Pend(
			ctx,
			inBattleSpirit.Namespace,
			battleName,
			battleGeneration,
			spiritName,
			spiritGeneration,
		)
		if err != nil {
			return nil, fmt.Errorf("actions queue pend: %w", err)
		}

		action, err := getAction([]string{actionName}, "", nil)
		if err != nil {
			return nil, fmt.Errorf("get action: %w", err)
		}

		return action, nil
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

func matchingSpirits(spirits []*spiritsv1alpha1.Spirit, spiritRefs []corev1.LocalObjectReference) bool {
	if len(spirits) != len(spiritRefs) {
		return false
	}

	for i := range spirits {
		if spirits[i].Name != spiritRefs[i].Name {
			return false
		}
	}

	return true
}
