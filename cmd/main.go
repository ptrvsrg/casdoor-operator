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

package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-logr/zapr"
	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"resty.dev/v3"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	kconfig "sigs.k8s.io/controller-runtime/pkg/config"
	"sigs.k8s.io/controller-runtime/pkg/healthz"

	casdoorv1alpha1 "github.com/ptrvsrg/casdoor-operator/api/v1alpha1"
	"github.com/ptrvsrg/casdoor-operator/config"
	"github.com/ptrvsrg/casdoor-operator/internal/controller"
	"github.com/ptrvsrg/casdoor-operator/internal/logging"
	"github.com/ptrvsrg/casdoor-operator/internal/version"

	"github.com/urfave/cli/v3"

	_ "github.com/joho/godotenv/autoload"
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/metrics/filters"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	// +kubebuilder:scaffold:imports
)

const (
	gracefulShutdownTimeout = 10 * time.Minute
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")

	cmd = &cli.Command{
		Name:                   os.Args[0],
		Version:                version.AppVersion,
		Authors:                []any{"ptrvsrg"},
		Copyright:              fmt.Sprintf("© %d ptrvsrg", time.Now().Year()),
		Usage:                  "The cli application for Crack-Hash manager",
		UseShortOptionHandling: true,
		EnableShellCompletion:  true,
		Action:                 runManager,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "Config file path",
				HideDefault: false,
				Required:    false,
				Local:       true,
				Value:       "config.yaml",
				Sources:     cli.NewValueSourceChain(cli.EnvVar("CONFIG_FILE_PATH")),
			},
			&cli.BoolFlag{
				Name:        "version",
				Aliases:     []string{"v"},
				Usage:       "Print version information",
				HideDefault: false,
				Required:    false,
				Local:       false,
				Action: func(_ context.Context, _ *cli.Command, enabled bool) error {
					if enabled {
						fmt.Printf("Application: %s\n", version.AppVersion)                 //nolint:forbidigo
						fmt.Printf("Runtime: %s %s\n", version.GoVersion, version.Platform) //nolint:forbidigo
					}
					return nil
				},
			},
		},
	}
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(casdoorv1alpha1.AddToScheme(scheme))
	// +kubebuilder:scaffold:scheme
}

func main() {
	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		setupLog.Error(err, "failed to run command")
	}
}

