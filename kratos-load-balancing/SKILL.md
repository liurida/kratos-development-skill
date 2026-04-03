---
name: kratos-load-balancing
description: Use when working with load balancing in a Kratos application.
---

# Kratos Load Balancing

This skill provides a guide to working with load balancing in a Kratos application.

## Overview

Kratos has several built-in load balancing algorithms that can be used when making client requests.

## Usage

Load balancing is configured on the client side when creating a new gRPC connection.

```go
import "github.com/go-kratos/kratos/v2/transport/grpc"

func main() {
	conn, err := grpc.Dial(
		context.Background(),
		grpc.WithEndpoint("discovery:///my-service"),
		grpc.WithDiscovery(d),
		grpc.WithBalancer("round_robin"),
	)
	// ...
}
```

## Examples

- [**selector.go**](./examples/selector.go): The core interfaces for the Kratos load balancing and node selection.

- `round_robin` (default)
- `p2c`
- `random`

## Related Skills

- **[kratos-service-discovery](./../kratos-service-discovery/SKILL.md)**: The load balancer uses the list of services from the discovery component.
