---
name: kratos-metrics
description: Use when working with metrics in a Kratos application.
---

# Kratos Metrics

This skill provides a guide to working with metrics in a Kratos application.

## Overview

Kratos provides a metrics interface that can be used to report service statistics to a monitoring platform.

## Usage

To use the metrics component, you need to create a new `Metrics` instance and register it with the application.

```go
import (
	"github.com/go-kratos/kratos/v2/metrics"
	"github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
)

func main() {
	// ...
	app := kratos.New(
		// ...
		kratos.Server(
			// ...
			http.Middleware(
				metrics.Server(
					metrics.WithSeconds(prom.NewHistogram(_metricSeconds)),
					metrics.WithRequests(prom.NewCounter(_metricRequests)),
				),
			),
		),
	)
	// ...
}
```

## Examples

- [**metrics.go**](./examples/metrics.go): The core implementation of the Kratos metrics middleware.

- `datadog`
- `prometheus`

## Related Skills

- **[kratos-middleware](./../kratos-middleware/SKILL.md)**: A metrics middleware is provided to instrument requests.
