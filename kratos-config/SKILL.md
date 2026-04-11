---
name: kratos-config
description: Use when loading application configuration from files or remote sources, and for handling dynamic configuration updates.
---

# Kratos Configuration

## Overview

This skill provides a guide to Kratos's powerful and flexible configuration management system. Kratos uses a unified interface to load configuration from various local or remote sources, merge them, and allow for dynamic updates by watching for changes.

## When to Use

Use this skill when you need to:
- Load application settings from a local file (e.g., `config.yaml`).
- Use a remote configuration center like Apollo, Nacos, Etcd, or Consul.
- Access configuration values within your application code.
- Implement hot-reloading of configuration when a value changes in the source.
- Define bootstrap configuration for your service.

## Core Pattern: Loading and Scanning

The most common pattern is to create a `Config` instance, provide one or more `Source`s, load the configuration, and then `Scan` it into a struct.

```go
import (
	"log"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

// Bootstrap struct must match the structure of your config file.
// It is recommended to use a tool like `koanf` to generate this.
type Bootstrap struct {
	Server struct {
		HTTP struct {
			Addr    string `json:"addr"`
			Timeout int    `json:"timeout"`
		} `json:"http"`
	} `json:"server"`
}

func main() {
	// 1. Create a new config instance with a file source.
	c := config.New(
		config.WithSource(
			file.NewSource("configs/config.yaml"),
		),
	)
	defer c.Close()

	// 2. Load the configuration sources.
	if err := c.Load(); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// 3. Scan the loaded configuration into your bootstrap struct.
	var bc Bootstrap
	if err := c.Scan(&bc); err != nil {
		log.Fatalf("failed to scan config: %v", err)
	}

	log.Printf("HTTP server address: %s", bc.Server.HTTP.Addr)
}
```

## Dynamic Updates: Watching for Changes

Kratos can watch for changes in the configuration source and trigger a callback.

```go
// Get a specific value.
addrValue, err := c.Value("server.http.addr")
if err != nil {
    log.Fatalf("could not get value: %v", err)
}

// Watch for changes to the value.
if err := addrValue.Watch(func(key string, value config.Value) {
    newAddr, _ := value.String()
    log.Printf("config changed! %s: %s", key, newAddr)
}); err != nil {
    log.Fatalf("could not watch value: %v", err)
}
```

## Available Sources

Kratos supports multiple sources, which are located in the `contrib/config/` directory of the core `kratos` repository.

- **`file`** (built-in)
- **`apollo`**
- **`consul`**
- **`etcd`**
- **`kubernetes`**
- **`nacos`**
- **`polaris`**

## Common Mistakes

- **Mismatched Struct Tags:** Ensure the field tags in your bootstrap struct (e.g., `json:"addr"`) exactly match the keys in your configuration file. Mismatches will cause fields to be zero-valued after scanning.
- **Forgetting `c.Close()`:** The `Config` instance may hold open connections to remote sources or file watchers. Always `defer c.Close()` to ensure a clean shutdown.
- **Ignoring `Load()` Errors:** `c.Load()` can fail if a file doesn't exist or a remote source is unreachable. Always check the error.
- **Blocking in Watch Callbacks:** The `Watch` callback is executed synchronously. Avoid long-running or blocking operations inside the callback to prevent deadlocks.

## Related Skills

- **kratos-encoding**: The config component uses codecs (like YAML, JSON) from the encoding package to parse configuration files.
