package controller

import (
	"context"
	"fmt"
	"sync"
	"time"

	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
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

//+kubebuilder:rbac:groups=ankeesler.github.com,resources=battles,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=ankeesler.github.com,resources=battles/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=ankeesler.github.com,resources=battles/finalizers,verbs=update

//+kubebuilder:rbac:groups=ankeesler.github.com,resources=spirits,verbs=get;list;watch;create;update;patch;delete

// BattleReconciler reconciles a Battle object
type BattleReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// SetupWithManager sets up the controller with the Manager.
func (r *BattleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&spiritsv1alpha1.Battle{}).
		Owns(&spiritsv1alpha1.Spirit{}).
		Complete(&reconciler[*spiritsv1alpha1.Battle, *spiritsinternal.Battle]{
			Client: r.Client,
			Scheme: r.Scheme,
			Handler: &battleHandler{
				Client: r.Client,
				Scheme: r.Scheme,
			},
		})
}

type battleHandler struct {
	client.Client
	Scheme *runtime.Scheme

	Battles sync.Map
}

func (h *battleHandler) NewExternal() *spiritsv1alpha1.Battle { return &spiritsv1alpha1.Battle{} }
func (h *battleHandler) NewInternal() *spiritsinternal.Battle { return &spiritsinternal.Battle{} }

func (h *battleHandler) OnDelete(ctx context.Context, log logr.Logger, req ctrl.Request) error {
	if cancel, ok := h.Battles.LoadAndDelete(req.NamespacedName.String()); ok {
		cancel.(context.CancelFunc)()
	}
	return nil
}

func (h *battleHandler) OnUpsert(
	ctx context.Context,
	log logr.Logger,
	req ctrl.Request,
	battle *spiritsinternal.Battle,
) error {
	// Update conditions on current battle status
	battle.Status.Conditions = []metav1.Condition{
		newCondition(battle, progressingCondition, h.progressBattle(ctx, log, battle)),
	}

	// Force the battle phase to be error, if there is one
	// Otherwise the battle phase will get updated by the battle callback
	if !meta.IsStatusConditionTrue(battle.Status.Conditions, progressingCondition) {
		battle.Status.Phase = spiritsinternal.BattlePhaseError
	}

	return nil
}

func (h *battleHandler) progressBattle(
	ctx context.Context,
	log logr.Logger,
	battle *spiritsinternal.Battle,
) error {
	// Get the spirits that are used in this battle
	inBattleSpirits, err := h.getInBattleSpirits(ctx, log, battle)
	if err != nil {
		return fmt.Errorf("get in battle spirits: %w", err)
	}

	// Go ahead and create a context for the battle, it will be canceled if
	// not used by the battle
	ctx, cancel := context.WithCancel(context.Background())
	cancelAny, exists := h.Battles.LoadOrStore(client.ObjectKeyFromObject(battle).String(), cancel)

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
	go battlerunner.Run(ctx, battle, inBattleSpirits, h.battleCallback)

	// Update the spirits that are running in this battle
	battle.Status.InBattleSpirits = []string{}
	for _, inBattleSpirit := range inBattleSpirits {
		battle.Status.InBattleSpirits = append(battle.Status.InBattleSpirits, inBattleSpirit.Name)
	}

	// TODO: does this work with multiple replicas?
	// TODO: what happens if multiple goroutines are running?

	return nil
}

func (h *battleHandler) getInBattleSpirits(
	ctx context.Context,
	log logr.Logger,
	battle *spiritsinternal.Battle,
) ([]*spiritsinternal.Spirit, error) {
	spirits, err := h.getSpirits(ctx, log, battle)
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
				// TODO: set owner ref
				// OwnerReferences: []metav1.OwnerReference{
				// 	*metav1.NewControllerRef(battle, battle.GroupVersionKind()),
				// },
			},
			Spec: spirit.Spec,
		}
		if err := h.createSpirit(ctx, &inBattleSpirit); err != nil && !k8serrors.IsAlreadyExists(err) {
			return nil, fmt.Errorf("create in battle spirit: %w", err)
		}

		inBattleSpirits = append(inBattleSpirits, &inBattleSpirit)
	}

	return inBattleSpirits, nil
}

func (h *battleHandler) getSpirits(
	ctx context.Context,
	log logr.Logger,
	battle *spiritsinternal.Battle,
) ([]*spiritsinternal.Spirit, error) {
	var spirits []*spiritsinternal.Spirit
	for _, spiritName := range battle.Spec.Spirits {
		spirit := spiritsinternal.Spirit{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: battle.Namespace,
				Name:      spiritName,
			},
		}
		if err := h.getSpirit(ctx, &spirit); err != nil {
			return nil, fmt.Errorf("get spirit: %w", err)
		}

		if !meta.IsStatusConditionTrue(spirit.Status.Conditions, readyCondition) {
			return nil, fmt.Errorf("spirit %s not ready", client.ObjectKeyFromObject(&spirit))
		}

		spirits = append(spirits, &spirit)
	}
	return spirits, nil
}

func (h *battleHandler) battleCallback(
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
	if err := createOrPatch(ctx, h.Client, h.Scheme, battle, &spiritsv1alpha1.Battle{}, func() error {
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
		if err := createOrPatch(ctx, h.Client, h.Scheme, spirit, &spiritsv1alpha1.Spirit{}, func() error {
			spirit.Spec = inBattleSpirit.Spec
			return nil
		}); err != nil {
			log.Log.Error(err, "create or patch spirit")
		}
	}
}

// TODO: we shoudl generic-ify these functions

func (h *battleHandler) getSpirit(ctx context.Context, spirit *spiritsinternal.Spirit) error {
	var externalSpirit spiritsv1alpha1.Spirit
	if err := h.Get(ctx, client.ObjectKeyFromObject(spirit), &externalSpirit); err != nil {
		return fmt.Errorf("get spirit: %w", err)
	}

	if err := h.Scheme.Convert(&externalSpirit, spirit, nil); err != nil {
		return fmt.Errorf("convert external spirit to internal spirit: %w", err)
	}

	// TODO: is this right calling this both places???
	var err error
	spirit.Spec.Internal.Action, err = getAction(spirit.Spec.Actions, spirit.Spec.Intelligence, nil)
	if err != nil {
		return fmt.Errorf("get action: %w", err)
	}

	return nil
}

func (h *battleHandler) createSpirit(ctx context.Context, spirit *spiritsinternal.Spirit) error {
	var externalSpirit spiritsv1alpha1.Spirit
	if err := h.Scheme.Convert(spirit, &externalSpirit, nil); err != nil {
		return fmt.Errorf("convert internal spirit to external spirit: %w", err)
	}

	if err := h.Create(ctx, &externalSpirit); err != nil {
		return fmt.Errorf("update spirit: %w", err)
	}

	return nil
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
