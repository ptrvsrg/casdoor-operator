# Casdoor CRD

The Casdoor Custom Resource Definition (CRD) is designed to allow other resources to interact with the Casdoor API. It
also monitors the state of the Casdoor instance and updates its status accordingly.

**Example:**

```yaml title="casdoor-sample-secret.yaml"
apiVersion: v1
kind: Secret
metadata:
  name: casdoor-test-secret
  labels:
    app.kubernetes.io/name: casdoor-operator
    app.kubernetes.io/instance: test
type: Opaque
data:
  clientSecret: MDRjNDFjODNhNGEzYTY1ZjI5MmQxZjZiMTM4MjEzZTc0YmE1Mzk3Mwo=
  jwtCertificate: |-
    LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUUyVENDQXNHZ0F3SUJBZ0lEQWVKQU1BMEdDU3FHU0liM0RRRUJDd1VBTUNZeERqQU1CZ05WQkFvVEJXRmsKYldsdU1SUXdFZ1lEVlFRRERBdGpaWEowWDJoMk5HZGlaekFlRncweU5UQTFNekV4T0RVeE5EVmFGdzAwTlRBMQpNekV4T0RVeE5EVmFNQ1l4RGpBTUJnTlZCQW9UQldGa2JXbHVNUlF3RWdZRFZRUUREQXRqWlhKMFgyaDJOR2RpClp6Q0NBaUl3RFFZSktvWklodmNOQVFFQkJRQURnZ0lQQURDQ0Fnb0NnZ0lCQU1mRGRQZmRnRWlleVY4dEs0M2gKakc0NHgwOFpHVGFBMm1MaHpwTzZyY0xQZGlpbmsrSDNxYlJadEpBRGIrMG9YNWRoVTNlSlZlejVzWWt4Ry9FTApVTXdjakxteXNLUHVSa3pXWkFjNUtiSGNxdWQvbnJqS1dJWFlJQTFSdDhrTzdIZ2JxRFJXWlhMUkh4enc5dFBhCm0yWHppMXY1b3RGcU9iWGVPQVlNVlcrMVp6Ym5oaHMrVFA4SW51NzNxT2V6c3c4ZGhPaTJWVGtrc3R6Q2JKL2EKQUsrZDIrbGRPZ0hqaXJFYnltY2JGZXY3WEdVTUtoRDR4YmViNG13RnlEeHliYkNTU3BRQzR6NWQyd1Q5V0pzWQpTSUtIaU5Ec25IYmNjK3hCY2JJaENyOFRJOU9hV25lYUZpOWpqREFGZ3VFUCtvV2pQVml5L0JTYXU1ZTRYbTd2CkY0UE94bVlOMUl6ai9QNDlHcVBLamk1aUE0aGVZWU5HdzAxOWd4dnRXa21TZ0o0emRuRlZXUVpkZ1R5SzI5emIKVXlqNzlJVVoxNVZ6MTJuSEdxZkJpUUVZclhsd1dueDdiY1owdWxiYTdpQTNOOUpzcVY3MWlKVjVhd2FGOGFYbAppNDRadE5udituVzg1bmRrYU9PSDUyRDdaWm50bjhYVjRCaXU2SGt5MDRlSHdWMy9GZkY1VWRiMW54RWU2MEpVCkxZRzZsMWV1cXVnNXVkdGhRWWlRRXRvVTRZRjdmdWpIOUdBUUczN1lFNko2UTFKQ0lNRlpIQ3hHSHpoZjJQTU8KT2w2d0pZNFdPYVhrVFdGU3FzK2NQTnZFQ0VlMC9iVWJPMjZZUmtzYTB1eU9rYStyK0tINS9rUmk4a3VuQ2dScQpqWGo4aW5IcUNpN29Uc2dDZGVJc2EyOTNBZ01CQUFHakVEQU9NQXdHQTFVZEV3RUIvd1FDTUFBd0RRWUpLb1pJCmh2Y05BUUVMQlFBRGdnSUJBSGFNOHVKS2JobFcySFNPVWVkREhFejBMUGd6eHdXWU81MXJFMWVUc0UzMytXekMKQWUrS0FwMHJlc1pNN24xN1F5Y3FIMGdGcHk3Nng1K3ZBTFNWbXJrS2c5UFc5YTl0NlQ5R3VScnRRNklaTWdjMQpjTjQvS1BGRm1ER2ZVTFJJbUdmUkRKMFFSaC9HblAvbERQMFN6MEU1TTRydkJPL294eEJwcTN1MXViMHVsWk03CkZTWG56UDdqeXlzUWtSVjd6a05pVmthUklUbnhrNWg3ei9vUFNvMHJFaGhqVU1UY2YyVEcxZVBMTGw0b1ZwaUsKNFdvK2xVL3FPRnFGMXhyNXFxV3N3OERib2toNzh4MDExV2syejV5a2duWUo4b2ZOcStRZ0p4a1NtWGpNMUQwZwo0My85SytOcjdJajhiNHNmTjZ1S1NlSWU0N00weXJvRUsyTzJOeE4zOEFxRGdjY2tPUTNickFiVjJ0S3dIdVZ6CnNaQ3pTVTNWNXNVWTBsNTdvbHNrKzdUSE94MlVoY2ZNWlJyNThCcWprMXI5cFptSnJXamZZYzRBNncxNUNJL2MKenNhK0tmaEcxZ2kwYndBbzJkQW9PMVpweTVseEdQUkJhc1R6c1FtZDZMQ2dkaVdGNnI3ZEFzakg5WnUxZXNQWgpyeitxM0UwSU1aNjZwQzFZc2NacWdmb0wxYmlsb25xSHZWQXRseEIxM0h5YzFGaXE4YUxLQ0pRNGEyc0dsczVICjByU2lYWDAxd2V0VEE4eWlQMHhJZnRDZWJPL3lZeDBzWUN1dEJ2THM3K0E2Z0IxUEk3YlcyQytMeXY4VTE5dUcKakh1WlJPbVFoNnZFNGFHYlQybmorZjlTbzNENXFqM3Y1eDBOUkd2a0JuZ3ZwTnRRRjFBR05FdzRDRzFQCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
```

