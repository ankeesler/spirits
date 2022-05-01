package controller

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	inBattleSpiritBattleNameLabel       = "spirits.ankeesler.github.com/battle-name"
	inBattleSpiritBattleGenerationLabel = "spirits.ankeesler.github.com/battle-generation"
	inBattleSpiritSpiritNameLabel       = "spirits.ankeesler.github.com/spirit-name"
	inBattleSpiritSpiritGenerationLabel = "spirits.ankeesler.github.com/spirit-generation"

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
