package apiserver

import (
	"context"
	"fmt"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	genericrequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/utils/trace"

	inputinternal "github.com/ankeesler/spirits/internal/apis/spirits/input"
)

type actionRequestHandler struct {
	ActionSink ActionSink
}

var _ interface {
	// Need this so that we can act as an apiserver's storage handler
	rest.Storage
	// Need this so that we get called on "create" verb
	rest.Creater

	// Not sure why/if we need these (2 of them are the same)...
	rest.NamespaceScopedStrategy
	rest.Scoper
} = (*actionRequestHandler)(nil)

func (h *actionRequestHandler) New() runtime.Object {
	return &inputinternal.ActionCall{}
}

func (h *actionRequestHandler) NamespaceScoped() bool {
	return true
}

func (h *actionRequestHandler) Create(
	ctx context.Context,
	obj runtime.Object,
	createValidation rest.ValidateObjectFunc,
	options *metav1.CreateOptions,
) (runtime.Object, error) {
	t := trace.FromContext(ctx).Nest("create", trace.Field{
		Key:   "kind",
		Value: "ActionRequest",
	})
	defer t.Log()

	// Run the provided creation validations
	if createValidation != nil {
		if err := createValidation(ctx, obj); err != nil {
			traceFailure(t, err.Error())
			return nil, err
		}
	}

	// Cast the input object
	actionCall, ok := obj.(*inputinternal.ActionCall)
	if !ok {
		return nil, apierrors.NewBadRequest(fmt.Sprintf("not an ActionRequest: %#v", obj))
	}

	// Post to the actions sink and return the result
	result := inputinternal.ActionCallResultAccepted
	message := ""
	if err := h.ActionSink.Post(
		genericrequest.NamespaceValue(ctx),
		actionCall.Spec.Spirit.Name,
		actionCall.Spec.Battle.Name,
		actionCall.Spec.ActionName,
	); err != nil {
		result = inputinternal.ActionCallResultRejected
		message = err.Error()
		traceFailure(t, message)
	} else {
		traceSuccess(t)
	}
	return &inputinternal.ActionCall{
		Status: inputinternal.ActionCallStatus{
			Result:  result,
			Message: message,
		},
	}, nil
}

func traceSuccess(t *trace.Trace) {
	t.Step("success")
}

func traceFailure(t *trace.Trace, message string) {
	t.Step("failure",
		trace.Field{Key: "message", Value: message},
	)
}
