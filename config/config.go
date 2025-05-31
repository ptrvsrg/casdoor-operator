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

//nolint:lll

package config

import "time"

type Config struct {
	SyncPeriod             time.Duration `yaml:"syncPeriod" env:"SYNC_PERIOD" env-default:"5m"`
	WatchNamespaces        string        `yaml:"watchNamespaces" env:"WATCH_NAMESPACES"`
	CustomResourceSelector string        `yaml:"customResourceSelector" env:"CUSTOM_RESOURCE_SELECTOR"`
	EnableHTTP2            *bool         `yaml:"enableHTTP2" env:"ENABLE_HTTP2" env-default:"true"`

	Metrics             MetricsConfig             `yaml:"metrics"`
	Probe               ProbeConfig               `yaml:"probe"`
	Pprof               PprofConfig               `yaml:"pprof"`
	Logging             LoggingConfig             `yaml:"logging"`
	LeaderElection      LeaderElectionConfig      `yaml:"leaderElection"`
	GlobalController    GlobalControllerConfig    `yaml:"globalController"`
	SpecificControllers SpecificControllersConfig `yaml:"specificControllers"`
}

type LoggingConfig struct {
	Level  string `yaml:"level" env:"LOG_LEVEL" env-default:"info"`
	Format string `yaml:"format" env:"LOG_FORMAT" env-default:"json"`
}

type LeaderElectionConfig struct {
	Enabled       *bool         `yaml:"enabled" env:"LEADER_ELECTION_ENABLED" env-default:"false"`
	ID            string        `yaml:"id" env:"LEADER_ELECTION_ID" env-default:"casdoor-operator-leader-election"`
	Namespace     string        `yaml:"namespace" env:"LEADER_ELECTION_NAMESPACE" env-default:"default"`
	RetryPeriod   time.Duration `yaml:"retryPeriod" env:"LEADER_ELECTION_RETRY_PERIOD" env-default:"2s"`
	RenewDeadline time.Duration `yaml:"renewDeadline" env:"LEADER_ELECTION_RENEW_DEADLINE" env-default:"10s"`
	LeaseDuration time.Duration `yaml:"leaseDuration" env:"LEADER_ELECTION_LEASE_DURATION" env-default:"15s"`
}

type MetricsConfig struct {
	Enabled *bool `yaml:"enabled" env:"METRICS_ENABLED" env-default:"false"`
	Port    int32 `yaml:"port" env:"METRICS_PORT" env-default:"8080"`
	Secure  *bool `yaml:"secure" env:"METRICS_SECURE" env-default:"false"`
}

type ProbeConfig struct {
	Enabled               *bool  `yaml:"enabled" env:"PROBE_ENABLED" env-default:"false"`
	Port                  int    `yaml:"port" env:"PROBE_PORT" env-default:"8082"`
	ReadinessEndpointName string `yaml:"readinessEndpointName" env:"PROBE_READINESS_ENDPOINT_NAME" env-default:"readyz"`
	LivenessEndpointName  string `yaml:"livenessEndpointName" env:"PROBE_LIVENESS_ENDPOINT_NAME" env-default:"healthz"`
}

type PprofConfig struct {
	Enabled *bool `yaml:"enabled" env:"PPROF_ENABLED" env-default:"false"`
	Port    int   `yaml:"port" env:"PPROF_PORT" env-default:"8081"`
}

type GlobalControllerConfig struct {
	SkipNameValidation      *bool          `yaml:"skipNameValidation" env:"GLOBAL_CONTROLLER_SKIP_NAME_VALIDATION"`
	GroupKindConcurrency    map[string]int `yaml:"groupKindConcurrency" env:"GLOBAL_CONTROLLER_GROUP_KIND_CONCURRENCY"`
	MaxConcurrentReconciles int            `yaml:"maxConcurrentReconciles" env:"GLOBAL_CONTROLLER_MAX_CONCURRENT_RECONCILES" env-default:"1" validate:"min=1"`
	CacheSyncTimeout        time.Duration  `yaml:"cacheSyncTimeout" env:"GLOBAL_CONTROLLER_CACHE_SYNC_TIMEOUT" env-default:"2m"`
	NeedLeaderElection      *bool          `yaml:"needLeaderElection" env:"GLOBAL_CONTROLLER_NEED_LEADER_ELECTION"`
}

type SpecificControllersConfig struct {
	Casdoor CasdoorControllerConfig `yaml:"casdoor"`
}

type CasdoorControllerConfig struct {
	SkipNameValidation      *bool         `yaml:"skipNameValidation" env:"CASDOOR_CONTROLLER_SKIP_NAME_VALIDATION"`
	MaxConcurrentReconciles int           `yaml:"maxConcurrentReconciles" env:"CASDOOR_CONTROLLER_MAX_CONCURRENT_RECONCILES" env-default:"1" validate:"min=1"`
	CacheSyncTimeout        time.Duration `yaml:"cacheSyncTimeout" env:"CASDOOR_CONTROLLER_CACHE_SYNC_TIMEOUT" env-default:"2m"`
	NeedLeaderElection      *bool         `yaml:"needLeaderElection" env:"CASDOOR_CONTROLLER_NEED_LEADER_ELECTION"`
}