```yaml title="casdoor-sample.yaml"
apiVersion: casdoor.ptrvsrg.github.com/v1alpha1
kind: Casdoor
metadata:
  labels:
    app.kubernetes.io/name: casdoor-operator
    app.kubernetes.io/instance: test
  name: casdoor-test
spec:
  url: http://test-casdoor.default.svc.cluster.local:8000
  organizationName: platform-admin-org
  applicationName: platform-admin-app
  clientID: 9bb2f3bdc6238f575900
  clientSecret:
    key: clientSecret
    name: casdoor-test-secret
  jwtCertificate:
    key: jwtCertificate
    name: casdoor-test-secret
  healthcheck:
    enabled: true
```

## Casdoor Schema Reference

| Property | Required | Property type                                    | Default | Description                                                                   |
|----------|----------|--------------------------------------------------|---------|-------------------------------------------------------------------------------|
| `spec`   | ✅        | [CasdoorSpec](#casdoorspec-schema-reference)     |         | The specification of the Casdoor instance, including configuration details.   |
| `status` |          | [CasdoorStatus](#casdoorstatus-schema-reference) |         | The current status of the Casdoor instance, reflecting its operational state. |

## CasdoorSpec Schema Reference

Used in [Casdoor](#casdoor-schema-reference)

| Property           | Required | Property type                                                                                                       | Default | Description                                                                                                                                     |
|--------------------|----------|---------------------------------------------------------------------------------------------------------------------|---------|-------------------------------------------------------------------------------------------------------------------------------------------------|
| `url`              | ✅        | string                                                                                                              | ""      | The URL of the Casdoor instance that will be accessed by other resources.                                                                       |
| `organizationName` | ✅        | string                                                                                                              | ""      | The name of Casdoor organization for management other Casdoor resources.                                                                        |
| `applicationName`  | ✅        | string                                                                                                              | ""      | The name of Casdoor application for management other Casdoor resources.                                                                         |
| `clientID`         | ✅        | string                                                                                                              | ""      | The client ID of Casdoor application for management other Casdoor resources.                                                                    |
| `clientSecret`     | ✅        | [SecretKeySelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#secretkeyselector-v1-core) |         | The reference to the key from a Kubernetes Secret containing the client secret of Casdoor application for management other Casdoor resources.   |
| `jwtCertificate`   | ✅        | [SecretKeySelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#secretkeyselector-v1-core) |         | The reference to the key from a Kubernetes Secret containing the JWT certificate of Casdoor application for management other Casdoor resources. |
| `healthcheck`      |          | [CasdoorHealthcheckSpec](#casdoorhealthcheckspec-schema-reference)                                                  |         | Configuration for health checks to monitor the availability and responsiveness of the Casdoor instance.                                         |

## CasdoorHealthcheckSpec Schema Reference

Used in [CasdoorSpec](#casdoorspec-schema-reference)

| Property  | Required | Property type | Default | Description                                                                                                                            |
|-----------|----------|---------------|---------|----------------------------------------------------------------------------------------------------------------------------------------|
| `enabled` |          | boolean       | false   | Enables or disables health checks for the Casdoor instance. When disabled, no periodic health checks will be performed.                |
| `method`  |          | string        | "GET"   | The HTTP method used for performing health checks. Supported methods include GET, HEAD, POST, etc.                                     |
| `path`    |          | string        | "/"     | The endpoint path on the Casdoor server that will be queried for health checks.                                                        |
| `timeout` |          | string        | "1m"    | The maximum time allowed for a single health check request to complete. If the request exceeds this duration, it is considered failed. |
| `retries` |          | integer       | 3       | The number of retry attempts for a failed health check before considering the Casdoor instance unhealthy.                              |

## CasdoorStatus Schema Reference

Used in [Casdoor](#casdoor-schema-reference)

| Property | Required | Property type                                      | Default   | Description                                                                                                                   |
|----------|----------|----------------------------------------------------|-----------|-------------------------------------------------------------------------------------------------------------------------------|
| `code`   | ✅        | string (one or more of [Ready, NotReady, Unknown]) | "Unknown" | The current status code of the Casdoor instance. Possible values include `Ready`, `NotReady` or `Unknown`.                    |
| `reason` |          | string                                             | ""        | A descriptive reason for the failure, if the status code is `NotReady`. This provides additional context for troubleshooting. |