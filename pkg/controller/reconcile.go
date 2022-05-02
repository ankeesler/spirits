package controller

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type reconcileHelper[T client.Object] struct {
	// Required

	client.Client
	Object T

	// Optional

	OnUpsert func(context.Context, logr.Logger, ctrl.Request, T) error
	OnDelete func(context.Context, logr.Logger, ctrl.Request) error
}

func reconcile[T client.Object](
	ctx context.Context,
	req ctrl.Request,
	reconcileHelper *reconcileHelper[T],
) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	if err := reconcileHelper.Get(ctx, req.NamespacedName, reconcileHelper.Object); err != nil {
		if apierrors.IsNotFound(err) {
			if reconcileHelper.OnDelete != nil {
				if err := reconcileHelper.OnDelete(ctx, log, req); err != nil {
					return ctrl.Result{}, fmt.Errorf("handle delete: %w", err)
				}
			}
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, fmt.Errorf("get: %w", err)
	}

	if _, err := controllerutil.CreateOrPatch(ctx, reconcileHelper.Client, reconcileHelper.Object, func() error {
		if reconcileHelper.OnUpsert != nil {
			if err := reconcileHelper.OnUpsert(ctx, log, req, reconcileHelper.Object); err != nil {
				return fmt.Errorf("handle upsert: %w", err)
			}
		}
		return nil
	}); err != nil {
		return ctrl.Result{}, fmt.Errorf("create or patch %w", err)
	}

	return ctrl.Result{}, nil
}
