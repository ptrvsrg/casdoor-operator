# Configuration

## Config schema reference

| Property                 | Environment variable       | Property type                                                           | Default | Description                                                                                        |
|--------------------------|----------------------------|-------------------------------------------------------------------------|---------|----------------------------------------------------------------------------------------------------|
| `watchNamespaces`        | `WATCH_NAMESPACES`         | string                                                                  | ""      | A comma-separated list of namespaces to watch for resources. If empty, all namespaces are watched. |
| `customResourceSelector` | `CUSTOM_RESOURCE_SELECTOR` | string                                                                  | ""      | A selector for custom resources.                                                                   |
| `enableHTTP2`            | `ENABLE_HTTP2`             | boolean                                                                 | true    | Enables or disables HTTP/2 support for the service.                                                |
| `metrics`                |                            | [MetricsConfig](#metricsconfig-schema-reference)                        |         | Configuration for metrics collection and exposure.                                                 |
| `probe`                  |                            | [ProbeConfig](#probeconfig-schema-reference)                            |         | Configuration for health and readiness probes.                                                     |
| `pprof`                  |                            | [PprofConfig](#pprofconfig-schema-reference)                            |         | Configuration for pprof profiling endpoints.                                                       |
| `logging`                |                            | [LoggingConfig](#loggingconfig-schema-reference)                        |         | Configuration for logging behavior and format.                                                     |
| `leaderElection`         |                            | [LeaderElectionConfig](#leaderelectionconfig-schema-reference)          |         | Configuration for leader election in a multi-instance setup.                                       |
| `globalController`       |                            | [GlobalControllerConfig](#globalcontrollerconfig-schema-reference)      |         | Global configuration for controllers managing resources.                                           |
| `specificControllers`    |                            | [SpecificControllerConfig](#specificcontrollersconfig-schema-reference) |         | Configuration specific to individual controllers.                                                  |

## MetricsConfig schema reference

Used in [Config](#config-schema-reference)

| Property  | Environment variable | Property type | Default | Description                                          |
|-----------|----------------------|---------------|---------|------------------------------------------------------|
| `enabled` | `METRICS_ENABLED`    | boolean       | false   | Enables or disables the metrics server.              |
| `port`    | `METRICS_PORT`       | integer       | 8080    | The port on which the metrics server listens.        |
| `secure`  | `METRICS_SECURE`     | boolean       | false   | Enables or disables secure (HTTPS) metrics endpoint. |

## ProbeConfig schema reference

Used in [Config](#config-schema-reference)

| Property                | Environment variable            | Property type | Default   | Description                                        |
|-------------------------|---------------------------------|---------------|-----------|----------------------------------------------------|
| `enabled`               | `PROBE_ENABLED`                 | boolean       | false     | Enables or disables health and readiness probes.   |
| `port`                  | `PROBE_PORT`                    | integer       | 8082      | The port on which the probe endpoints are exposed. |
| `readinessEndpointName` | `PROBE_READINESS_ENDPOINT_NAME` | string        | "readyz"  | The endpoint path for readiness checks.            |
| `livenessEndpointName`  | `PROBE_LIVENESS_ENDPOINT_NAME`  | string        | "healthz" | The endpoint path for liveness checks.             |

## PprofConfig schema reference

Used in [Config](#config-schema-reference)

| Property  | Environment variable | Property type | Default | Description                                        |
|-----------|----------------------|---------------|---------|----------------------------------------------------|
| `enabled` | `PPROF_ENABLED`      | boolean       | false   | Enables or disables the pprof profiling endpoints. |
| `port`    | `PPROF_PORT`         | integer       | 8081    | The port on which the pprof endpoints are exposed. |

## LoggingConfig schema reference

Used in [Config](#config-schema-reference)

| Property | Environment variable | Property type                                     | Default | Description                                                                       |
|----------|----------------------|---------------------------------------------------|---------|-----------------------------------------------------------------------------------|
| `level`  | `LOG_LEVEL`          | string (one of [debug, info, warn, error, fatal]) | "info"  | The log level for the application. Determines the verbosity of logs.              |
| `format` | `LOG_FORMAT`         | string (one of [console, json])                   | "json"  | The format of the logs. Can be human-readable (`console`) or structured (`json`). |

## LeaderElectionConfig schema reference

Used in [Config](#config-schema-reference)

| Property        | Environment variable             | Property type | Default                            | Description                                                 |
|-----------------|----------------------------------|---------------|------------------------------------|-------------------------------------------------------------|
| `enabled`       | `LEADER_ELECTION_ENABLED`        | boolean       | false                              | Enables or disables leader election.                        |
| `id`            | `LEADER_ELECTION_ID`             | string        | "casdoor-operator-leader-election" | The unique identifier for the leader election process.      |
| `namespace`     | `LEADER_ELECTION_NAMESPACE`      | string        | "default"                          | The namespace where leader election resources are created.  |
| `retryPeriod`   | `LEADER_ELECTION_RETRY_PERIOD`   | string        | "2s"                               | The interval at which leader election attempts are retried. |
| `renewDeadline` | `LEADER_ELECTION_RENEW_DEADLINE` | string        | "10s"                              | The deadline by which the leader must renew its lease.      |
| `leaseDuration` | `LEADER_ELECTION_LEASE_DURATION` | string        | "15s"                              | The duration for which a leader's lease is valid.           |

## GlobalControllerConfig schema reference

Used in [Config](#config-schema-reference)

| Property                  | Environment variable                          | Property type  | Default | Description                                                        |
|---------------------------|-----------------------------------------------|----------------|---------|--------------------------------------------------------------------|
| `skipNameValidation`      | `GLOBAL_CONTROLLER_SKIP_NAME_VALIDATION`      | boolean        | false   | Skips validation of resource names if enabled.                     |
| `groupKindConcurrency`    | `GLOBAL_CONTROLLER_GROUP_KIND_CONCURRENCY`    | map[string]int |         | Specifies the concurrency limits for different resource types.     |
| `maxConcurrentReconciles` | `GLOBAL_CONTROLLER_MAX_CONCURRENT_RECONCILES` | integer        | 1       | The maximum number of concurrent reconciliation operations.        |
| `cacheSyncTimeout`        | `GLOBAL_CONTROLLER_CACHE_SYNC_TIMEOUT`        | string         | "2m"    | The timeout for syncing the cache with the API server.             |
| `needLeaderElection`      | `GLOBAL_CONTROLLER_NEED_LEADER_ELECTION`      | boolean        | true    | Indicates whether leader election is required for this controller. |

## SpecificControllersConfig schema reference

Used in [Config](#config-schema-reference)

| Property  | Environment variable | Property type                                                        | Default | Description                                       |
|-----------|----------------------|----------------------------------------------------------------------|---------|---------------------------------------------------|
| `casdoor` |                      | [CasdoorControllerConfig](#casdoorcontrollerconfig-schema-reference) |         | Configuration specific to the Casdoor controller. |

## CasdoorControllerConfig schema reference

Used in [SpecificControllersConfig](#specificcontrollersconfig-schema-reference)

| Property                  | Environment variable                           | Property type | Default | Description                                                                       |
|---------------------------|------------------------------------------------|---------------|---------|-----------------------------------------------------------------------------------|
| `skipNameValidation`      | `CASDOOR_CONTROLLER_SKIP_NAME_VALIDATION`      | boolean       | false   | Skips validation of Casdoor resource names if enabled.                            |
| `maxConcurrentReconciles` | `CASDOOR_CONTROLLER_MAX_CONCURRENT_RECONCILES` | integer       | 1       | The maximum number of concurrent reconciliation operations for Casdoor resources. |
| `cacheSyncTimeout`        | `CASDOOR_CONTROLLER_CACHE_SYNC_TIMEOUT`        | string        | "2m"    | The timeout for syncing the cache with the API server for Casdoor resources.      |
| `needLeaderElection`      | `CASDOOR_CONTROLLER_NEED_LEADER_ELECTION`      | boolean       | true    | Indicates whether leader election is required for the Casdoor controller.         |
