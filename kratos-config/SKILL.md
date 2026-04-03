---
name: kratos-config
description: Use when managing configuration in a Kratos application.
---

# Kratos Configuration

This skill provides a guide to managing configuration in a Kratos application.

## Overview

Kratos provides a unified interface for loading configuration from various sources, such as local files or remote configuration centers. It also supports dynamic configuration updates.

## Usage

To use the configuration component, you need to create a new `Config` instance and provide one or more `Source`s.

```go
import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

func main() {
	c := config.New(
		config.WithSource(
			file.NewSource("configs/config.yaml"),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc struct{
		Server struct{
			Http struct{
				Addr string
				Timeout int
			}
		}
	}

	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
}
```

## Examples

- [**config.yaml**](./examples/config.yaml): A typical configuration file for a Kratos service, defining server and data source settings.

- `file` (built-in)
- `apollo`
- `etcd`
- `kubernetes`
- `nacos`

## Related Skills

- **[kratos-encoding](./../kratos-encoding/SKILL.md)**: The config component uses codecs from the encoding package.
