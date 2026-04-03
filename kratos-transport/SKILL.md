---
name: kratos-transport
description: Use when working with transport protocols (HTTP and gRPC) in a Kratos application.
---

# Kratos Transport

This skill provides a guide to working with transport protocols (HTTP and gRPC) in a Kratos application.

## Overview

Kratos provides a unified interface for working with different transport protocols. It has built-in support for HTTP and gRPC.

## HTTP

To create an HTTP server:

```go
import "github.com/go-kratos/kratos/v2/transport/http"

func main() {
	hs := http.NewServer(
		http.Address(":8000"),
	)
	// ...
}
```

## gRPC

To create a gRPC server:

```go
import "github.com/go-kratos/kratos/v2/transport/grpc"

func main() {
	gs := grpc.NewServer(
		grpc.Address(":9000"),
	)
	// ...
}
```

## Examples

- [**transport.go**](./examples/transport.go): The core interfaces for the Kratos transport layer.

- **[kratos-middleware](./../kratos-middleware/SKILL.md)**: Middleware is applied to transports to add functionality.
- **[kratos-api-definition](./../kratos-api-definition/SKILL.md)**: Your defined APIs are exposed over transports.
