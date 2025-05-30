# Casdoor Schema Reference

The Casdoor Custom Resource Definition (CRD) is designed to allow other resources to interact with the Casdoor API. It
also monitors the state of the Casdoor instance and updates its status accordingly.

**Example:**

```yaml title="casdoor-sample-secret.yaml"
apiVersion: v1
kind: Secret
metadata:
  name: casdoor-sample-secret
  labels:
    app.kubernetes.io/name: casdoor-operator
type: Opaque
data:
  clientSecret: MzhjMDA1ODMyYjQ4NDY4Zjg0NmU5NGQ0YWQ3YTdiYzRkZjM5YmM5Ygo=
  jwtCertificate: |
    LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUU5VENDQXQyZ0F3SUJBZ0lEQWVKQU1BMEdDU3FHU0liM0RRRUJDd1VBTURReEhEQWFCZ05WQkFvTUUyOXkKWjJGdWFYcGhkR2x2Ymw5cVpXOXRZV3d4RkRBU0JnTlZCQU1NQzJObGNuUmZOV3c1Y0RBM01CNFhEVEkxTURVeQpOekUyTURRd05Wb1hEVFExTURVeU56RTJNRFF3TlZvd05ERWNNQm9HQTFVRUNnd1RiM0puWVc1cGVtRjBhVzl1ClgycGxiMjFoYkRFVU1CSUdBMVVFQXd3TFkyVnlkRjgxYkRsd01EY3dnZ0lpTUEwR0NTcUdTSWIzRFFFQkFRVUEKQTRJQ0R3QXdnZ0lLQW9JQ0FRRE1JalNNVklXZ2hmaVVkWlpGUDdLOUc5dHRrUTNaVUMvaXFMaVZkZlVMNzBqeAo0NFVTYWovbFVFMkF5cGZuRFBxdTNxbmo2VDYrNUpERTIxY1QwVEZwUkJOQjRwSVQ3Z2dxS1RBQitON2loM2NPCnBVRjVyS2dLRndCeTRSaWsrM3oydFNtMjhWQ2I4WFNGYnM5WERYUXhYN3dvOVBsb0NPckpNSkhzM00wMFlpclkKcFVFaTJRMWlmMGJIL25yY283WTRFUUp0N1l6SEZsSlIrYTFJcTcrL01xakFUWWlSSUtEZHYrMFh0dGQwQkN0KwpCalZJMnB0OVBiekdwSnJ4TVRoVWZKTUxoUExXRUI5YTdFS084SS9TTTRXRldZWlc4WWtld25JMlUyS0NlUjhpCkhBbkhxUnhGWDFXc0xXUTV2MHZHekllNlZtajFUSlIxQTQ3d3MzeVFWaVVucExONXNPcS9FME5BWWVUc0doQkYKdUQyVTBhVVpoN0pHQnFQUWt6MDZGZTNoQjVINzN1L1NwUmFtWmREVGpCTHY2OG4zakpaa3RLTG5YQ2twMUxZbgp2V3d2Y3RveXpLbitubnRmeEs2Q1RrbGt5MGJxQ2hIVGF0VmFyYk5tSUtrdzhUNHFUQUNIRW1aaVNhdGhadksxCmhrM3RVMG9LdWVXMFNPcTlCWTlzVkljTkxzRG8xY1JENW5hdldMUnJUTnc3bmZoM0VkNGNtUHNEZjFWZU1MYWgKYngzSlhSR0JVajRmTHYzU3N2aUUyN3JDbWlBazNycmhmSURJQmxDOWxUMitpWEhHRGxGYldPbTVkS29LaUg3SApIbjBQd2JvaWZscVova1MwV3c2b3ZVbXdZVTZuQWIxVFVhNkJGbm1KYkJydy9GOWZJTHhhRGJhNEZnVEhEUUlECkFRQUJveEF3RGpBTUJnTlZIUk1CQWY4RUFqQUFNQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUNBUUFwL3FDNE80M1IKaVBnOTNGME9ZMURUYXZ1dkFoSnpDZ2Q0dWdrNnlrZmJsZTdVTDFaYzBmZllJcTZOUmV3VENrVWFaMkVPTkxYawp3TkI4SW1XOWI4bEVEV2RQOUJpODRhSWYrTWFBMzQybmdsNGVoenAzTzJIckd1ZUNmZWNkMjVrN2FkV25SMDFPCnpFbHVzSWQ3dFRseWhQdzlSZVA1ZjljZndXemM0K1JEWVZJZElXRVdOZkY3NXoxUTkvZjNXL0VqVjNSTkQ5anAKVzdSbEo4WTNQSnJIVFR2QkZlOGIwMEdtejJoTjZkYWw2VTM5NDB5TmNxdVBGMVBFNjFaSDhiQXBucFFna21VeAp5Ky9xRmcyU1JUUVROVVBRRDhXSi9ib2FTdXVXbzloZXVTWjZRUnBYM0kwSS9KdjcvODhqdFZDaVlueE1lL1o2ClpXdnhJOXNXU2g1eVFzT1pBVzNCNkxTdDgwNG1NK28yaVVZd1hrclFzM2o5OWZWS1NMTGhRelpzSjhwS0ZsN2MKQllvSXVVaVpzTnBUQ0piTmVIMVF0UDFNMTFWZ0k4ZU1uTGxsQ1BqZDZlNXpRUmNQWVM1bVdWNEhrTHk4dnhDMApNSVNYRU5HNjlFa0EweWN6YXRoakhoOWpXcGJLYk5ydTRkOUttUXFENG1HZ0YxVUhJTkhIakJjYlZNVWxYL2MwCnZlZERyQmdSM293VWJKVHNjRzIzd3ZzVkh4MEZvRXA4THNvQlJubmlDMDJhanJYT0NEMDI0L3FnQUZmeUtZMFkKSGlYWTB3Y0w1dzBmT0ZyNHRmbEwyWTZLY1NOME0vUCtRM3V6aklrTmxYVGFreWdtaEdkM3lDcFZjNDA2bVpZeApVY2xDeHJndk80dXpCZk5JbFFuZk00b3UxRHdxZ3ZYNFhnPT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
```

```yaml title="casdoor-sample.yaml"
apiVersion: casdoor.ptrvsrg.github.com/v1alpha1
kind: Casdoor
metadata:
  labels:
    app.kubernetes.io/name: casdoor-operator
  name: casdoor-sample
spec:
  url: http://casdoor-sample.default.svc.cluster.local:8000
  organizationName: test
  applicationName: test
  clientID: dc3d4ebc690f0eb4f427
  clientSecret:
    key: clientSecret
    name: casdoor-sample-secret
  jwtCertificate:
    key: jwtCertificate
    name: casdoor-sample-secret
```

## Casdoor Schema Reference

| Property | Required | Property type                                | Default | Description                                                                 |
|----------|----------|----------------------------------------------|---------|-----------------------------------------------------------------------------|
| `spec`   | ✅        | [CasdoorSpec](#casdoorspec-schema-reference) |         | The specification of the Casdoor instance, including configuration details. |

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
