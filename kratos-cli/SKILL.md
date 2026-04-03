---
name: kratos-cli
description: Use when working with the Kratos command-line interface.
---

# Kratos CLI

This skill provides a guide to using the Kratos command-line interface.

## Main Commands

- **`kratos new <project-name>`**: Creates a new Kratos project.
    - `-r`: Specify a custom layout repository.
    - `-b`: Specify a branch.
    - `--nomod`: Create a service in a monorepo.
- **`kratos proto add <path>`**: Adds a new Protobuf file.
- **`kratos proto client <path>`**: Generates client-side code from a Protobuf file.
- **`kratos proto server <path> -t <target-dir>`**: Generates server-side code from a Protobuf file.
- **`kratos run`**: Runs the Kratos application.
- **`kratos upgrade`**: Upgrades the Kratos tool and dependencies.
- **`kratos changelog`**: Shows the changelog.
- **`kratos -v`**: Shows the version.

## Examples

- [**common-usage.sh**](./examples/common-usage.sh): A shell script demonstrating a typical workflow using the Kratos CLI.

For more detailed information, see the [Kratos CLI Usage](./reference/usage.md) document.
