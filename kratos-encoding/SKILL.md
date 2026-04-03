---
name: kratos-encoding
description: Use when working with different data serialization formats in a Kratos application.
---

# Kratos Encoding

This skill provides a guide to working with different data serialization formats in a Kratos application.

## Overview

Kratos provides an encoding package that supports various data serialization formats. The framework will automatically select the appropriate codec based on the `Content-Type` header.

## Supported Formats

- `json`
- `proto`
- `yaml`
- `xml`
- `form`

## Examples

- [**encoding.go**](./examples/encoding.go): The core interface for the Kratos encoding and decoding.

To use a specific codec, you need to register it with the application.

```go
import "github.com/go-kratos/kratos/v2/encoding/json"

func main() {
	// The json codec is registered by default
}
```

## Related Skills

- **[kratos-config](./../kratos-config/SKILL.md)**: The config component uses codecs from the encoding package to parse configuration files.
