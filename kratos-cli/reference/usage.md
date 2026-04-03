
---
id: usage
title: Usage
---

## Installation

```bash
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

## Project Creation

To create a new project:
```bash
kratos new helloworld
```

Use `-r` to specify the source

```bash
# If pull fails in China, you can use gitee source.
kratos new helloworld -r https://gitee.com/go-kratos/kratos-layout.git
# You can also use custom templates
kratos new helloworld -r xxx-layout.git
# You can also specify the source through the environment variable
KRATOS_LAYOUT_REPO=xxx-layout.git
kratos new helloworld
```

Use `-b` to specify the branch

```bash
kratos new helloworld -b main
```

Use `--nomod` to add services and working together using ` go.mod `, large warehouse mode

```bash
kratos new helloworld
cd helloworld
kratos new app/user --nomod
```

## Adding Proto files

```bash
kratos proto add api/helloworld/demo.proto
```

## Generate Proto Codes
```bash
kratos proto client api/helloworld/demo.proto
```

## Generate Service Codes
kratos can generate the bootstrap codes from the proto file.
```bash
kratos proto server api/helloworld/demo.proto -t internal/service
```

## Run project

- If there are multiple items under the subdirectory, the selection menu appears

```bash
kratos run
```

## View Version

To show the tool version

```bash
kratos -v
```

## Tool upgrade

The following tools will be upgraded

- Kratos and the tool itself
- Protoc related build plugins

```bash
kratos upgrade
```

## Changelog

```bash
# Equivalent to printing the version changelog of https://github.com/go-kratos/kratos/releases/latest 
kratos changelog

# Print the update log of the specified version
kratos changelog v2.1.4

# View the changelog since the latest release to now
kratos changelog dev
```
