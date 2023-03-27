# HTTP route matching rules

## Attributes

| attr          | superset                                                               |
|---------------|------------------------------------------------------------------------|
| hostname      | `^(\*\.)?[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9](?:\.[a-zA-Z]{2,})+$` |
| method        | `GET\|HEAD\|POST\|PUT\|DELETE\|CONNECT\|OPTIONS\|TRACE\|PATCH`         |
| path          | `^(\/\w*)*\*?$`                                                        |
| query_string  | `^(\?(\w+=[\w\d]+(&\w+=[\w\d]+)*)?\|\*)$`                              |
| headers       | `^((\w+:\s?[\w\d]+(,\w+:\s?[\w\d]+)*)?\|\*)$`                          |

## Example

### Rules

|  # | matcher                                                                      | policy     |
|---:|------------------------------------------------------------------------------|------------|
|  1 | **hostname**=`*.acme.com`                                                    | 500rps     |
|  2 | **hostname**=`*.acme.com`, **method=**`POST`                                 | def:100rps |
|  3 | **hostname**=`*.acme.com`                                                    | ip-deny    |
|  4 | **hostname**=`*.acme.internal`                                               | def:mtls   |
|  5 | **hostname**=`toys.acme.com,toys.acme.internal`                              | 150rps     |
|  6 | **hostname**=`toys.acme.com,toys.acme.internal`, **headers=**`X-Env: canary` | 50rps      |
|  7 | **hostname**=`toys.acme.internal`, **path=**`/admin/*`                       | ∞rps       |
|  8 | **hostname**=`toys.acme.com`                                                 | api-key    |
|  9 | **hostname**=`toys.acme.com`, **method=**`DELETE`                            | deny       |
| 10 | **hostname**=`toys.acme.com`, **path=**`/admin/*`                            | deny       |
| 11 | **hostname**=`toys.acme.com,toys.acme.internal`, **path=**`/orgs/*`          | check-org  |
| 12 | **hostname**=`toys.acme.internal`, **path=**`/admin/*`                       | k8s-authn  |
| 13 | **hostname**=`*.telemetry.acme.internal`                                     | oidc/jwt   |
| 14 | **hostname**=`foo.telemetry.acme.internal`                                   | +api-key   |

### Disjoint sets per attribute

| attr          | sets                                                                                                                               |
|---------------|------------------------------------------------------------------------------------------------------------------------------------|
| hostname      | `*.acme.com`, `*.acme.internal`, `toys.acme.com`, `toys.acme.internal`, `*.telemetry.acme.internal`, `foo.telemetry.acme.internal` |
| method        | `GET`, `POST`, `PUT`, `DELETE`                                                                                                     |
| path          | `/*`, `/admin/*`, `/orgs/*`                                                                                                        |
| query_string  | `*`                                                                                                                                |
| headers       | `*`, `X-Env: canary`                                                                                                               |

#### Hostnames

1. `toys.acme.com`
2. (`*.acme.com` - `toys.acme.com`)
3. `toys.acme.internal`
4. `foo.telemetry.acme.internal`
5. (`*.telemetry.acme.internal` - `foo.telemetry.acme.internal`)
6. (`*.acme.internal` - `toys.acme.internal` - `*.telemetry.acme.internal`)

#### Methods

1. `GET`
2. `POST`
3. `PUT`
4. `DELETE`

#### Paths

1. `/admin/*`
2. `/orgs/*`
3. (`/*` - `/admin/*` - `/orgs/*`)

#### Query string

1. `*`

#### Headers

1. `X-Env: canary`
2. (`*` - `X-Env: canary`)

### All the disjoint sets

**Number of disjoint sets:** 6 ⨉ 4 ⨉ 3 ⨉ 1 ⨉ 2 = 144<br/>
<sub>(#hostnames ⨉ #methods ⨉ #paths ⨉ #query_string ⨉ #headers)</sub>
