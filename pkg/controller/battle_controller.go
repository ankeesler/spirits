package controller

import (
	"context"
	"fmt"
	"sync"
	"time"

	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
	"github.com/ankeesler/spirits/internal/battlerunner"
	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
	"github.com/go-logr/logr"
)

// BattleReconciler reconciles a Battle object
type BattleReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	Battles sync.Map
}

// SetupWithManager sets up the controller with the Manager.
func (r *BattleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&spiritsv1alpha1.Battle{}).
		Owns(&spiritsv1alpha1.Spirit{}).
		Complete(r)
}

//+kubebuilder:rbac:groups=ankeesler.github.com,resources=battles,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=ankeesler.github.com,resources=battles/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=ankeesler.github.com,resources=battles/finalizers,verbs=update

//+kubebuilder:rbac:groups=ankeesler.github.com,resources=spirits,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *BattleReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// Get battle - if it doesn't exist, then: stop it, delete it from the cache, and return.
	var externalBattle spiritsv1alpha1.Battle
	if err := r.Get(ctx, req.NamespacedName, &externalBattle); err != nil {
		if k8serrors.IsNotFound(err) {
			if cancel, ok := r.Battles.LoadAndDelete(req.NamespacedName.String()); ok {
				cancel.(context.CancelFunc)()
			}
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, fmt.Errorf("get battle: %w", err)
	}

	var battle spiritsinternal.Battle
	if err := r.Scheme.Convert(&externalBattle, &battle, nil); err != nil {
		return ctrl.Result{}, fmt.Errorf("convert: %w", err)
	}

	battlePatch := client.MergeFrom(battle.DeepCopyObject().(client.Object))

	// Update conditions on current battle status
	battle.Status.Conditions = []metav1.Condition{
		newCondition(&battle, "Progress", r.progressBattle(ctx, log, &battle)),
	}

	if err := r.Patch(ctx, &battle, battlePatch); err != nil {
		return ctrl.Result{}, fmt.Errorf("patch battle: %w", err)
	}

	log.Info("reconciled battle")

	return ctrl.Result{}, nil
}

func (r *BattleReconciler) progressBattle(
	ctx context.Context,
	log logr.Logger,
	battle *spiritsinternal.Battle,
) error {
	// Get the spirits that are used in this battle
	inBattleSpirits, err := r.getInBattleSpirits(ctx, log, battle)
	if err != nil {
		return fmt.Errorf("get in battle spirits: %w", err)
	}

	// Go ahead and create a context for the battle, it will be canceled if
	// not used by the battle
	ctx, cancel := context.WithCancel(context.Background())
	cancelAny, exists := r.Battles.LoadOrStore(battle.Name, cancel)

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
	go battlerunner.Run(ctx, battle, inBattleSpirits, r.battleCallback)

	// TODO: what happens

	return nil
}

func (r *BattleReconciler) getInBattleSpirits(
	ctx context.Context,
	log logr.Logger,
	battle *spiritsinternal.Battle,
) ([]*spiritsinternal.Spirit, error) {
	spirits, err := r.getSpirits(ctx, log, battle)
	if err != nil {
		return nil, fmt.Errorf("get spirits: %w", err)
	}

	var inBattleSpirits []*spiritsinternal.Spirit
	for _, spirit := range spirits {
		inBattleSpirit := spiritsinternal.Spirit{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: battle.Namespace,
				Name:      fmt.Sprintf("%s-%s-%d", battle.Name, spirit.Name, spirit.Generation),
				Labels: map[string]string{
					inBattleSpiritBattleNameLabel:       battle.Name,
					inBattleSpiritSpiritNameLabel:       spirit.Name,
					inBattleSpiritSpiritGenerationLabel: fmt.Sprintf("%d", spirit.Generation),
				},
			},
			Spec: spirit.Spec,
		}
		if err := r.Create(ctx, &inBattleSpirit); err != nil && !k8serrors.IsAlreadyExists(err) {
			return nil, fmt.Errorf("create in battle spirit: %w", err)
		}
	}

	return inBattleSpirits, nil
}

func (r *BattleReconciler) getSpirits(
	ctx context.Context,
	log logr.Logger,
	battle *spiritsinternal.Battle,
) ([]*spiritsinternal.Spirit, error) {
	var spirits []*spiritsinternal.Spirit
	for _, spiritName := range battle.Spec.Spirits {
		externalSpirit := spiritsv1alpha1.Spirit{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: battle.Namespace,
				Name:      spiritName,
			},
		}
		if err := r.Get(ctx, client.ObjectKeyFromObject(&externalSpirit), &externalSpirit); err != nil {
			return nil, fmt.Errorf("get spirit: %w", err)
		}

		var spirit spiritsinternal.Spirit
		if err := r.Scheme.Convert(&externalSpirit, &spirit, nil); err != nil {
			return nil, fmt.Errorf("convert: %w", err)
		}
	}
	return spirits, nil
}

func (r *BattleReconciler) battleCallback(
	battle *spiritsinternal.Battle,
	spirits []*spiritsinternal.Spirit,
	err error,
) {
	// Set a really long timeout, just in case
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	if err != nil {
		battle.Status.Phase = spiritsinternal.BattlePhaseError
		battle.Status.Message = err.Error()
	} else {
		battle.Status.Phase = spiritsinternal.BattlePhaseRunning
	}

	if err := r.Update(ctx, battle); err != nil {
		log.Log.Error(err, "update battle")
	}
	for _, spirit := range spirits {
		if err := r.Update(ctx, spirit); err != nil {
			log.Log.Error(err, "update spirit")
		}
	}
}

func matchingSpirits(spirits []*spiritsinternal.Spirit, names []string) bool {
	if len(spirits) != len(names) {
		return false
	}

	for i := range spirits {
		if spirits[i].Name != names[i] {
			return false
		}
	}

	return true
}
