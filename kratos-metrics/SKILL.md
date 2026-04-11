---
name: kratos-metrics
description: Use for instrumenting a Kratos service with metrics, particularly for exposing a Prometheus endpoint to track request latency and counts.
---

# Kratos Metrics

## Overview

This skill provides a guide to instrumenting a Kratos application with metrics. Kratos uses a middleware-based approach to automatically track key indicators like request latency and total request counts. The primary supported backend is Prometheus, a popular open-source monitoring system.

## When to Use

Use this skill when you need to:
- Monitor the health and performance of your Kratos service.
- Expose a `/metrics` endpoint for a Prometheus server to scrape.
- Track the duration (latency) of your API requests.
- Count the total number of requests handled by your service.
- Create and update custom application-specific metrics (e.g., `active_users`).

## Core Pattern: Instrumenting with Prometheus

This workflow shows how to add the metrics middleware to your HTTP server to expose a Prometheus endpoint.

```go
import (
	"net/http"

	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/prometheus/client_golang/prometheus"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
)

// 1. Define your Prometheus metrics vectors.
// These are typically global variables in your server setup.
var (
	_metricSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "server",
		Subsystem: "requests",
		Name:      "duration_sec",
		Help:      "server requests duration(sec).",
	}, []string{"kind", "operation"})

	_metricRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "client",
		Subsystem: "requests",
		Name:      "code_total",
		Help:      "The total number of processed requests",
	}, []string{"kind", "operation", "code", "reason"})
)

func NewHTTPServer(c *conf.Server, greeter *service.GreeterService) *http.Server {
	// ... server setup

	// 2. Add the metrics middleware to your HTTP server.
	srv.HandlePrefix("/", http.Middleware(
		metrics.Server(
			metrics.WithSeconds(prom.NewHistogram(_metricSeconds)),
			metrics.WithRequests(prom.NewCounter(_metricRequests)),
		),
	)(greeter.ServeHTTP))

	// 3. Expose the /metrics endpoint for Prometheus to scrape.
	srv.Handle("/metrics", promhttp.Handler())

	return srv
}
```

## How it Works

1.  The `metrics.Server` middleware intercepts every incoming request.
2.  It starts a timer before passing the request to your service.
3.  When the request is finished, it stops the timer and records the duration in the `_metricSeconds` Prometheus histogram.
4.  It also increments the `_metricRequests` counter, including labels for the request `kind` (e.g., `server`), `operation` (e.g., `/helloworld.v1.Greeter/SayHello`), and result `code`.
5.  The `promhttp.Handler()` exposes all registered Prometheus metrics (including yours) via an HTTP handler at the `/metrics` path.

## Common Mistakes

- **Forgetting to Expose the `/metrics` Endpoint:** Adding the middleware is not enough. You must also explicitly expose the `promhttp.Handler()` on a path (usually `/metrics`) so the Prometheus scraper can access the data.
- **Not Registering the Metrics:** While the example shows global variables, in a real application you might need to call `prometheus.MustRegister(_metricSeconds, _metricRequests)` if they are not registered automatically. The Kratos middleware handles this for you, but it's a common issue when creating custom metrics.
- **High Cardinality Labels:** Be careful about the labels you add to your metrics. Using labels with unbounded values (like a `user_id`) can cause a "cardinality explosion" and overwhelm your Prometheus server. Use labels for low-cardinality values like `operation` or `status_code`.

## Related Skills

- **kratos-middleware**: The metrics instrumentation is implemented as a standard Kratos middleware.
- **kratos-transport**: The metrics endpoint is exposed via the HTTP transport server.
