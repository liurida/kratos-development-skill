---
id: errors
title: Errors
---
You can define errors in protos and generate enums with protoc-gen-go.


Error in errors pacakge implements GRPCStatus() interface,The conversion between grpc and HTTP error code is realized, and the business reason is returned through errorinfo.

```json
{
    // The error code is consistent with HTTP status and can be converted into grpc status in grpc.
    "code": 500,
    // The error reason is defined as the business judgment error code.
    "reason": "USER_NOT_FOUND",
    // Error information is user-readable information and can be used as user prompt content.
    "message": "invalid argument error",
    // Error meta information, add additional extensible information for the error.
    "metadata": {
      "foo": "bar"
    }
}
```

### Installation

```bash
go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2
```

### Error Defination

api/helloworld/v1/helloworld.proto

```protobuf
syntax = "proto3";

package api.kratos.v1;
import "errors/errors.proto";

// Define the package name for source code reference.
option go_package = "kratos/api/helloworld;helloworld";
option java_multiple_files = true;
option java_package = "api.helloworld";

enum ErrorReason {
  // Set default error code.
  option (errors.default_code) = 500;

  // Set the error code separately for an enumeration.
  USER_NOT_FOUND = 0 [(errors.code) = 404];

  CONTENT_MISSING = 1 [(errors.code) = 400];
}
```

### Error Generation

To generate code with protoc.

```bash
protoc --proto_path=. \
         --proto_path=./third_party \
         --go_out=paths=source_relative:. \
         --go-errors_out=paths=source_relative:. \
         $(API_PROTO_FILES)
```

Or use the makefile directive at the root of the project
```
make errors
```

### Usage
```go
import "kratos/api/helloworld"

err := wrong()

if errors.Is(err,errors.BadRequest("USER_NAME_EMPTY","")) {
// do something
}

e := errors.FromError(err)
if  e.Reason == "USER_NAME_EMPTY" && e.Code == 500 {
// do something
}

if helloworld.IsUserNotFound(err) {
// do something
})
```
