---
id: encoding
title: Encoding
keywords:
  - Go
  - Kratos
  - Toolkit
  - Framework
  - Microservices
  - Protobuf
  - gRPC
  - HTTP
---
We've abstracted the `Codec` interface to unify the serialization/deserialization logic for processing requests, and you can implement your own Codec to support more formats. The specific source code is in [encoding](https://github.com/go-kratos/kratos/tree/main/encoding)。

These formats are battery-included.
* form
* json
* protobuf
* xml
* yaml

### Interface

You should implement the following Codec interface for your custom codec.

```go
// Codec interface is for serialization and deserialization, notice that these methods must be thread-safe.
type Codec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
	Name() string
}
```

### Usage

#### Register Custom Codec

```go
encoding.RegisterCodec(codec{})
```

#### Get the Codec

```go
jsonCodec := encoding.GetCodec("json")
```
