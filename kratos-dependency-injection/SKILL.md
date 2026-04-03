
---
name: kratos-dependency-injection
description: Use when working with dependency injection in a Kratos application.
---

# Kratos Dependency Injection

This skill provides a guide to working with dependency injection in a Kratos application using Wire.

## Core Concepts

- **Provider**: A Go function that provides a value.
- **Injector**: A function that calls providers to build a complete object graph.
- **ProviderSet**: A collection of providers for a module.

## Usage

1.  Define a `ProviderSet` in each module (`data`, `biz`, `service`, `server`).
2.  In `cmd/server/wire.go`, use `wire.Build` to assemble the `ProviderSet`s and the `newApp` injector.
3.  Run `go generate ./...` to generate the dependency injection code in `wire_gen.go`.

## Example

`cmd/server/wire.go`:
```go
//go:build wireinject
// +build wireinject

package main

import (
    "github.com/go-kratos/kratos/v2"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/google/wire"
    "helloworld/internal/biz"
    "helloworld/internal/conf"
    "helloworld/internal/data"
    "helloworld/internal/server"
    "helloworld/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
    panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
```

## Related Skills

- **[kratos-layout](./../kratos-layout/SKILL.md)**: Dependency injection is a key part of the Kratos project structure.
