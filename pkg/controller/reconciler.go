package controller

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
)

// TODO: this type could be a function...

type handler[ExternalT, InternalT client.Object] interface {
	NewExternal() ExternalT
	NewInternal() InternalT
	OnUpsert(context.Context, logr.Logger, ctrl.Request, InternalT) error
	OnDelete(context.Context, logr.Logger, ctrl.Request) error
}

type reconciler[ExternalT, InternalT client.Object] struct {
	client.Client
	Scheme  *runtime.Scheme
	Handler handler[ExternalT, InternalT]
}

var _ reconcile.Reconciler = &reconciler[*spiritsv1alpha1.Spirit, *spiritsinternal.Spirit]{}

func (r *reconciler[ExternalT, InternalT]) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	externalT := r.Handler.NewExternal()
	if err := r.Get(ctx, req.NamespacedName, externalT); err != nil {
		if apierrors.IsNotFound(err) {
			if err := r.Handler.OnDelete(ctx, log, req); err != nil {
				return ctrl.Result{}, fmt.Errorf("handle delete: %w", err)
			}
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, fmt.Errorf("get external object %w", err)
	}

	if _, err := controllerutil.CreateOrPatch(ctx, r.Client, externalT, func() error {
		internalT := r.Handler.NewInternal()
		if err := r.Scheme.Convert(externalT, internalT, nil); err != nil {
			return fmt.Errorf("convert from external object to internal object %w", err)
		}

		if err := r.Handler.OnUpsert(ctx, log, req, internalT); err != nil {
			return fmt.Errorf("handle upsert: %w", err)
		}

		if err := r.Scheme.Convert(internalT, externalT, nil); err != nil {
			return fmt.Errorf("convert from internal object to external object %w", err)
		}

		return nil
	}); err != nil {
		return ctrl.Result{}, fmt.Errorf("create or patch external object %w", err)
	}

	return ctrl.Result{}, nil
}
