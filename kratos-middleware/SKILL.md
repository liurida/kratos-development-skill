
---
name: kratos-middleware
description: Use when working with middleware in a Kratos application.
---

# Kratos Middleware

This skill provides a guide to working with middleware in a Kratos application.

## Overview

Middleware in Kratos is used to add cross-cutting concerns to requests, such as logging, tracing, and recovery. Middleware is executed in a "first in, last out" (FILO) order.

## Usage

Middleware is registered when creating a new HTTP or gRPC server.

```go
// http
var opts = []http.ServerOption{
    http.Middleware(
        recovery.Recovery(),
        tracing.Server(),
        logging.Server(),
    ),
}
http.NewServer(opts...)

//grpc
var opts = []grpc.ServerOption{
    grpc.Middleware(
        recovery.Recovery(),
        status.Server(),
        tracing.Server(),
        logging.Server(),
    ),
}
grpc.NewServer(opts...)
```

## Custom Middleware

You can create custom middleware by implementing the `middleware.Middleware` interface.

```go
func MyMiddleware() middleware.Middleware {
    return func(handler middleware.Handler) middleware.Handler {
        return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
            // Pre-processing
            defer func() { /* Post-processing */ }()
            return handler(ctx, req)
        }
    }
}
```

## Related Skills

- **[kratos-transport](./../kratos-transport/SKILL.md)**: Middleware is applied to transports (HTTP and gRPC).
- **[kratos-logging](./../kratos-logging/SKILL.md)**: The logging middleware is a common use case.
