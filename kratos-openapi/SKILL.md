---
name: kratos-openapi
description: Use when generating OpenAPI/Swagger documentation from Protobuf definitions in a Kratos application.
---

# Kratos OpenAPI

This skill provides a guide to generating OpenAPI/Swagger documentation from Protobuf definitions in a Kratos application.

## Overview

Kratos can automatically generate OpenAPI v2 (Swagger) documentation from your Protobuf API definitions.

## Usage

1.  Install the `protoc-gen-openapiv2` tool:
    ```bash
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
    ```

2.  Add the `(google.api.http)` option to your Protobuf service methods.

3.  Generate the OpenAPI documentation:
    ```bash
    protoc --proto_path=. \
             --proto_path=./third_party \
             --go_out=paths=source_relative:. \
             --go-http_out=paths=source_relative:. \
             --openapiv2_out . \
             api/helloworld/v1/greeter.proto
    ```

## Examples

- [**openapi.yaml**](./examples/openapi.yaml): A snippet of a generated OpenAPI v2 (Swagger) specification.

- **[kratos-api-definition](./../kratos-api-definition/SKILL.md)**: OpenAPI documentation is generated from your Protobuf API definitions.
