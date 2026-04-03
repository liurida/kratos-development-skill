---
name: kratos-service-discovery
description: Use when working with service discovery and registration in a Kratos application.
---

# Kratos Service Discovery

This skill provides a guide to working with service discovery and registration in a Kratos application.

## Overview

Kratos provides a unified interface for service registration and discovery, allowing you to easily connect to various service registries.

## Usage

To use the service discovery component, you need to create a `Registry` and a `Discovery` instance and provide them to the application.

```go
import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	consulAPI "github.com/hashicorp/consul/api"
)

func main() {
	// ...
	consulClient, err := consulAPI.NewClient(consulAPI.DefaultConfig())
	if err != nil {
		panic(err)
	}
	reg := consul.New(consulClient)

	app := kratos.New(
		// ...
		kratos.Registrar(reg),
	)
	// ...
}
```

## Examples

- [**registry.go**](./examples/registry.go): The core interfaces for service registration and discovery.

- `consul`
- `discovery`
- `etcd`
- `kubernetes`
- `nacos`
- `zookeeper`

## Related Skills

- **[kratos-transport](./../kratos-transport/SKILL.md)**: Service discovery is used by clients to find server endpoints.
- **[kratos-load-balancing](./../kratos-load-balancing/SKILL.md)**: The list of discovered services is used by the load balancer.
