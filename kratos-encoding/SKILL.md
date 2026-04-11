---
name: kratos-encoding
description: Use when you need to handle different data serialization formats like JSON, Protobuf, and XML, or when you need to implement a custom codec.
---

# Kratos Encoding

## Overview

This skill provides a guide to Kratos's encoding package, which allows the framework to seamlessly handle multiple data serialization formats. Kratos automatically selects the appropriate codec for marshalling and unmarshalling based on the `Content-Type` and `Accept` HTTP headers.

## When to Use

Use this skill when you need to:
- Understand which serialization formats Kratos supports out of the box.
- Implement a client that needs to request a specific response format (e.g., `application/xml`).
- Create and register a custom codec for a format not supported by default (e.g., `msgpack`).

## Quick Reference: Built-in Codecs

By default, the following codecs are registered via blank imports in the `transport` package. Kratos will automatically use them based on the `Content-Type` header.

| Name | `Content-Type` | Package |
|---|---|---|
| `json` | `application/json` | `github.com/go-kratos/kratos/v2/encoding/json` |
| `proto`| `application/proto`| `github.com/go-kratos/kratos/v2/encoding/proto`|
| `yaml` | `application/yaml` | `github.com/go-kratos/kratos/v2/encoding/yaml` |
| `xml` | `application/xml` | `github.com/go-kratos/kratos/v2/encoding/xml` |
| `form` | `application/x-www-form-urlencoded` | `github.com/go-kratos/kratos/v2/encoding/form` |

## Core Pattern: Implementing a Custom Codec

To support a new serialization format, you must implement the `encoding.Codec` interface and register it.

```go
// 1. Implement the Codec interface

import (
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/vmihailenco/msgpack/v5"
)

const Name = "msgpack"

// Define your custom codec struct
type Codec struct{}

func (c Codec) Marshal(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

func (c Codec) Unmarshal(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}

func (c Codec) Name() string {
	return Name
}

// 2. Register the codec in your main.go

func init() {
	encoding.RegisterCodec(Codec{})
}

func main() {
	// Now Kratos can handle "Content-Type: application/msgpack"
	// ... your app startup code
}
```

## Common Mistakes

- **Forgetting to Register a Custom Codec:** If you implement a custom codec but forget to call `encoding.RegisterCodec(MyCodec{})` in an `init()` function, Kratos will not be aware of it and will return a `415 Unsupported Media Type` error.
- **Incorrect `Content-Type` Header:** When making requests, ensure the `Content-Type` header exactly matches the name of a registered codec (e.g., `application/json`, `application/xml`).
- **Blank Importing a Codec:** The built-in codecs are enabled via blank imports (e.g., `_ "github.com/go-kratos/kratos/v2/encoding/json"`). If you build a very minimal binary, you might need to explicitly add these blank imports if they are not already pulled in by another package like `transport`.

## Related Skills

- **kratos-config**: The config component uses these codecs to parse configuration files in different formats.
- **kratos-transport**: The HTTP transport uses codecs to automatically handle request and response bodies based on `Content-Type` headers.
