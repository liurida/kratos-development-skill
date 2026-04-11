---
name: kratos-middleware
description: Use for adding cross-cutting concerns like logging, tracing, or authentication to Kratos services, and for creating custom middleware.
---

# Kratos Middleware

## Overview

This skill provides a guide to Kratos's middleware system. Middleware allows you to wrap your service's business logic with reusable components that handle cross-cutting concerns. It's a powerful mechanism for keeping your service logic clean and separating concerns like logging, metrics, authentication, and error recovery.

## When to Use

Use this skill when you need to:
- Add a standard behavior (like logging or tracing) to all incoming requests.
- Implement a custom authentication or authorization check.
- Create a middleware to transform requests or responses.
- Understand the order in which middleware is executed.

## Core Pattern: Creating Custom Middleware

Middleware is a function that takes a `Handler` and returns a `Handler`. The `Handler` is the next step in the chain, which could be another middleware or your actual service logic.

This example shows a simple middleware that measures the execution time of a request.

```go
import (
	"context"
	"log"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
)

func LatencyLogger() middleware.Middleware {
	// The outer function is called once during server setup.
	return func(handler middleware.Handler) middleware.Handler {
		// The inner function is called for every request.
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			// 1. Pre-processing: Code here runs BEFORE the handler.
			startTime := time.Now()
			log.Printf("Request started for: %T", req)

			// 2. Call the next handler in the chain.
			reply, err = handler(ctx, req)

			// 3. Post-processing: Code here runs AFTER the handler.
			latency := time.Since(startTime)
			log.Printf("Request finished in %s", latency)

			return reply, err
		}
	}
}
```

## Middleware Execution Order (FILO)

Middleware in Kratos is executed in a **First-In, Last-Out (FILO)** or "onion" model. The first middleware you register is the outermost layer, and the last one you register is the innermost.

Consider this setup:

```go
http.NewServer(http.Middleware(
    MiddlewareA, // Outermost
    MiddlewareB,
    MiddlewareC, // Innermost
))
```

The execution flow for a request is:
1.  Request enters `MiddlewareA` (pre-processing)
2.  Request enters `MiddlewareB` (pre-processing)
3.  Request enters `MiddlewareC` (pre-processing)
4.  Request is handled by your `service` logic
5.  Response leaves `MiddlewareC` (post-processing)
6.  Response leaves `MiddlewareB` (post-processing)
7.  Response leaves `MiddlewareA` (post-processing)

## Common Mistakes

- **Forgetting to Call `handler(ctx, req)`:** If you don't call the next handler in your middleware, the request chain will be broken, and your service logic will never be executed. The client will likely time out.
- **Modifying the Wrong Context:** If you need to add a value to the context for downstream middleware or your service to read, you must create a new context and pass it to the handler (e.g., `handler(newCtx, req)`). You cannot modify the incoming `ctx` directly.
- **Incorrect Order:** The order of middleware matters. For example, a `tracing` middleware should generally come before a `logging` middleware so that the logger can pick up the trace ID that tracing adds to the context.

## Related Skills

- **kratos-transport**: Middleware is applied to transports (HTTP and gRPC) when the server is created.
- **kratos-logging**: The logging middleware is a common and important middleware for observing your service.
- **kratos-metrics**: The metrics middleware is another crucial piece for monitoring service health.
