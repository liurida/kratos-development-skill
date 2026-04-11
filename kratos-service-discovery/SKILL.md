---
name: kratos-service-discovery
description: Use for registering a Kratos service with a registry like Consul or Nacos, and for discovering and connecting to other services.
---

# Kratos Service Discovery & Registration

## Overview

This skill provides a guide to Kratos's service discovery and registration mechanism. In a microservices architecture, services need a way to find each other. Kratos provides a unified `Registry` and `Discovery` interface to connect with various service registries like Consul, Nacos, and Etcd. This allows services to dynamically register themselves upon startup and discover the locations of other services they need to call.

## When to Use

Use this skill when you need to:
- Make your service discoverable by other services.
- Connect a Kratos client to a service without hardcoding its IP address.
- Implement a resilient architecture where service instances can be added or removed dynamically.
- Integrate with a service registry like Consul, Nacos, or Etcd.

## Core Pattern: Registration and Discovery

The pattern has two sides: the server registers itself, and the client discovers the server.

**1. Server-Side: Registration**

On the server side, you create a `Registrar` instance and pass it to your Kratos app. The app will then automatically register itself with the service registry on startup and de-register on shutdown.

```go
// cmd/server/main.go

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	consulAPI "github.com/hashicorp/consul/api"
)

func main() {
	// 1. Create a client for your chosen service registry (e.g., Consul).
	consulClient, err := consulAPI.NewClient(consulAPI.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// 2. Create a new Kratos registrar instance.
	reg := consul.New(consulClient)

	// 3. Pass the registrar to the Kratos app.
	app := kratos.New(
		kratos.Name("my-service"),
		kratos.Registrar(reg),
		// ... other app options
	)

	// ...
}
```

**2. Client-Side: Discovery**

On the client side, you create a `Discovery` instance and pass it to the gRPC client dialer. You use a special `discovery:///` scheme in the endpoint URI to tell the client to use the discovery mechanism.

```go
// somewhere in your client code

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	consulAPI "github.com/hashicorp/consul/api"
)

func main() {
	// 1. Create a client for the service registry.
	consulClient, err := consulAPI.NewClient(consulAPI.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// 2. Create a new Kratos discovery instance.
	discovery := consul.New(consulClient)

	// 3. Dial the service using the "discovery:///" scheme and the service name.
	conn, err := grpc.Dial(
		context.Background(),
		grpc.WithEndpoint("discovery:///my-service"),
		grpc.WithDiscovery(discovery),
	)
	// ...
}
```

## Available Implementations

Kratos provides implementations for many popular service registries in the `contrib/registry/` directory:
- `consul`
- `etcd`
- `kubernetes`
- `nacos`
- `polaris`
- `zookeeper`
- `eureka`

## Common Mistakes

- **Forgetting the `discovery:///` Scheme:** If you omit this scheme from the client endpoint, the gRPC client will try to connect to an address named "my-service" directly, which will fail. The scheme is what activates the discovery mechanism.
- **Mismatched Service Names:** The service name used in the client's discovery endpoint (`discovery:///my-service`) must exactly match the `kratos.Name()` provided by the server when it registered itself.
- **Firewall or Network Issues:** Ensure your client and server can both reach the service registry (e.g., the Consul server). Network policies can often block this communication.
- **Not Passing a Registrar:** If you create a Kratos app but don't pass it a `kratos.Registrar`, it will run but will not be discoverable by any other services.

## Related Skills

- **kratos-load-balancing**: Once the discovery component returns a list of available service instances, the load balancer is responsible for choosing which one to send the request to.
- **kratos-transport**: Both the server-side registration and client-side discovery are configured as part of the core Kratos application and gRPC client transport setup.
