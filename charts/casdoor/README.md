# Casdoor

Casdoor is an open-source identity and access management platform that provides authentication, authorization, and user
management capabilities.

## Parameters

The following tables list the configurable parameters of the Casdoor Helm chart and their default values.

---

### Global Configuration

| Parameter              | Description                                   | Default Value |
|------------------------|-----------------------------------------------|---------------|
| `replicaCount`         | Number of replicas for the Casdoor deployment | `1`           |
| `nameOverride`         | Override the name of the chart                | `""`          |
| `fullnameOverride`     | Override the full name of the chart           | `""`          |
| `revisionHistoryLimit` | Number of old ReplicaSets to retain           | `10`          |

---

### Image Configuration

| Parameter          | Description                                                      | Default Value  |
|--------------------|------------------------------------------------------------------|----------------|
| `image.registry`   | Docker registry for the Casdoor image                            | `docker.io`    |
| `image.repository` | Repository for the Casdoor image                                 | `casbin`       |
| `image.name`       | Name of the Casdoor image                                        | `casdoor`      |
| `image.tag`        | Tag for the Casdoor image (defaults to the Chart's `appVersion`) | `""`           |
| `image.pullPolicy` | Image pull policy                                                | `IfNotPresent` |
| `imagePullSecrets` | Secrets for pulling images from private registries               | `[]`           |

---

### Service Configuration

| Parameter                | Description                                                   | Default Value |
|--------------------------|---------------------------------------------------------------|---------------|
| `service.type`           | Type of service (`ClusterIP`, `NodePort`, `LoadBalancer`)     | `ClusterIP`   |
| `service.httpPort`       | Port for HTTP traffic                                         | `8000`        |
| `service.ldapPort`       | Port for LDAP traffic                                         | `8001`        |
| `service.ldapsPort`      | Port for LDAPS traffic                                        | `8002`        |
| `service.radiusPort`     | Port for RADIUS traffic                                       | `8003`        |
| `service.httpNodePort`   | NodePort for HTTP traffic (if `service.type` is `NodePort`)   | `30000`       |
| `service.ldapNodePort`   | NodePort for LDAP traffic (if `service.type` is `NodePort`)   | `30001`       |
| `service.ldapsNodePort`  | NodePort for LDAPS traffic (if `service.type` is `NodePort`)  | `30002`       |
| `service.radiusNodePort` | NodePort for RADIUS traffic (if `service.type` is `NodePort`) | `30003`       |

---

### Resource Management

#### Resource Requests and Limits

| Parameter   | Description                                             | Default Value |
|-------------|---------------------------------------------------------|---------------|
| `resources` | CPU/memory resource requests/limits for the Casdoor pod | `{}`          |

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

| Parameter            | Description                                | Default Value |
|----------------------|--------------------------------------------|---------------|
| `podSecurityContext` | Security context for the Casdoor pod       | `{}`          |
| `securityContext`    | Security context for the Casdoor container | `{}`          |

---

### Probes

| Parameter                 | Description                  | Default Value |
|---------------------------|------------------------------|---------------|
| `probe.enabled`           | Enable health probes         | `true`        |
| `probe.readinessEndpoint` | Endpoint for readiness probe | `/api/health` |
| `probe.livenessEndpoint`  | Endpoint for liveness probe  | `/`           |

---

### Ingress Configuration

| Parameter             | Description                                | Default Value         |
|-----------------------|--------------------------------------------|-----------------------|
| `ingress.enabled`     | Enable ingress resource                    | `false`               |
| `ingress.className`   | Ingress class name                         | `""`                  |
| `ingress.annotations` | Annotations for the ingress resource       | `{}`                  |
| `ingress.hosts`       | Hosts and paths for the ingress resource   | `chart-example.local` |
| `ingress.tls`         | TLS configuration for the ingress resource | `[]`                  |

---

### Database Configuration

| Parameter                  | Description                                             | Default Value |
|----------------------------|---------------------------------------------------------|---------------|
| `database.driver`          | Database driver (`sqlite`, `mysql`, `postgresql`, etc.) | `sqlite`      |
| `database.tableNamePrefix` | Prefix for database table names                         | `""`          |
| `database.dbName`          | Name of the database                                    | `casdoor`     |
| `database.options`         | Additional database options                             | `{}`          |
| `database.host`            | Database host (for MySQL, PostgreSQL, etc.)             | `""`          |
| `database.port`            | Database port (for MySQL, PostgreSQL, etc.)             | `""`          |
| `database.username`        | Database username                                       | `""`          |
| `database.password`        | Database password                                       | `""`          |

---

### Advanced Features

#### Redis

| Parameter        | Description              | Default Value |
|------------------|--------------------------|---------------|
| `redis.enabled`  | Enable Redis integration | `false`       |
| `redis.endpoint` | Redis endpoint           | `""`          |

#### LDAP

| Parameter         | Description        | Default Value |
|-------------------|--------------------|---------------|
| `ldap.enabled`    | Enable LDAP server | `false`       |
| `ldap.serverPort` | LDAP server port   | `389`         |

#### LDAPS

| Parameter          | Description              | Default Value |
|--------------------|--------------------------|---------------|
| `ldaps.enabled`    | Enable LDAPS server      | `false`       |
| `ldaps.certId`     | Certificate ID for LDAPS | `""`          |
| `ldaps.serverPort` | LDAPS server port        | `636`         |

#### RADIUS

| Parameter                    | Description                     | Default Value |
|------------------------------|---------------------------------|---------------|
| `radius.enabled`             | Enable RADIUS server            | `false`       |
| `radius.serverPort`          | RADIUS server port              | `1812`        |
| `radius.defaultOrganization` | Default organization for RADIUS | `"built-in"`  |
| `radius.secret`              | RADIUS shared secret            | `"secret"`    |

---

### Logging

| Parameter         | Description                           | Default Value        |
|-------------------|---------------------------------------|----------------------|
| `log.showSql`     | Show SQL logs                         | `false`              |
| `log.logPostOnly` | Log only POST requests                | `true`               |
| `log.adapter`     | Log adapter (`file`, `console`, etc.) | `"file"`             |
| `log.filename`    | Log file path                         | `"logs/casdoor.log"` |
| `log.maxDays`     | Maximum days to retain logs           | `30`                 |
| `log.perm`        | File permissions for log files        | `"0770"`             |

---

### Application Configuration

| Parameter                     | Description                                 | Default Value               |
|-------------------------------|---------------------------------------------|-----------------------------|
| `app.name`                    | Application name                            | `"casdoor"`                 |
| `app.mode`                    | Application mode (`dev`, `prod`, etc.)      | `"dev"`                     |
| `app.httpPort`                | HTTP port for the application               | `8000`                      |
| `app.SessionOn`               | Enable session management                   | `true`                      |
| `app.copyRequestBody`         | Copy request body for logging               | `true`                      |
| `app.isUsernameLowered`       | Force lowercase usernames                   | `false`                     |
| `app.staticBaseUrl`           | Base URL for static resources               | `"https://cdn.casbin.org "` |
| `app.batchSize`               | Batch size for processing                   | `100`                       |
| `app.enableErrorMask`         | Mask error messages                         | `false`                     |
| `app.enableGzip`              | Enable Gzip compression                     | `true`                      |
| `app.inactiveTimeoutMinutes`  | Timeout for inactive sessions (in minutes)  | `1440`                      |
| `app.verificationCodeTimeout` | Timeout for verification codes (in minutes) | `10`                        |
| `app.initScore`               | Initial score for new users                 | `0`                         |
| `app.origin`                  | Origin URL for the application              | `""`                        |
| `app.originFrontend`          | Frontend origin URL                         | `""`                        |

---

### Quota Configuration

| Parameter            | Description                                        | Default Value |
|----------------------|----------------------------------------------------|---------------|
| `quota.organization` | Maximum number of organizations (-1 for unlimited) | `-1`          |
| `quota.user`         | Maximum number of users (-1 for unlimited)         | `-1`          |
| `quota.application`  | Maximum number of applications (-1 for unlimited)  | `-1`          |
| `quota.provider`     | Maximum number of providers (-1 for unlimited)     | `-1`          |

---

### Advanced Configuration

| Parameter           | Description                                   | Default Value |
|---------------------|-----------------------------------------------|---------------|
| `initContainers`    | Additional init containers                    | `[]`          |
| `extraContainers`   | Additional sidecar containers                 | `[]`          |
| `extraVolumeMounts` | Additional volume mounts                      | `[]`          |
| `extraVolumes`      | Additional volumes                            | `[]`          |
| `extraEnvs`         | Additional environment variables              | `[]`          |
| `extraEnvsFrom`     | Additional environment variables from sources | `[]`          |

---