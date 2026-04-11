---
name: kratos-docker
description: Use when building Docker images for Kratos services, including understanding multi-stage builds for smaller production images.
---

# Kratos and Docker

## Overview

This skill provides a guide for containerizing Kratos applications using Docker. The default `kratos-layout` project includes a pre-configured, multi-stage `Dockerfile` designed to create a minimal, secure, and efficient production image.

## When to Use

Use this skill when you need to:
- Build a Docker image for your Kratos service.
- Understand the purpose of the multi-stage `Dockerfile` provided by the Kratos layout.
- Run your Kratos service as a Docker container.
- Customize the `Dockerfile` for your specific needs.

## Core Pattern: Multi-Stage Dockerfile

The `kratos-layout` `Dockerfile` uses a multi-stage build to keep the final production image as small as possible. This is a crucial best practice for security and performance.

```dockerfile
# 1. Build Stage: Uses a full Go build environment
FROM golang:1.19 AS builder

WORKDIR /app

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application binary
# CGO_ENABLED=0 disables CGO, creating a static binary
# -o /bin/server specifies the output path for the compiled binary
RUN CGO_ENABLED=0 go build -o /bin/server ./cmd/server

# 2. Final Stage: Uses a minimal, non-root base image
FROM scratch

# Copy only the compiled binary and configuration from the builder stage
COPY --from=builder /bin/server /bin/server
COPY --from=builder /configs /configs

# Expose the default HTTP and gRPC ports
EXPOSE 8000
EXPOSE 9000

# Set the container's entrypoint
ENTRYPOINT ["/bin/server", "-conf", "/configs"]
```

**Why Multi-Stage?**
- **Security:** The final image doesn't contain the Go toolchain, source code, or any build-time dependencies, reducing the attack surface.
- **Size:** The final image is tiny (often < 20MB) because it only contains the compiled binary and configuration, whereas the `builder` stage can be hundreds of MBs.

## Usage Workflow

1.  **Build the Docker Image:** From the root of your project, run the `docker build` command. The `-t` flag tags the image with a name and version.

    ```bash
    docker build -t my-app:latest .
    ```

2.  **Run the Docker Container:** Use `docker run` to start your service. The `-p` flags map your local ports to the container's exposed ports.

    ```bash
    # Maps local port 8000 to container port 8000 (HTTP)
    # Maps local port 9000 to container port 9000 (gRPC)
    docker run -p 8000:8000 -p 9000:9000 my-app:latest
    ```

## Common Mistakes

- **Not Using `.dockerignore`:** Without a `.dockerignore` file, your build context can become very large, including local `vendor` directories, `.git` history, and other unnecessary files, slowing down your build.
- **Running as Root:** The final stage uses `FROM scratch`, which is a good practice. Avoid changing this to a full OS image like `ubuntu` and running as the root user in production.
- **Copying Unnecessary Files:** Only copy the compiled binary and the `configs` directory into the final image. Do not copy the entire source code.

## Related Skills

- **kratos-layout**: The `Dockerfile` is designed to work with the standard Kratos project layout, especially the `cmd/server` and `configs` directories.
