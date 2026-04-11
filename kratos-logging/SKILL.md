---
name: kratos-logging
description: Use for structured, contextual logging in a Kratos application, including integrating with common logging libraries like Zap or Logrus.
---

# Kratos Logging

## Overview

This skill provides a guide to Kratos's structured logging interface. The logging system is designed around a core `Logger` interface and a `Helper` wrapper to provide a simple, powerful, and extensible way to log key-value pairs. This approach makes your logs machine-parseable and easy to query in log management systems.

## When to Use

Use this skill when you need to:
- Add logging to your business logic (`biz`) or service layers (`service`).
- Log structured data using key-value pairs.
- Add contextual information (like a `trace_id`) to all log messages within a request.
- Integrate a third-party logging library like Zap or Logrus into your Kratos application.

## Core Pattern: Logger and Helper

Kratos logging has two main components:
- **`log.Logger`**: A simple, low-level interface that takes key-value pairs.
- **`log.Helper`**: A wrapper around a `Logger` that provides level-based logging methods (`Info`, `Warn`, `Error`, etc.) and printf-style formatting.

It is recommended to use the `Helper` for most application-level logging.

```go
import (
	"os"
	"github.com/go-kratos/kratos/v2/log"
)

func main() {
	// 1. Create a base logger. The default is a simple standard library logger.
	logger := log.NewStdLogger(os.Stdout)

	// 2. Create a Helper for level-based logging.
	help := log.NewHelper(logger)

	// 3. Log messages with structured key-value pairs.
	help.Infow("event", "user.login", "user_id", 12345)
	// Output: level=INFO event=user.login user_id=12345
}
```

## Contextual Logging

You can add fields to a logger's context. These fields will be included in all subsequent log messages created from that logger instance. This is extremely useful for adding request-scoped data like a trace ID.

```go
// Add a trace_id to the logger's context.
ctxLogger := log.With(logger, "trace_id", "abc-123")

helper := log.NewHelper(ctxLogger)

helper.Info("this is a message")
// Output: level=INFO trace_id=abc-123 msg="this is a message"

helper.Warn("this is another message")
// Output: level=WARN trace_id=abc-123 msg="this is another message"
```

## Available Logger Implementations

Kratos provides integrations for several popular logging libraries in the `contrib/log/` directory.

- **`std`** (built-in)
- **`fluent`**
- **`logrus`**
- **`zap`**
- **`zerolog`**

To use one, simply create it and pass it to `log.NewHelper`.

```go
// Example using the Zap logger integration
import "github.com/go-kratos/contrib/log/zap/v2"

rawLogger, _ := zap.NewProduction()
logger := zap.NewLogger(rawLogger)

helper := log.NewHelper(logger)
helper.Info("Hello from Zap!")
```

## Common Mistakes

- **Using Printf for Structured Data:** Avoid doing `log.Infof("user %d logged in", userID)`. The better, structured way is `log.Infow("event", "user.login", "user_id", userID)`. This makes your logs searchable by `user_id`.
- **Not Using Contextual Loggers:** Creating a new logger for each request with request-specific context (like `trace_id`) is a best practice. The logging middleware in Kratos does this for you automatically.
- **Creating Helpers Repeatedly:** Create the `log.Helper` once and reuse it. There is no need to call `log.NewHelper` every time you want to log a message.

## Related Skills

- **kratos-middleware**: The logging middleware automatically creates a request-scoped logger with a `trace_id` and other useful context, which is then passed into your service methods.
- **kratos-dependency-injection**: The logger is typically initialized in `cmd/server/main.go` and provided to the rest of the application via dependency injection.
