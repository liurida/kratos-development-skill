---
name: kratos-logging
description: Use when logging in a Kratos application.
---

# Kratos Logging

This skill provides a guide to logging in a Kratos application.

## Overview

Kratos' logging module consists of two parts:

1.  **`Logger`**: A low-level logging interface.
2.  **`Helper`**: A high-level logging interface with helper functions for log levels and formatting.

## Usage

It is recommended to use the `Helper` for business logic.

```go
import "github.com/go-kratos/kratos/v2/log"

func main() {
    logger := log.NewStdLogger(os.Stdout)
    log := log.NewHelper(logger)

    log.Info("Hello Kratos")
    log.Infof("Hello %s", "Kratos")
}
```

## Examples

- [**log.go**](./examples/log.go): The core implementation of the Kratos logging interface.

- `std` (built-in)
- `fluent`
- `zap`

## Related Skills

- **[kratos-middleware](./../kratos-middleware/SKILL.md)**: A logging middleware is provided to log request details.
