# Casdoor Operator

## About The Project

Simply put, the Casdoor Operator is a tool for automating the resource management of Casdoor instances in Kubernetes
environments.

Casdoor is an open-source identity and access management platform that provides authentication, authorization, and user
management capabilities. Managing Casdoor deployments manually in Kubernetes can be repetitive, error-prone, and
time-consuming. This is where the Casdoor Operator comes in to streamline the process.

## Why We Built It

The Casdoor Operator was created to solve a common problem faced by developers and DevOps teams: deploying and managing
applications in Kubernetes is complex and often tedious. Without automation, you end up writing custom scripts or
performing manual steps to handle tasks like provisioning, configuration, scaling, and health checks.

Typically, these scripts are not reusable, difficult to maintain, and often depend on multiple tools, making it hard to
replicate the process across different environments. The Casdoor Operator eliminates the need for these custom solutions
by providing a robust, reusable, and Kubernetes-native way to manage Casdoor instances.

With the Casdoor Operator, instead of writing complex scripts or manually managing deployments, you define your desired
state in a simple custom resource definition (CRD) file. The operator takes care of the rest, ensuring that your Casdoor
instances are deployed, configured, and maintained according to your specifications.

## How It Helps

We have worked hard to make the Casdoor Operator intuitive and powerful, focusing on simplicity, security, and
scalability. Here’s how it benefits users:

* **Automation:** Simplifies the deployment and lifecycle management of Casdoor instances.
* **Consistency:** Ensures consistent configurations across environments using Kubernetes-native tools.
* **Health Monitoring:** Continuously monitors the health of Casdoor instances and updates their status accordingly.
* **Extensibility:** Provides flexibility to customize deployments based on your specific needs.
* **Security:** Supports secure practices such as secret management and signed artifacts.

Our goal is to empower users to focus on building great applications while the Casdoor Operator handles the operational
complexities of managing Casdoor in Kubernetes.

## Parameters

The following tables list the configurable parameters of the Casdoor Operator Helm chart and their default values.

### Global Configuration

| Parameter                | Description                                    | Default Value |
|--------------------------|------------------------------------------------|---------------|
| `replicas`               | Number of replicas for the operator deployment | `1`           |
| `nameOverride`           | Override the name of the chart                 | `""`          |
| `fullnameOverride`       | Override the full name of the chart            | `""`          |
| `revisionHistoryLimit`   | Number of old ReplicaSets to retain            | `10`          |
| `enableHTTP2`            | Enable HTTP/2 support                          | `true`        |
| `watchNamespaces`        | List of namespaces the operator should watch   | `[]`          |
| `customResourceSelector` | Selector for custom resources                  | `""`          |

---

### Image Configuration

