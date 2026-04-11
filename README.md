# Kratos Development Skills

A structured repository of Kratos microservices framework best practices and patterns optimized for AI agents and LLMs.

## Structure

- `kratos-*/SKILL.md` - Individual skill files (one per topic)
- `SKILL.md` - Main router skill with overview and navigation
- `kratos-*/examples/` - Code examples for each skill

## Skills

### Project Setup (CRITICAL)

- `kratos-layout/` - Project structure and organization
- `kratos-cli/` - Using the Kratos command-line interface
- `kratos-api-definition/` - Defining service APIs using Protobuf

### Transport Layer (HIGH)

- `kratos-transport/` - Configuring HTTP and gRPC servers/clients
- `kratos-middleware/` - Adding cross-cutting concerns (logging, auth, tracing)
- `kratos-encoding/` - Handling different data serialization formats

### Configuration & DI (HIGH)

- `kratos-config/` - Loading and managing application configuration
- `kratos-dependency-injection/` - Managing dependencies with Google's Wire

### Error Handling & Observability (HIGH)

- `kratos-error-handling/` - Defining and handling structured errors
- `kratos-logging/` - Structured logging and integration with logging libraries
- `kratos-metrics/` - Instrumenting your application with metrics

### Service Communication (MEDIUM)

- `kratos-service-discovery/` - Service registration and discovery patterns
- `kratos-load-balancing/` - Client-side load balancing strategies
- `kratos-metadata/` - Passing contextual information between services

### Documentation & Integration (MEDIUM)

- `kratos-openapi/` - Generating OpenAPI/Swagger documentation
- `kratos-ent-integration/` - Using the Ent ORM for data access

### Deployment (LOW)

- `kratos-docker/` - Building Docker images for Kratos services

## Using These Skills

### Installation

To make these skills available in another project, you can use the provided `link-skills.sh` script to create symbolic links. This keeps your skills centralized and easy to update.

From within the `kratos-development` directory, run the script and provide the **_ABSOLUTE path_** to your target project:

```bash
# Make the script executable (only needs to be done once)
chmod +x ./link-skills.sh

# Link the skills to your chatbot project using an ABSOLUTE path
./link-skills.sh /path/to/your/project
```

**IMPORTANT:** You **MUST** use an absolute path to your project directory. The script will validate this and return an error if a relative path is used. This approach ensures the links are stable, but it means they are specific to your machine's file structure.

_(Note: The script is smart enough to handle paths that already include `.claude/skills`.)_

### Usage

Each skill follows this structure:

```markdown
---
name: kratos-skill-name
description: When to use this skill
---

# Skill Title

Brief description and overview.

## Overview
[Conceptual explanation]

## Basic Usage
[Code examples with HTTP and gRPC]

## Examples
[Links to example files]

## Related Skills
[Links to related skills]
```

## Impact Levels

- `CRITICAL` - Required for any Kratos project
- `HIGH` - Core functionality used in most microservices
- `MEDIUM` - Important for production applications
- `LOW` - Advanced features for specific use cases

## Quick Start

```bash
# Install Kratos CLI
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest

# Create new project
kratos new myservice

# Generate API code from proto
kratos proto client api/myservice/v1/myservice.proto

# Run the service
cd myservice && go run ./cmd/myservice
```

## Project Structure

```
myservice/
├── api/                    # Protobuf definitions
│   └── myservice/v1/
│       └── myservice.proto
├── cmd/                    # Application entry points
│   └── myservice/
│       └── main.go
├── configs/                # Configuration files
├── internal/
│   ├── biz/               # Business logic layer
│   ├── data/              # Data access layer
│   ├── server/            # Server setup (HTTP/gRPC)
│   └── service/           # Service implementations
└── go.mod
```

## Core Patterns

```go
// Define service in proto
service MyService {
  rpc CreateUser (CreateUserRequest) returns (CreateUserReply);
}

// Implement service
type MyServiceService struct {
  pb.UnimplementedMyServiceServer
  uc *biz.UserUsecase
}

func (s *MyServiceService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
  user, err := s.uc.CreateUser(ctx, &biz.User{Name: req.Name})
  if err != nil {
    return nil, err
  }
  return &pb.CreateUserReply{Id: user.ID}, nil
}

// Wire up servers
func main() {
  app := kratos.New(
    kratos.Name("myservice"),
    kratos.Server(httpSrv, grpcSrv),
  )
  app.Run()
}
```

## References

- [Kratos Official Documentation](https://go-kratos.dev/docs/)
- [Kratos GitHub Repository](https://github.com/go-kratos/kratos)
- [Kratos Examples](https://github.com/go-kratos/examples)
