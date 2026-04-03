---
name: kratos-metadata
description: Use when working with metadata in a Kratos application.
---

# Kratos Metadata

This skill provides a guide to working with metadata in a Kratos application.

## Overview

Kratos provides a metadata package that can be used to pass and retrieve metadata between services.

## Usage

### Passing Metadata

```go
import "github.com/go-kratos/kratos/v2/metadata"

md := metadata.New(map[string]string{"key": "value"})
ctx := metadata.NewClientContext(context.Background(), md)
```

### Retrieving Metadata

```go
import "github.com/go-kratos/kratos/v2/metadata"

if md, ok := metadata.FromServerContext(ctx); ok {
    val := md.Get("key")
}
```

## Examples

- [**metadata.go**](./examples/metadata.go): The core implementation of the Kratos metadata package.

- **[kratos-transport](./../kratos-transport/SKILL.md)**: Metadata is passed between services via the transport layer.