| Parameter          | Description                                                       | Default Value      |
|--------------------|-------------------------------------------------------------------|--------------------|
| `image.registry`   | Docker registry for the operator image                            | `ghcr.io`          |
| `image.repository` | Repository for the operator image                                 | `ptrvsrg`          |
| `image.name`       | Name of the operator image                                        | `casdoor-operator` |
| `image.tag`        | Tag for the operator image (defaults to the Chart's `appVersion`) | `""`               |
| `image.pullPolicy` | Image pull policy                                                 | `IfNotPresent`     |
| `imagePullSecrets` | Secrets for pulling images from private registries                | `[]`               |

---

### Service Configuration

| Parameter                 | Description                                                    | Default Value |
|---------------------------|----------------------------------------------------------------|---------------|
| `service.type`            | Type of service (`ClusterIP`, `NodePort`, `LoadBalancer`)      | `ClusterIP`   |
| `service.metricsPort`     | Port for metrics                                               | `8080`        |
| `service.probePort`       | Port for health probes                                         | `8082`        |
| `service.pprofPort`       | Port for pprof debugging                                       | `8081`        |
| `service.metricsNodePort` | NodePort for metrics (if `service.type` is `NodePort`)         | `30080`       |
| `service.probeNodePort`   | NodePort for health probes (if `service.type` is `NodePort`)   | `30082`       |
| `service.pprofNodePort`   | NodePort for pprof debugging (if `service.type` is `NodePort`) | `30081`       |

---

### Metrics and Probes

#### Metrics Configuration

| Parameter         | Description               | Default Value |
|-------------------|---------------------------|---------------|
| `metrics.enabled` | Enable metrics endpoint   | `false`       |
| `metrics.port`    | Port for exposing metrics | `8080`        |
| `metrics.secure`  | Use HTTPS for metrics     | `false`       |

#### Probe Configuration

| Parameter                             | Description                                 | Default Value |
|---------------------------------------|---------------------------------------------|---------------|
| `probe.enabled`                       | Enable health probes                        | `false`       |
| `probe.port`                          | Port for health probes                      | `8082`        |
| `probe.readiness.endpointName`        | Endpoint for readiness probe                | `readyz`      |
| `probe.readiness.initialDelaySeconds` | Initial delay before readiness probe starts | `10`          |
| `probe.readiness.periodSeconds`       | Interval between readiness probes           | `30`          |
| `probe.liveness.endpointName`         | Endpoint for liveness probe                 | `healthz`     |
| `probe.liveness.initialDelaySeconds`  | Initial delay before liveness probe starts  | `10`          |
| `probe.liveness.periodSeconds`        | Interval between liveness probes            | `30`          |

---

### Resource Management

#### Resource Requests and Limits

| Parameter   | Description                                              | Default Value |
|-------------|----------------------------------------------------------|---------------|
| `resources` | CPU/memory resource requests/limits for the operator pod | `{}`          |

#### Autoscaling

| Parameter                                       | Description                                  | Default Value |
|-------------------------------------------------|----------------------------------------------|---------------|
| `autoscaling.enabled`                           | Enable Horizontal Pod Autoscaler (HPA)       | `false`       |
| `autoscaling.minReplicas`                       | Minimum number of replicas for HPA           | `1`           |
| `autoscaling.maxReplicas`                       | Maximum number of replicas for HPA           | `100`         |
| `autoscaling.targetCPUUtilizationPercentage`    | Target CPU utilization percentage for HPA    | `80`          |
| `autoscaling.targetMemoryUtilizationPercentage` | Target memory utilization percentage for HPA | `""`          |

---

### Security Context

| Parameter            | Description                                 | Default Value |
|----------------------|---------------------------------------------|---------------|
| `podSecurityContext` | Security context for the operator pod       | `{}`          |
| `securityContext`    | Security context for the operator container | `{}`          |

---

### Logging and Debugging

| Parameter        | Description                       | Default Value |
|------------------|-----------------------------------|---------------|
| `logging.level`  | Log level (`info`, `debug`, etc.) | `info`        |
| `logging.format` | Log format (`json`, `text`, etc.) | `json`        |
| `pprof.enabled`  | Enable pprof debugging            | `false`       |
| `pprof.port`     | Port for pprof debugging          | `8081`        |

---

### RBAC and Permissions

| Parameter                    | Description                                             | Default Value |
|------------------------------|---------------------------------------------------------|---------------|
| `rbac.create`                | Create RBAC resources                                   | `true`        |
| `serviceAccount.create`      | Create a service account for the operator               | `true`        |
| `serviceAccount.annotations` | Annotations for the service account                     | `{}`          |
| `serviceAccount.name`        | Name of the service account (if not creating a new one) | `""`          |

---

### Leader Election

| Parameter                      | Description                        | Default Value                      |
|--------------------------------|------------------------------------|------------------------------------|
| `leaderElection.enabled`       | Enable leader election             | `false`                            |
| `leaderElection.id`            | ID for leader election             | `casdoor-operator-leader-election` |
| `leaderElection.namespace`     | Namespace for leader election      | `default`                          |
| `leaderElection.retryPeriod`   | Retry period for leader election   | `2s`                               |
| `leaderElection.renewDeadline` | Renew deadline for leader election | `10s`                              |
| `leaderElection.leaseDuration` | Lease duration for leader election | `15s`                              |

---

### Controller Configuration

#### Global Controller Settings

| Parameter                                  | Description                                   | Default Value |
|--------------------------------------------|-----------------------------------------------|---------------|
| `globalController.skipNameValidation`      | Skip name validation for resources            | `false`       |
| `globalController.groupKindConcurrency`    | Concurrency settings for group/kind resources | `{}`          |
| `globalController.maxConcurrentReconciles` | Maximum concurrent reconciles                 | `1`           |
| `globalController.cacheSyncTimeout`        | Timeout for cache synchronization             | `2m`          |
| `globalController.needLeaderElection`      | Require leader election for reconciliation    | `false`       |

#### Casdoor-Specific Controller Settings

| Parameter                                             | Description                                             | Default Value |
|-------------------------------------------------------|---------------------------------------------------------|---------------|
| `specificControllers.casdoor.skipNameValidation`      | Skip name validation for Casdoor resources              | `false`       |
| `specificControllers.casdoor.maxConcurrentReconciles` | Maximum concurrent reconciles for Casdoor resources     | `1`           |
| `specificControllers.casdoor.cacheSyncTimeout`        | Timeout for cache synchronization for Casdoor resources | `2m`          |
| `specificControllers.casdoor.needLeaderElection`      | Require leader election for Casdoor reconciliation      | `false`       |

---

### Advanced Configuration

| Parameter             | Description                                           | Default Value |
|-----------------------|-------------------------------------------------------|---------------|
| `extraEnvs`           | Additional environment variables for the operator pod | `[]`          |
| `extraVolumes`        | Additional volumes for the operator pod               | `[]`          |
| `extraVolumeMounts`   | Additional volume mounts for the operator pod         | `[]`          |
| `extraArgs`           | Additional arguments for the operator container       | `[]`          |
| `extraInitContainers` | Additional init containers for the operator pod       | `[]`          |
| `extraContainers`     | Additional sidecar containers for the operator pod    | `[]`          |
