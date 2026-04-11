---
name: kratos-api-definition
description: Use when defining service APIs in a Kratos application using Protocol Buffers (Protobuf) for gRPC and HTTP.
---

# Kratos API Definition

## Overview

This skill provides a comprehensive guide to defining service APIs in a Kratos application using Protocol Buffers (Protobuf). In Kratos, Protobuf is the single source of truth for defining your service contracts, including RPC methods, message structures, and HTTP bindings.

## When to Use

Use this skill when you are:
- Creating a new service and need to define its API endpoints.
- Adding new RPC methods to an existing service.
- Defining request and response message structures.
- Mapping your gRPC services to RESTful HTTP endpoints using `google.api.annotations`.
- Unsure about the correct package structure and naming conventions for your `.proto` files.

### When NOT to Use
- For defining internal business logic; Protobufs should only define the public contract.
- When you are not using gRPC or HTTP as a primary transport.

## Quick Reference: Naming Conventions

| Element | Convention | Example |
|---|---|---|
| **Package** | `my.package.v1` | `helloworld.v1` |
| **File** | `lower_snake_case.proto` | `greeter_service.proto` |
| **Message** | `CamelCase` | `SayHelloRequest` |
| **Field** | `snake_case` | `user_name` |
| **Service** | `CamelCase` | `GreeterService` |
| **Method** | `CamelCase` | `SayHello` |

## Core Pattern: Defining a Service

All API definitions are located in the `api/` directory. The core components are the `service`, `message`, and `rpc` definitions, along with HTTP binding options.

```protobuf
// api/helloworld/v1/greeter.proto

syntax = "proto3";

package helloworld.v1;

import "google/api/annotations.proto";

// The go_package option specifies the import path for the generated Go code.
option go_package = "github.com/go-kratos/kratos-layout/api/helloworld/v1;v1";

// 1. Define the service
service Greeter {
  // 2. Define the RPC method, request, and response
  rpc SayHello (HelloRequest) returns (HelloReply)  {
    // 3. Define the HTTP binding
    option (google.api.http) = {
      get: "/helloworld/{name}"
    };
  }
}

// 4. Define the request message
message HelloRequest {
  // Field numbers 1-15 use 1 byte encoding, use them for frequently used fields.
  string name = 1;
}

// 5. Define the response message
message HelloReply {
  string message = 1;
}
```

## Common Mistakes

- **Incorrect `go_package` Path:** Ensure the `go_package` option correctly points to your project's module path. An incorrect path will cause import errors in your Go code.
- **Forgetting `google/api/annotations.proto`:** If you are defining HTTP bindings, you **must** import `google/api/annotations.proto`. Forgetting this will cause `protoc` to fail.
- **Using Relative Paths in Imports:** Always use fully qualified paths for imports within your `.proto` files if they are in different packages.
- **Manual JSON Serialization:** Do not manually create JSON structs that mirror your Protobuf messages. Kratos handles the gRPC to JSON transcoding automatically based on your HTTP bindings.

## Related Skills

- **kratos-transport**: Your API definitions are served over transports.
- **kratos-openapi**: Generate OpenAPI/Swagger specs from your Protobuf definitions.
- **kratos-cli**: Use the `kratos proto` commands to generate server and client stubs from your definitions.
