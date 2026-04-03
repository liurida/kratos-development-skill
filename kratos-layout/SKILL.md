
---
name: kratos-layout
description: Use when working with the project structure of a Kratos application.
---

# Kratos Project Layout

This skill provides a guide to the project structure of a Kratos application.

## Directory Structure

A Kratos project is organized as follows:

- `api/`: Contains `.proto` API files and generated `.go` files.
- `cmd/`: The entry point of the application.
- `configs/`: Configuration files for local development.
- `internal/`: Private application and business logic.
    - `biz/`: Business logic layer (similar to domain layer in DDD).
    - `conf/`: Configuration struct definitions.
    - `data/`: Data access layer (repositories, database connections).
    - `server/`: HTTP and gRPC server instances.
    - `service/`: Service layer implementing the API.

## Reference

For more detailed information, see the [Kratos Project Layout](./reference/layout.md) document.
