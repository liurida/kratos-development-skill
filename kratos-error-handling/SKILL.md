---
name: kratos-error-handling
description: Use when defining, creating, and asserting structured errors in a Kratos application, especially when using Protobuf to define error codes.
---

# Kratos Error Handling

## Overview

This skill provides a guide to Kratos's structured error handling system. Kratos enables you to define your business errors directly in your `.proto` files, which then generate helper functions in Go. This ensures consistency between your API contract and your business logic, and allows Kratos to automatically map your business errors to the correct gRPC and HTTP status codes.

## When to Use

Use this skill when you need to:
- Define a new business error (e.g., "user not found", "invalid parameter").
- Return a structured error from your service's business logic (`biz`) or service layer (`service`).
- Check for a specific type of Kratos error in your code.
- Understand how Kratos maps your custom error `Reason` to HTTP status codes.

## Core Workflow: Proto to Go

This workflow shows the end-to-end process of defining an error in Protobuf, generating code, creating it, and asserting it.

**1. Define Errors in Protobuf:**
In your `api/.../v1/error_reason.proto` file, define an `enum` for your errors. The `errors.code` option maps the enum value to an HTTP status code.

```protobuf
// api/helloworld/v1/error_reason.proto

enum ErrorReason {
  // The default_code will be used if an explicit code is not set.
  option (errors.default_code) = 500;

  // Maps USER_NOT_FOUND to HTTP 404 Not Found.
  USER_NOT_FOUND = 0 [(errors.code) = 404];

  // Maps CONTENT_MISSING to HTTP 400 Bad Request.
  CONTENT_MISSING = 1 [(errors.code) = 400];
}
```

**2. Generate Go Code:**
Run `go generate ./...` or `kratos proto server` to trigger the `protoc-gen-go-errors` plugin. This generates two key files:
- `error_reason.pb.go`: The base Protobuf generated code.
- `error_reason_errors.pb.go`: Contains the Kratos error helper functions.

For `USER_NOT_FOUND`, it will generate:
- `IsUserNotFound(err error) bool`: A helper to check if an error is of this type.
- `ErrorUserNotFound(format string, a ...interface{}) *errors.Error`: A constructor to create a new error of this type.

**3. Create and Return Errors:**
In your service or business logic, use the generated constructor to create and return the error.

```go
// internal/biz/greeter.go

import "github.com/go-kratos/kratos/v2/errors"

func (uc *GreeterUsecase) GetUser(ctx context.Context, id int64) (*User, error) {
    if id <= 0 {
        // You can also create errors manually.
        return nil, errors.New(400, "INVALID_ARGUMENT", "user id must be greater than 0")
    }
    user, err := uc.repo.FindUserByID(ctx, id)
    if err != nil {
        // Use the generated error constructor.
        return nil, api.ErrorUserNotFound("user %d not found", id)
    }
    return user, nil
}
```

**4. Assert Errors:**
To handle a specific error, use the generated `Is...` helper function.

```go
// internal/service/greeter.go

import "path/to/api/helloworld/v1"

user, err := uc.GetUser(ctx, 123)
if err != nil {
    if api.IsUserNotFound(err) {
        // Handle the "not found" case specifically
        return nil, errors.NotFound("USER_NOT_FOUND", "The requested user does not exist")
    }
    // Handle other errors
    return nil, err
}
```

## Common Mistakes

- **Forgetting the `errors.code` Option:** If you don't specify `(errors.code)`, Kratos will use the `default_code` (e.g., 500), which is often not what you want. Always explicitly map your error reasons to HTTP status codes.
- **Asserting with `==`:** Do not use `err == api.ErrorUserNotFound(...)`. The error may be wrapped. Always use the `errors.Is()` function or the generated `Is...()` helpers to correctly check the error chain.
- **Creating Vague Errors:** Use specific `Reason` strings (e.g., `USER_NAME_EMPTY` instead of just `INVALID_INPUT`) to make your errors easier to debug and handle programmatically.

## Related Skills

- **kratos-api-definition**: Errors are a core part of your API contract and are defined alongside your services and messages in Protobuf.
- **kratos-middleware**: The `recovery` middleware is responsible for catching panics and converting them into Kratos errors, which are then returned to the client.
