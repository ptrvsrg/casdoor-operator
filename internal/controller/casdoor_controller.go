/*
Copyright 2025 ptrvsrg.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	casdoorv1alpha1 "github.com/ptrvsrg/casdoor-operator/api/v1alpha1"
	"github.com/ptrvsrg/casdoor-operator/config"
	"github.com/ptrvsrg/casdoor-operator/internal/logging"
)

// CasdoorReconciler reconciles a Casdoor object
type CasdoorReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	logger   *zap.Logger
	cfg      config.CasdoorControllerConfig
	recorder record.EventRecorder
}

func NewCasdoorReconciler(
	client client.Client, scheme *runtime.Scheme, cfg config.CasdoorControllerConfig,
) *CasdoorReconciler {
	return &CasdoorReconciler{
		Client: client,
		Scheme: scheme,
		cfg:    cfg,
	}
}

// RBAC for Casdoor
// +kubebuilder:rbac:groups=casdoor.ptrvsrg.github.com,resources=casdoors,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=casdoor.ptrvsrg.github.com,resources=casdoors/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=casdoor.ptrvsrg.github.com,resources=casdoors/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.0/pkg/reconcile
func (r *CasdoorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// Setup loggers
	logger := log.FromContext(ctx)
	logger.Info("start to reconcile Casdoor")

	casdoor := &casdoorv1alpha1.Casdoor{}
	r.logger = logging.GetReconcileLogger(req, casdoor)

	// Get Casdoor
	if err := r.Get(ctx, req.NamespacedName, casdoor); err != nil {
		if kerrors.IsNotFound(err) {
			r.logger.Info("resource Casdoor was deleted before reconcile")
			return ctrl.Result{}, nil
		}

		r.logger.Error("failed to get Casdoor", zap.Error(err))
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Delete Casdoor
	if !casdoor.ObjectMeta.DeletionTimestamp.IsZero() {
		return r.delete(ctx, casdoor)
	}

	// Add finalizer
	if err := r.addFinalizer(ctx, casdoor); err != nil {
		r.logger.Error("failed to add finalizer", zap.Error(err))
		return ctrl.Result{}, err
	}

	// Skip reconcile
	if _, ok := casdoor.Annotations[skipReconcileAnnotation]; ok {
		logger.Info("skip reconcile Casdoor")
		return ctrl.Result{}, nil
	}

	logger.Info("reconcile Casdoor successfully")

	return ctrl.Result{}, nil
}

func (r *CasdoorReconciler) delete(ctx context.Context, casdoor *casdoorv1alpha1.Casdoor) (ctrl.Result, error) {
	if err := r.Delete(ctx, casdoor); err != nil {
		if kerrors.IsNotFound(err) {
			r.logger.Info("resource Casdoor was deleted before reconcile")
			return ctrl.Result{}, nil
		}

		r.logger.Error("failed to delete Casdoor", zap.Error(err))
		return ctrl.Result{}, fmt.Errorf("failed to delete Casdoor: %w", err)
	}

	r.logger.Info("resource Casdoor was deleted")
	r.recorder.Eventf(casdoor, v1.EventTypeNormal, "Deleted", "Casdoor %s was deleted", casdoor.Name)

	return ctrl.Result{}, nil
}

func (r *CasdoorReconciler) addFinalizer(ctx context.Context, casdoor *casdoorv1alpha1.Casdoor) error {
	if controllerutil.AddFinalizer(casdoor, buildDefaultFinalizerName(casdoor)) {
		if err := r.Status().Update(ctx, casdoor); err != nil {
			r.logger.Error("failed to add Casdoor finalizer", zap.Error(err))
			return err
		}
	}

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CasdoorReconciler) SetupWithManager(_ context.Context, mgr ctrl.Manager) error {
	r.recorder = mgr.GetEventRecorderFor(
		(&casdoorv1alpha1.Casdoor{}).GetResourceKind().String(),
	)

	err := ctrl.NewControllerManagedBy(mgr).
		For(
			&casdoorv1alpha1.Casdoor{},
		).
		WithOptions(
			controller.Options{
				SkipNameValidation:      r.cfg.SkipNameValidation,
				MaxConcurrentReconciles: r.cfg.MaxConcurrentReconciles,
				CacheSyncTimeout:        r.cfg.CacheSyncTimeout,
				NeedLeaderElection:      r.cfg.NeedLeaderElection,
			},
		).
		WithEventFilter(
			predicate.GenerationChangedPredicate{},
		).
		Complete(r)
	if err != nil {
		return fmt.Errorf("failed to create application controller: %w", err)
	}

	return nil
}

func (r *CasdoorReconciler) Close() error {
	return nil
}
