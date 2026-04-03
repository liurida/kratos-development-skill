---
name: kratos-docker
description: Use when working with Docker in a Kratos application.
---

# Kratos Docker

This skill provides a guide to working with Docker in a Kratos application.

## Overview

The Kratos layout project includes a `Dockerfile` that can be used to build a Docker image for your application.

## Usage

1.  Build the Docker image:
    ```bash
    docker build -t my-app:latest .
    ```

2.  Run the Docker container:
    ```bash
    docker run -p 8000:8000 -p 9000:9000 my-app:latest
    ```

## Examples

- [**Dockerfile**](./examples/Dockerfile): A multi-stage Dockerfile for building a Kratos service.

- **[kratos-layout](./../kratos-layout/SKILL.md)**: The Dockerfile is designed to work with the standard Kratos project layout.
