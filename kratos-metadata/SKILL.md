---
name: kratos-metadata
description: Use for passing cross-cutting, request-scoped data like trace IDs or tenancy information between Kratos services.
---

# Kratos Metadata

## Overview

This skill provides a guide to Kratos's metadata package, which allows for the propagation of request-scoped information between services. Metadata is essentially a map of string key-value pairs that gets attached to a request's context and is automatically passed along via gRPC metadata or HTTP headers.

## When to Use

Use this skill when you need to:
- Pass a trace ID from a client to a server for distributed tracing.
- Propagate tenancy or user information across microservice calls.
- Send custom headers from a client that can be read by a server.
- Implement a middleware that reads or writes request-scoped data.

## Core Pattern: Client to Server Propagation

Metadata is set on a client context, and then retrieved from the server context. Kratos handles the transport automatically.

**1. On the Client Side:**
Attach metadata to the `context.Context` before making a gRPC or HTTP call.

```go
import (
	"context"
	"github.com/go-kratos/kratos/v2/metadata"
)

func main() {
	// Create a new metadata map.
	md := metadata.New(map[string]string{"x-trace-id": "abc-123", "x-user-id": "42"})

	// Attach it to a new client context.
	ctx := metadata.NewClientContext(context.Background(), md)

	// Make a request with this context.
	// Kratos will automatically convert the metadata to gRPC metadata or HTTP headers.
	greeterClient.SayHello(ctx, &pb.HelloRequest{Name: "world"})
}
```

**2. On the Server Side:**
Retrieve the metadata from the server context provided to your gRPC or HTTP handler.

```go
import "github.com/go-kratos/kratos/v2/metadata"

func (s *GreeterService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	// Retrieve the metadata from the incoming server context.
	if md, ok := metadata.FromServerContext(ctx); ok {
		traceID := md.Get("x-trace-id")
		userID := md.Get("x-user-id")

		log.Infof("Trace ID: %s, User ID: %s", traceID, userID)
	}

	// ... your service logic
	return &pb.HelloReply{Message: "Hello " + req.Name}, nil
}
```

## How it Works

- **gRPC:** Metadata is sent as gRPC metadata, which is efficient and standard.
- **HTTP:** Metadata is sent as HTTP headers. Keys are typically canonicalized (e.g., `x-trace-id` becomes `X-Trace-Id`).

## Common Mistakes

- **Using the Wrong Context Function:** Always use `metadata.NewClientContext` on the client side and `metadata.FromServerContext` on the server side. Using the wrong one will not work.
- **Assuming Case-Sensitivity:** HTTP header keys are case-insensitive by spec. While gRPC metadata is case-sensitive, it's best practice to use lowercase keys to avoid confusion when moving between protocols.
- **Overloading Metadata:** Metadata is for small, request-scoped data. Do not try to send large payloads or entire data objects in the metadata; that's what the request body is for.

## Related Skills

- **kratos-transport**: The transport layer is responsible for the actual transmission of metadata over HTTP headers or gRPC metadata.
- **kratos-middleware**: Middleware is the perfect place to read incoming metadata (like a trace ID) and attach it to the logger context for the duration of the request.
