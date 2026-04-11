---
name: kratos-transport
description: Use for configuring and running HTTP and gRPC servers within a Kratos application, and for understanding the unified server lifecycle.
---

# Kratos Transport

## Overview

This skill provides a guide to the Kratos transport layer. Kratos abstracts server implementations (like HTTP and gRPC) behind a unified `transport.Server` interface. This allows a Kratos application to seamlessly manage the lifecycle (start, stop) of multiple servers at once.

## When to Use

Use this skill when you need to:
- Configure the address and timeout for your HTTP or gRPC servers.
- Add middleware to your HTTP or gRPC servers.
- Understand how to register your implemented services with the transport servers.
- Run both an HTTP and a gRPC server in the same application.

## Core Pattern: Creating and Running Servers

The standard pattern is to create your HTTP and/or gRPC server instances and then pass them to the main `kratos.App`.

```go
// cmd/server/main.go

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"path/to/your/service"
	"path/to/your/conf"
)

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.Name("my-service"),
		kratos.Server(
			// Register both your HTTP and gRPC servers with the app.
			hs,
			gs,
		),
		// ... other app options
	)
}

// This function will be part of your dependency injection setup.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService) *http.Server {
	var opts = []http.ServerOption{
		http.Address(c.Http.Addr),
		http.Middleware(
			recovery.Recovery(),
		),
	}
	// Create the HTTP server
	srv := http.NewServer(opts...)
	// Register your service implementation with the HTTP server.
	pb.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}

// This function will be part of your dependency injection setup.
func NewGRPCServer(c *conf.Server, greeter *service.GreeterService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Address(c.Grpc.Addr),
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	// Create the gRPC server
	srv := grpc.NewServer(opts...)
	// Register your service implementation with the gRPC server.
	pb.RegisterGreeterServer(srv, greeter)
	return srv
}
```

## How it Works

1.  Both `http.NewServer` and `grpc.NewServer` create servers that implement the `transport.Server` interface.
2.  You register your generated service handlers (`pb.RegisterGreeterHTTPServer`, `pb.RegisterGreeterServer`) with their respective servers.
3.  You pass both server instances to `kratos.New` via the `kratos.Server()` option.
4.  When `app.Run()` is called, the Kratos app will start both servers concurrently and manage their lifecycles, ensuring they both shut down gracefully.

## Common Mistakes

- **Forgetting to Register the Service:** If you create a server but forget to call `pb.Register...Server(srv, myService)`, the server will start, but it won't know about your service's RPC methods, resulting in `Not Found` or `Unimplemented` errors.
- **Mismatched Ports:** Ensure the addresses you configure for your HTTP and gRPC servers use different ports (e.g., `:8000` for HTTP, `:9000` for gRPC). Trying to use the same port will cause the application to fail at startup.
- **Applying Middleware in the Wrong Place:** Middleware should be applied using `http.Middleware()` or `grpc.Middleware()` when creating the server, not later.

## Related Skills

- **kratos-middleware**: Middleware is applied to transports to add functionality like logging, metrics, and recovery.
- **kratos-api-definition**: The `protoc-gen-go-http` and `protoc-gen-go-grpc` plugins generate the `Register...Server` functions that you use to attach your services to the transports.
- **kratos-service-discovery**: The transport servers, once started, are registered with the service registry so clients can find them.
