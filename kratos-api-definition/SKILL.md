
---
name: kratos-api-definition
description: Use when defining APIs in a Kratos application using Protobuf.
---

# Kratos API Definition

This skill provides a guide to defining APIs in a Kratos application using Protobuf.

## Directory Structure

Protobuf files are located in the `api/` directory of the project.

```
api/
└── helloworld
    └── v1
        ├── greeter.proto
        └── ...
```

## Naming Conventions

- **Package**: `my.package.v1`
- **File**: `lower_snake_case.proto`
- **Message**: `CamelCase` (e.g., `MyRequest`)
- **Field**: `underscore_separated_names` (e.g., `my_field`)
- **Service**: `CamelCase` (e.g., `MyService`)
- **Method**: `CamelCase` (e.g., `MyMethod`)

## Example

```protobuf
syntax = "proto3";

package helloworld.v1;

import "google/api/annotations.proto";

option go_package = "github.com/go-kratos/kratos-layout/api/helloworld/v1;v1";

// The greeting service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply)  {
        option (google.api.http) = {
            get: "/helloworld/{name}"
        };
    }
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}
```

## Related Skills

- **[kratos-transport](./../kratos-transport/SKILL.md)**: Your API definitions are served over transports.
- **[kratos-openapi](./../kratos-openapi/SKILL.md)**: Generate OpenAPI specs from your Protobuf definitions.
