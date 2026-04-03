
---
id: api-protobuf
title: Protobuf Guideline
---

This documentation is the guideline of Protobuf definition which recommended in Kratos project.

The API definition is based on HTTP and gRPC, written with Protobuf format They should includes all the Request, Reply and the corresponding Errors.

## Directory Structure

The definition of Proto could be either in `api` directory of the project or in a unified repository, likes `googleapis`, `envoy-api`, `istio-api`.

For the proto in project, the api should be used as the root of package name.

```
kratos-demo：
|____api // The definition of service's API
| |____kratos
| | |____demo
| | | |____v1
| | | | |____demo.proto
```

