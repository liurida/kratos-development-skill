---
name: kratos-load-balancing
description: Use when configuring client-side load balancing policies like round-robin, P2C, or random for gRPC clients in a Kratos application.
---

# Kratos Client-Side Load Balancing

## Overview

This skill provides a guide to Kratos's client-side load balancing capabilities. When a client needs to communicate with a service that has multiple instances, the Kratos client can distribute requests across those instances using a configured load balancing policy. This is configured on the gRPC client at connection time.

## When to Use

Use this skill when you need to:
- Configure how a gRPC client distributes requests among multiple service instances.
- Choose a load balancing strategy that fits your application's needs (e.g., simple distribution vs. latency-aware).
- Understand the different load balancing algorithms available in Kratos.

## Quick Reference: Built-in Balancers

| Name | Description | When to Use |
|---|---|---|
| `round_robin` | (Default) Rotates through the list of available nodes sequentially. | Good for evenly distributing load across healthy, homogenous services. |
| `p2c` | (Power of Two Choices) Randomly picks two nodes and chooses the one with the least load. | Excellent for heterogeneous environments or when services can become overloaded. It naturally favors healthier/faster nodes. |
| `random` | Picks a node at random. | Simple, but less effective at even distribution than round-robin or P2C. |

## Core Pattern: Configuring a Balancer

Load balancing is configured on the client side when creating a gRPC connection using `grpc.Dial`. You must also provide a discovery mechanism to supply the list of available service instances.

```go
import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/registry"

	// Import your chosen discovery implementation
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
)

func main() {
	// 1. Set up your service discovery component.
	// This example uses Consul, but it could be etcd, nacos, etc.
	consulClient, err := consul.NewClient(consulapi.DefaultConfig())
	if err != nil {
		panic(err)
	}
	discovery := consul.New(consulClient)

	// 2. Dial the service using a discovery endpoint.
	// The "discovery:///" scheme tells the client to use the service discovery system.
	conn, err := grpc.Dial(
		context.Background(),
		// The service to connect to.
		grpc.WithEndpoint("discovery:///my-service"),
		// Provide the discovery component.
		grpc.WithDiscovery(discovery),
		// 3. Specify the load balancing policy.
		grpc.WithBalancer("p2c"), // Use Power of Two Choices
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// ... use the connection to make RPC calls
}
```

## Common Mistakes

- **Forgetting `WithDiscovery`:** The load balancer has no nodes to choose from if you don't provide a discovery mechanism via `grpc.WithDiscovery`. The client will be unable to connect.
- **Using a Direct Endpoint:** If you provide a direct IP address endpoint (e.g., `grpc.WithEndpoint("127.0.0.1:9000")`) instead of a discovery endpoint (`"discovery:///..."`), load balancing will not be activated because there is only one node to connect to.
- **Choosing the Wrong Policy:** While `round_robin` is a safe default, `p2c` is often a better choice in production environments as it can adapt to nodes that are slower or temporarily overloaded, leading to better overall latency.

## Related Skills

- **kratos-service-discovery**: The load balancer relies on the service discovery system to get the list of available service instances to balance between.
- **kratos-transport**: Load balancing is configured as part of the client-side gRPC transport setup.
