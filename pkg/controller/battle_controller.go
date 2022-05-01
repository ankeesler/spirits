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
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
	"github.com/ankeesler/spirits/internal/battlerunner"
	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
	"github.com/go-logr/logr"
)

//+kubebuilder:rbac:groups=ankeesler.github.com,resources=battles,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=ankeesler.github.com,resources=battles/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=ankeesler.github.com,resources=battles/finalizers,verbs=update

//+kubebuilder:rbac:groups=ankeesler.github.com,resources=spirits,verbs=get;list;watch;create;update;patch;delete

// BattleReconciler reconciles a Battle object
type BattleReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	ActionSource     ActionSource
	BattleCancelFuns sync.Map
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
	if cancel, ok := r.BattleCancelFuns.LoadAndDelete(req.NamespacedName.String()); ok {
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
	battle.Status.Conditions = []metav1.Condition{
		newCondition(battle, progressingCondition, r.progressBattle(ctx, log, battle)),
	}

	// Force the battle phase to be error, if there is one
	// Otherwise the battle phase will get updated by the battle callback
	if !meta.IsStatusConditionTrue(battle.Status.Conditions, progressingCondition) {
		battle.Status.Phase = spiritsv1alpha1.BattlePhaseError
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

	// Go ahead and create a context for the battle, it will be canceled if
	// not used by the battle
	ctx, cancel := context.WithCancel(context.Background())
	cancelAny, exists := r.BattleCancelFuns.LoadOrStore(client.ObjectKeyFromObject(battle).String(), cancel)

	// If the battle exists...
	if exists {
		// ...and it is running with the expected spirits, then we are done
		if matchingSpirits(inBattleSpirits, battle.Status.InBattleSpirits) {
			// The context we created is not used by the battle, so trash it
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

	var inBattleSpirits []*spiritsv1alpha1.Spirit
	for _, spirit := range spirits {
		// Get in-battle external spirit
		inBattleSpirit := spiritsv1alpha1.Spirit{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: battle.Namespace,
				Name:      fmt.Sprintf("%s-%s-%d", battle.Name, spirit.Name, spirit.Generation),
				Labels: map[string]string{
					inBattleSpiritBattleNameLabel:       battle.Name,
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
	for _, spiritName := range battle.Spec.Spirits {
		spirit := spiritsv1alpha1.Spirit{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: battle.Namespace,
				Name:      spiritName.Name,
			},
		}
		if err := r.Get(ctx, client.ObjectKeyFromObject(&spirit), &spirit); err != nil {
			return nil, fmt.Errorf("get spirit: %w", err)
		}

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
	log.Log.V(1).Info("battle callback", "battle", battle, "spirits", inBattleSpirits, "err", err)

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

		return nil
	}); err != nil {
		log.Log.Error(err, "create or patch battle")
	}

	// Update the spirits
	for _, inBattleSpirit := range inBattleSpirits {
		spirit := inBattleSpirit.DeepCopy()
		if err := r.convertAndCreateOrPatch(ctx, spirit, &spiritsv1alpha1.Spirit{}, func() error {
			spirit.Spec = inBattleSpirit.Spec
			return nil
		}); err != nil {
			log.Log.Error(err, "create or patch spirit")
		}
	}
}

func (r *BattleReconciler) convertToInternalBattle(
	battle *spiritsv1alpha1.Battle,
	spirits []*spiritsv1alpha1.Spirit,
) (*spiritsinternal.Battle, []*spiritsinternal.Spirit, error) {
	var internalBattle spiritsinternal.Battle
	if err := r.Scheme.Convert(battle, &internalBattle, nil); err != nil {
		return nil, nil, fmt.Errorf("convert external battle to internal battle: %w", err)
	}

	internalSpirits := []*spiritsinternal.Spirit{}
	for _, spirit := range spirits {
		var internalSpirit spiritsinternal.Spirit
		if err := r.Scheme.Convert(&internalSpirit, spirit, nil); err != nil {
			return nil, nil, fmt.Errorf("convert external spirit to internal spirit: %w", err)
		}

		var err error
		internalSpirit.Spec.Internal.Action, err = getAction(
			spirit.Spec.Actions,
			spirit.Spec.Intelligence,
			r.getLazyActionFunc(battle, spirit),
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
	spirit *spiritsv1alpha1.Spirit,
) func(ctx context.Context) (spiritsinternal.Action, error) {
	return func(ctx context.Context) (spiritsinternal.Action, error) {
		battleName, ok := spirit.Labels[inBattleSpiritBattleNameLabel]
		if !ok {
			return nil, errors.New("unknown battle name")
		}

		battleGeneration, ok := spirit.Labels[inBattleSpiritBattleGenerationLabel]
		if !ok {
			return nil, errors.New("unknown battle name")
		}

		spiritName, ok := spirit.Labels[inBattleSpiritSpiritGenerationLabel]
		if !ok {
			return nil, errors.New("unknown spirit name")
		}

		spiritGeneration, ok := spirit.Labels[inBattleSpiritSpiritGenerationLabel]
		if !ok {
			return nil, errors.New("unknown spirit generation")
		}

		actionName, err := r.ActionSource.Pend(
			ctx,
			spirit.Namespace,
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
		if err := r.Scheme.Convert(externalObj, internalObj, nil); err != nil {
			return fmt.Errorf("convert external object to internal object: %w", err)
		}

		if err := mutateFunc(); err != nil {
			return err
		}

		if err := r.Scheme.Convert(internalObj, externalObj, nil); err != nil {
			return fmt.Errorf("convert internal object to external object: %w", err)
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}

func matchingSpirits(spirits []*spiritsv1alpha1.Spirit, spiritNames []corev1.LocalObjectReference) bool {
	if len(spirits) != len(spiritNames) {
		return false
	}

	for i := range spirits {
		if spirits[i].Name != spiritNames[i].Name {
			return false
		}
	}

	return true
}
