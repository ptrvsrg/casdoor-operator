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
	url2 "net/url"
	"path"
	"time"

	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"resty.dev/v3"
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
	Scheme     *runtime.Scheme
	logger     *zap.Logger
	cfg        config.CasdoorControllerConfig
	httpClient *resty.Client
	recorder   record.EventRecorder
}

func NewCasdoorReconciler(
	client client.Client, scheme *runtime.Scheme, cfg config.CasdoorControllerConfig, httpClient *resty.Client,
) *CasdoorReconciler {
	return &CasdoorReconciler{
		Client:     client,
		Scheme:     scheme,
		cfg:        cfg,
		httpClient: httpClient,
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

	// Setup RequeueAfter
	var requeueAfter time.Duration
	if casdoor.Spec.Healthcheck != nil && casdoor.Spec.Healthcheck.Enabled {
		requeueAfter = casdoor.Spec.Healthcheck.Interval.Duration
	}

	// Delete Casdoor
	if !casdoor.ObjectMeta.DeletionTimestamp.IsZero() {
		return r.delete(ctx, casdoor)
	}

	// Skip reconcile
	if _, ok := casdoor.Annotations[skipReconcileAnnotation]; ok {
		logger.Info("skip reconcile Casdoor")
		return ctrl.Result{}, nil
	}

	// Add finalizer
	if err := r.addFinalizer(ctx, casdoor); err != nil {
		r.logger.Error("failed to add finalizer", zap.Error(err))
		return ctrl.Result{RequeueAfter: requeueAfter}, err
	}

	// Check health
	if err := r.checkHealth(ctx, casdoor); err != nil {
		// Error health
		r.logger.Error("failed to check health", zap.Error(err))

		if err := r.applyErrorHealthcheck(ctx, casdoor, err); err != nil {
			r.logger.Error("failed to apply error healthcheck", zap.Error(err))
			return ctrl.Result{RequeueAfter: requeueAfter}, err
		}

		return ctrl.Result{RequeueAfter: requeueAfter}, err
	}

	// Successful health
	if err := r.applySuccessfulHealthcheck(ctx, casdoor); err != nil {
		r.logger.Error("failed to apply successful healthcheck", zap.Error(err))
		return ctrl.Result{RequeueAfter: requeueAfter}, err
	}

	logger.Info("reconcile Casdoor successfully")

	return ctrl.Result{RequeueAfter: requeueAfter}, nil
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
			r.logger.Error("failed to add finalizer", zap.Error(err))
			return err
		}
	}

	return nil
}

func (r *CasdoorReconciler) checkHealth(ctx context.Context, casdoor *casdoorv1alpha1.Casdoor) error {
	if casdoor.Spec.Healthcheck == nil || !casdoor.Spec.Healthcheck.Enabled {
		return nil
	}

	url, err := url2.Parse(casdoor.Spec.URL)
	if err != nil {
		return fmt.Errorf("failed to parse url: %w", err)
	}

	url.Path = path.Join(url.Path, casdoor.Spec.Healthcheck.Path)

	r.logger.Debug(
		"send health check request",
		zap.String("url", url.String()),
		zap.String("method", casdoor.Spec.Healthcheck.Method),
	)
	res, err := r.httpClient.
		R().
		SetContext(ctx).
		SetRetryCount(casdoor.Spec.Healthcheck.Retries).
		SetTimeout(casdoor.Spec.Healthcheck.Timeout.Duration).
		Execute(casdoor.Spec.Healthcheck.Method, url.String())

	if err != nil {
		return fmt.Errorf("failed to check health: %w", err)
	}

	if res.IsError() {
		return fmt.Errorf(
			"failed to check health: status code not success: status=%s body=%s",
			res.Status(),
			res.String(),
		)
	}

	return nil
}

func (r *CasdoorReconciler) applySuccessfulHealthcheck(ctx context.Context, casdoor *casdoorv1alpha1.Casdoor) error {
	// Send event
	r.recorder.Eventf(casdoor, v1.EventTypeNormal, "Ready", "Casdoor %s is ready", casdoor.Name)

	// Update status
	casdoor.Status.Code = casdoorv1alpha1.CasdoorStatusReady
	casdoor.Status.Reason = ""
	if err := r.Status().Update(ctx, casdoor); err != nil {
		r.logger.Error("failed to update status", zap.Error(err))
		return err
	}

	return nil
}

func (r *CasdoorReconciler) applyErrorHealthcheck(
	ctx context.Context, casdoor *casdoorv1alpha1.Casdoor, healthCheckError error,
) error {
	// Send event
	r.recorder.Eventf(
		casdoor,
		v1.EventTypeWarning,
		"Failed",
		"Casdoor %s is failed: %s",
		casdoor.Name,
		healthCheckError.Error(),
	)

	// Update status
	casdoor.Status.Code = casdoorv1alpha1.CasdoorStatusFailed
	casdoor.Status.Reason = healthCheckError.Error()
	if err := r.Status().Update(ctx, casdoor); err != nil {
		r.logger.Error("failed to update status", zap.Error(err))
		return err
	}

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CasdoorReconciler) SetupWithManager(_ context.Context, mgr ctrl.Manager) error {
	r.recorder = mgr.GetEventRecorderFor((&casdoorv1alpha1.Casdoor{}).GetResourceKind().String())

	err := ctrl.NewControllerManagedBy(mgr).
		For(&casdoorv1alpha1.Casdoor{}).
		WithOptions(
			controller.Options{
				SkipNameValidation:      r.cfg.SkipNameValidation,
				MaxConcurrentReconciles: r.cfg.MaxConcurrentReconciles,
				CacheSyncTimeout:        r.cfg.CacheSyncTimeout,
				NeedLeaderElection:      r.cfg.NeedLeaderElection,
			},
		).
		WithEventFilter(predicate.GenerationChangedPredicate{}).
		Complete(r)
	if err != nil {
		return fmt.Errorf("failed to create application controller: %w", err)
	}

	return nil
}

func (r *CasdoorReconciler) Close() error {
	return nil
}
