# Casdoor Schema Reference

The Casdoor Custom Resource Definition (CRD) is designed to allow other resources to interact with the Casdoor API. It
also monitors the state of the Casdoor instance and updates its status accordingly.

**Example:**

```yaml title="casdoor.yaml"
apiVersion: casdoor.ptrvsrg.github.com/v1alpha1
kind: Casdoor
metadata:
  labels:
    app.kubernetes.io/name: casdoor-operator
  name: casdoor-sample
  namespace: default
spec:
  adminCredentialsSecret:
    name: casdoor-sample-admin-credentials
  healthcheck:
    enabled: true
  url: "http://casdoor-sample.default.svc.cluster.local:8000"
```

## Casdoor Schema Reference

| Property | Required | Property type                                    | Default | Description                                                                   |
|----------|----------|--------------------------------------------------|---------|-------------------------------------------------------------------------------|
| `spec`   | ✅        | [CasdoorSpec](#casdoorspec-schema-reference)     |         | The specification of the Casdoor instance, including configuration details.   |
| `status` | ✅        | [CasdoorStatus](#casdoorstatus-schema-reference) |         | The current status of the Casdoor instance, reflecting its operational state. |

## CasdoorSpec Schema Reference

Used in [Casdoor](#casdoor-schema-reference)

| Property                 | Required | Property type                                                                                                             | Default | Description                                                                                             |
|--------------------------|----------|---------------------------------------------------------------------------------------------------------------------------|---------|---------------------------------------------------------------------------------------------------------|
| `url`                    | ✅        | string                                                                                                                    | ""      | The URL of the Casdoor instance that will be accessed by other resources.                               |
| `adminCredentialsSecret` | ✅        | [LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#localobjectreference-v1-core) |         | A reference to the Kubernetes Secret containing the admin credentials for the Casdoor instance.         |
| `healthcheck`            |          | [CasdoorHealthcheckSpec](#casdoorhealthcheckspec-schema-reference)                                                        |         | Configuration for health checks to monitor the availability and responsiveness of the Casdoor instance. |

## CasdoorHealthcheckSpec Schema Reference

Used in [CasdoorSpec](#casdoorspec-schema-reference)

| Property   | Required | Property type | Default | Description                                                                                                                            |
|------------|----------|---------------|---------|----------------------------------------------------------------------------------------------------------------------------------------|
| `enabled`  |          | boolean       | false   | Enables or disables health checks for the Casdoor instance. When disabled, no periodic health checks will be performed.                |
| `method`   |          | string        | "GET"   | The HTTP method used for performing health checks. Supported methods include GET, HEAD, POST, etc.                                     |
| `path`     |          | string        | "/"     | The endpoint path on the Casdoor server that will be queried for health checks.                                                        |
| `interval` |          | string        | "1m"    | The time interval between consecutive health check requests. Specified as a duration (e.g., "1m" for one minute).                      |
| `timeout`  |          | string        | "1m"    | The maximum time allowed for a single health check request to complete. If the request exceeds this duration, it is considered failed. |
| `retries`  |          | integer       | 3       | The number of retry attempts for a failed health check before considering the Casdoor instance unhealthy.                              |

## CasdoorStatus Schema Reference

Used in [Casdoor](#casdoor-schema-reference)

| Property | Required | Property type                           | Default | Description                                                                                                                                                  |
|----------|----------|-----------------------------------------|---------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `code`   | ✅        | string (one or more of [Ready, Failed]) | ""      | The current status code of the Casdoor instance. Possible values include `Ready` (indicating the instance is operational) or `Failed` (indicating an issue). |
| `reason` |          | string                                  | ""      | A descriptive reason for the failure, if the status code is `Failed`. This provides additional context for troubleshooting.                                  |
