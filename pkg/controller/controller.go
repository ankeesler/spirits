package controller

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	spiritapi "github.com/ankeesler/spirits/pkg/apis/spirits"
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

func getPhase(conditions []metav1.Condition) spiritapi.Phase {
	for i := range conditions {
		if conditions[i].Status == metav1.ConditionFalse {
			return spiritapi.PhaseError
		}
	}
	return spiritapi.PhaseReady
}
