package controllers

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	spiritsdevv1alpha1 "github.com/ankeesler/spirits/api/v1alpha1"
)

func newCondition(obj spiritsdevv1alpha1.Object, teyep string, err error) metav1.Condition {
	condition := metav1.Condition{
		Type:               teyep,
		Status:             metav1.ConditionTrue,
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

func getPhase(obj spiritsdevv1alpha1.Object) spiritsdevv1alpha1.Phase {
	for _, condition := range *obj.Conditions() {
		if condition.Status == metav1.ConditionFalse {
			return spiritsdevv1alpha1.PhaseError
		}
	}
	return spiritsdevv1alpha1.PhaseReady
}
