---
name: kratos-ent-integration
description: Use when integrating the Ent ORM for database access within a Kratos application, including schema definition and client setup.
---

# Integrating the Ent ORM with Kratos

## Overview

This skill provides a guide to integrating Ent, a powerful entity framework for Go, into a Kratos application. The standard Kratos layout uses the `internal/data` package as the home for data access logic, which is the perfect place to manage your Ent schema and client.

## When to Use

Use this skill when you need to:
- Set up a database connection in your Kratos service.
- Define your data model (schema) using Ent's code-first approach.
- Generate database query code from your schema.
- Inject the Ent client into your data layer repositories.

## Core Workflow

1.  **Define the Schema:** In the `internal/data/` directory, create a new file for your schema definition (e.g., `internal/data/user.go`).

    ```go
    // internal/data/user.go
    package data

    import (
        "entgo.io/ent"
        "entgo.io/ent/schema/field"
    )

    // User holds the schema definition for the User entity.
    type User struct {
        ent.Schema
    }

    // Fields of the User.
    func (User) Fields() []ent.Field {
        return []ent.Field{
            field.Int64("id"),
            field.String("username"),
            field.String("password"),
        }
    }
    ```

2.  **Generate Ent Code:** From your project root, run `go generate ./...`. This will read your schema and generate all the necessary ORM code in `internal/data/ent/`.

    ```bash
    go generate ./...
    ```

3.  **Initialize the Ent Client:** In `internal/data/data.go`, create a `NewData` provider that initializes the Ent client and injects it into the `Data` struct. This uses the database configuration from your `conf.Data` struct.

    ```go
    // internal/data/data.go
    func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
        // ...
        client, err := ent.Open(c.Database.Driver, c.Database.Source)
        if err != nil {
            // ... handle error
        }

        // Run the auto migration tool.
        if err := client.Schema.Create(context.Background()); err != nil {
            // ... handle error
        }
        d := &Data{db: client}
        // ...
        return d, cleanup, nil
    }
    ```

4.  **Use the Client in Repositories:** Your repositories, also defined in the `data` package, will receive the `*Data` struct and can use the `db` field to access the Ent client and query the database.

## Common Mistakes

- **Forgetting to Run `go generate`:** After any change to your Ent schema files, you **must** run `go generate ./...` to update the generated ORM code. Forgetting this will lead to compile errors or runtime panics.
- **Incorrect Database Configuration:** Double-check that the `Driver` and `Source` in your `config.yaml` are correct for your database (e.g., `mysql`, `postgres`, `sqlite3`).
- **Running Migrations in Production:** The `client.Schema.Create(context.Background())` line is for automatic migrations during development. In a production environment, you should use Ent's versioned migrations to manage schema changes safely.

## Related Skills

- **kratos-layout**: The Ent schema and generated code reside in the `internal/data` directory as defined by the Kratos project layout.
- **kratos-dependency-injection**: The Ent client (`*ent.Client`) is provided to the rest of the application via dependency injection managed by Wire.
