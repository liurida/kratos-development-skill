
---
id: wire
title: Dependency Injection
---

**Wire** is a compile-time dependency injection tool.

It is recommended that doing explicit initialization rather than using global variables.

Generating the initialization codes by *Wire* can reduce the coupling among components and increase the maintainability of the project.

### Installation

```bash
# Import into project
go get -u github.com/google/wire

# Install cmd
go install github.com/google/wire/cmd/wire
```

### Terms

There are two basic terms in wire, *Provider* and *Injector*.

Provider is a *Go Func*, it can also receive the values from other *Provider*s for dependency injection.

```go
// provides a config file
func NewConfig() *conf.Data {...}

// provides the data component (the initialization of database, cache and etc.) which depends on the data config.
func NewData(c *conf.Data) (*Data, error) {...}

// provides persistence components (implementation of CRUD persistence) which depends on the data component.
func NewUserRepo(d *data.Data) (*UserRepo, error) {...}
```
