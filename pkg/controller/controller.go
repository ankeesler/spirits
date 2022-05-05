package controller

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//+kubebuilder:rbac:groups=spirits.ankeesler.github.io,resources=spirits,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=spirits.ankeesler.github.io,resources=spirits/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=spirits.ankeesler.github.io,resources=spirits/finalizers,verbs=update

//+kubebuilder:rbac:groups=spirits.ankeesler.github.io,resources=battles,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=spirits.ankeesler.github.io,resources=battles/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=spirits.ankeesler.github.io,resources=battles/finalizers,verbs=update

const (
	inBattleSpiritBattleNameLabel       = "spirits.ankeesler.github.io/battle-name"
	inBattleSpiritBattleGenerationLabel = "spirits.ankeesler.github.io/battle-generation"
	inBattleSpiritSpiritNameLabel       = "spirits.ankeesler.github.io/spirit-name"
	inBattleSpiritSpiritGenerationLabel = "spirits.ankeesler.github.io/spirit-generation"

	readyCondition       = "Ready"
	progressingCondition = "Progressing"
)

func newCondition(
	obj metav1.Object,
	teyep string,
	err error,
) metav1.Condition {
	condition := metav1.Condition{
		Type:               teyep,
		Status:             metav1.ConditionTrue,
		Reason:             "Success",
		Message:            "success",
		ObservedGeneration: obj.GetGeneration(),
		LastTransitionTime: metav1.NewTime(time.Now()),
	}
	if err != nil {
		condition.Status = metav1.ConditionFalse
		condition.Reason = "Error"
		condition.Message = err.Error()
	}
	return condition
}
