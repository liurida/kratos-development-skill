---
name: kratos-openapi
description: Use for generating and serving OpenAPI v2 (Swagger) documentation from your Protobuf API definitions in a Kratos service.
---

# Kratos OpenAPI (Swagger) Generation

## Overview

This skill provides a guide to automatically generating OpenAPI v2 (also known as Swagger) documentation from your Protobuf service definitions. Kratos leverages the `grpc-gateway` toolchain to convert your gRPC service definitions and HTTP annotations into a standard `openapi.yaml` file, which can then be served via a Swagger UI.

## When to Use

Use this skill when you need to:
- Provide interactive API documentation for your Kratos service.
- Generate a machine-readable `openapi.yaml` specification for other tools to consume.
- Ensure your HTTP API documentation is always in sync with your Protobuf definitions.

## Core Workflow

**1. Install the Generator:**
First, you need to install the `protoc-gen-openapiv2` tool which reads your Protobuf files and outputs an OpenAPI spec.

```bash
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
```

**2. Annotate your Protobuf Service:**
Ensure your RPC methods have `(google.api.http)` annotations, as these are used to generate the OpenAPI paths and methods.

```protobuf
// api/helloworld/v1/greeter.proto

rpc SayHello (HelloRequest) returns (HelloReply)  {
  option (google.api.http) = {
    get: "/helloworld/{name}"
  };
}
```

**3. Integrate Generation into your Build Process:**
Manually running `protoc` is tedious. The best practice is to add the generation step to your project's `Makefile` or a `generate.go` file. The default `kratos-layout` uses a `buf.gen.yaml` file for this.

```yaml
# buf.gen.yaml
version: v1
plugins:
  # ... other plugins like go, go-http, go-grpc
  - plugin: openapiv2
    out: .
```

Then, you can generate all artifacts with a single command:

```bash
# From your project root
buf generate
```
This will create an `openapi.yaml` file in the same directory as your `.proto` file.

**4. Serve the Swagger UI:**
Kratos does not serve the Swagger UI by default. You need to add an HTTP handler to serve the generated `openapi.yaml` and the Swagger UI assets.

```go
// cmd/server/main.go

import (
	"net/http"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

func main() {
	// ... server setup

	// Create the OpenAPI handler
	openapiHandler := openapiv2.NewHandler()

	// Create a new HTTP server
	httpSrv := http.NewServer(
		// ... other options
	)

	// Mount the OpenAPI handler
	httpSrv.HandlePrefix("/q/", openapiHandler)

	// ... app run
}
```
When you run your server, you can now access the interactive documentation at `http://localhost:8000/q/swagger-ui/`.

## Common Mistakes

- **Missing HTTP Annotations:** If your RPC methods don't have `(google.api.http)` options, `protoc-gen-openapiv2` will have nothing to generate, resulting in an empty or incomplete spec.
- **Incorrectly Serving the UI:** The `openapiv2.NewHandler()` serves both the `openapi.yaml` file and the Swagger UI assets. Ensure you mount it with `HandlePrefix` as shown above. Mounting it with `Handle` will not work correctly.
- **Out-of-Sync Documentation:** If you change your `.proto` file but forget to run `buf generate` (or `protoc`), your served documentation will be out of date. Always regenerate after making API changes.

## Related Skills

- **kratos-api-definition**: The OpenAPI documentation is generated directly from the annotations and definitions in your Protobuf files.
- **kratos-transport**: The generated Swagger UI is served via an HTTP handler on the Kratos HTTP transport.