func runManager(ctx context.Context, command *cli.Command) error { //nolint:cyclop,funlen,gocognit
	// Setup base logging
	logging.SetupLogger(
		config.LoggingConfig{
			Level:  "info",
			Format: "console",
		},
	)
	ctrl.SetLogger(zapr.NewLogger(zap.L()))

	// Read config
	var cfg config.Config
	configPath := command.String("config")
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	// Validate config
	err = validator.New().Struct(cfg)
	if err != nil {
		return fmt.Errorf("failed to validate config: %w", err)
	}

	// Setup logging
	logging.SetupLogger(cfg.Logging)
	ctrl.SetLogger(zapr.NewLogger(zap.L()))

	// Print config
	setupLog.Info("loaded config", "config", cfg)

	// Setup TLS
	var tlsOpts []func(*tls.Config)
	if cfg.EnableHTTP2 != nil && !(*cfg.EnableHTTP2) {
		tlsOpts = append(tlsOpts, disableHTTP2)
	}

	// Setup metrics
	metricsServerOptions := metricsserver.Options{
		BindAddress: "0",
	}
	if cfg.Metrics.Enabled != nil && *cfg.Metrics.Enabled {
		metricsServerOptions.BindAddress = fmt.Sprintf(":%d", cfg.Metrics.Port)
		metricsServerOptions.SecureServing = *cfg.Metrics.Secure
		metricsServerOptions.TLSOpts = tlsOpts

		if cfg.Metrics.Secure != nil && *cfg.Metrics.Secure {
			metricsServerOptions.FilterProvider = filters.WithAuthenticationAndAuthorization
		}
	}

	// Setup probes
	probeAddr := ""
	if cfg.Probe.Enabled != nil && *cfg.Probe.Enabled {
		probeAddr = fmt.Sprintf(":%d", cfg.Probe.Port)
	}

	// Setup pprof
	pprofAddr := ""
	if cfg.Pprof.Enabled != nil && *cfg.Pprof.Enabled {
		pprofAddr = fmt.Sprintf(":%d", cfg.Pprof.Port)
	}

	// Setup webhook
	webhookServer := webhook.NewServer(
		webhook.Options{
			TLSOpts: tlsOpts,
		},
	)

	// Setup cache
	cacheOpts := cache.Options{
		DefaultNamespaces: make(map[string]cache.Config),
	}
	for _, namespace := range strings.Split(cfg.WatchNamespaces, ",") {
		cacheOpts.DefaultNamespaces[namespace] = cache.Config{}
	}

	// Build manager options
	managerOpts := ctrl.Options{
		Logger:                        zapr.NewLogger(zap.L().Named("manager")),
		Scheme:                        scheme,
		Metrics:                       metricsServerOptions,
		WebhookServer:                 webhookServer,
		HealthProbeBindAddress:        probeAddr,
		PprofBindAddress:              pprofAddr,
		Cache:                         cacheOpts,
		LeaderElection:                cfg.LeaderElection.Enabled != nil && *cfg.LeaderElection.Enabled,
		LeaderElectionID:              cfg.LeaderElection.ID,
		LeaderElectionNamespace:       cfg.LeaderElection.Namespace,
		LeaderElectionReleaseOnCancel: true,
		RetryPeriod:                   &cfg.LeaderElection.RetryPeriod,
		LeaseDuration:                 &cfg.LeaderElection.LeaseDuration,
		RenewDeadline:                 &cfg.LeaderElection.RenewDeadline,
		GracefulShutdownTimeout:       lo.ToPtr(gracefulShutdownTimeout),
		Controller: kconfig.Controller{
			SkipNameValidation:      cfg.GlobalController.SkipNameValidation,
			GroupKindConcurrency:    cfg.GlobalController.GroupKindConcurrency,
			MaxConcurrentReconciles: cfg.GlobalController.MaxConcurrentReconciles,
			CacheSyncTimeout:        cfg.GlobalController.CacheSyncTimeout,
			NeedLeaderElection:      cfg.GlobalController.NeedLeaderElection,
		},
	}

	// Create manager
	setupLog.Info("creating manager")
	kubeconfig := ctrl.GetConfigOrDie()
	mgr, err := ctrl.NewManager(kubeconfig, managerOpts)
	if err != nil {
		return fmt.Errorf("failed to create manager: %w", err)
	}

	if err := mgr.AddHealthzCheck(cfg.Probe.LivenessEndpointName, healthz.Ping); err != nil {
		return fmt.Errorf("failed to set up health check: %w", err)
	}
	if err := mgr.AddReadyzCheck(cfg.Probe.ReadinessEndpointName, healthz.Ping); err != nil {
		return fmt.Errorf("failed to set up ready check: %w", err)
	}

	// Register controllers
	controllers := []controller.Controller{
		controller.NewCasdoorReconciler(mgr.GetClient(), mgr.GetScheme(), cfg.SpecificControllers.Casdoor),
	}
	if err = controller.SetupWithManager(ctx, mgr, controllers...); err != nil {
		return fmt.Errorf("failed to setup controllers: %w", err)
	}

	defer func(ctrls ...controller.Controller) {
		err := controller.Close(ctrls...)
		if err != nil {
			setupLog.Error(err, "failed to close controller")
		}
	}(controllers...)

	// +kubebuilder:scaffold:builder

	// Start the manager
	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		return fmt.Errorf("failed to start manager: %w", err)
	}

	return nil
}

func disableHTTP2(c *tls.Config) {
	setupLog.Info("disabling http/2")
	c.NextProtos = []string{"http/1.1"}
}
