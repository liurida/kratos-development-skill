---
name: kratos-cli
description: Use when scaffolding new projects, managing Protobuf files, or running a Kratos application from the command line.
---

# Kratos CLI

## Overview

This skill is a reference guide for the Kratos command-line interface (CLI), the primary tool for creating, managing, and running Kratos projects. It covers project scaffolding, Protobuf code generation, and application execution.

## When to Use

Use this skill when you need to:
- Create a new Kratos project from scratch (`kratos new`).
- Add a new Protobuf service definition to your project (`kratos proto add`).
- Generate gRPC and HTTP client or server stubs from your `.proto` files (`kratos proto client/server`).
- Run your Kratos application locally for development (`kratos run`).
- Upgrade the `kratos` tool to the latest version (`kratos upgrade`).

## Quick Reference: Main Commands

| Command | Description | Common Flags |
|---|---|---|
| `kratos new <project>` | Creates a new Kratos project. | `-r <repo>` (custom layout), `--nomod` (monorepo) |
| `kratos proto add <path>` | Adds a new Protobuf file. | |
| `kratos proto client <path>` | Generates client-side code from Protobuf. | |
| `kratos proto server <path>`| Generates server-side code from Protobuf. | `-t <dir>` (target directory) |
| `kratos run` | Compiles and runs the application. | |
| `kratos upgrade` | Upgrades the Kratos tool and layout. | |
| `kratos changelog` | Generates a changelog from git history. | `dev` (for unreleased changes) |

## Core Workflow: Creating a New Service

This workflow demonstrates the most common CLI use case: creating a new service, defining its API, and running it.

```bash
# 1. Create a new project using the default layout
# This scaffolds a complete project structure.
kratos new helloworld

# 2. Navigate into the new project directory
cd helloworld

# 3. Add a new Protobuf definition for a "greeter" service
# This creates api/helloworld/v1/greeter.proto and updates api/helloworld/helloworld.proto
kratos proto add api/helloworld/v1/greeter.proto

# 4. Generate the server implementation for the new proto
# This creates internal/service/greeter.go
kratos proto server api/helloworld/v1/greeter.proto -t internal/service

# 5. Run the application
# This command compiles and runs your main.go
kratos run
```

## Common Mistakes

- **Running `kratos proto` from the Wrong Directory:** Always run `kratos proto` commands from the root of your Kratos project, not from within the `api/` or `internal/` directories.
- **Forgetting the `-t` Flag:** When using `kratos proto server`, you must specify the target directory with the `-t` flag (e.g., `-t internal/service`) or the command will fail.
- **Manually Creating Proto Files:** Avoid creating `.proto` files by hand. Use `kratos proto add` to ensure they are correctly registered and included in the build process.
- **`kratos upgrade` Failing:** If `kratos upgrade` fails, it might be due to a dirty git working directory. Commit or stash your changes before running the upgrade.

## Related Skills

- **kratos-layout**: The CLI creates projects based on the structure defined in this skill.
- **kratos-api-definition**: The CLI is used to generate code from the API definitions you create.
