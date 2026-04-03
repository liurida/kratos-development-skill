---
id: log
title: Logger
description: Kratos contains only the simplest Log interface for business-adapted log access. When your business logic needs to use custom logs inside the kratos framework, you only need to implement the Log method simply.
keywords:
  - Go
  - Kratos
  - Toolkit
  - Framework
  - Microservices
  - Protobuf
  - gRPC
  - HTTP
---

We can use logs to observe program behavior, diagnose problems, or configure corresponding alarms. And defining a well structured log can improve search efficiency and facilitate handling of problems.

## Design concept

For convenience, Kratos defines two levels of abstraction. Logger unifies the access mode of logs and helper interface unifies the call mode of logstore.

Different companies and infrastructures may have different requirements for the printing method, format and output location of logs. Kratos abstracts the log component in order to adapt and migrate to various environments more flexibly, so that the use of logs in business code can be isolated from the specific implementation of the underlying log, and the overall maintainability can be improved.

Log of kratos has the following characteristics:

- Logger is used to connect various log libraries or log platforms, which can be implemented by off the shelf or by yourself.
- Helper is actually called in your project code, it is used to print logs in business code
- Filter is used to filter or modify the output log (usually used for log desensitization)
- Valuer is used to bind some global fixed or dynamic values (such as timestamp, traceID or instance ID) to the output log.

### Helper - log in project code

[Helper](https://github.com/go-kratos/kratos/blob/main/log/helper.go):Advanced log interface, which provides a series of help functions with log levels and formatting methods. This is usually recommended in business logic, which can simplify log code.
You can think of it as a wrapper for the logger, which simplifies the parameters that need to be passed in when printing.
Its usage is basically the following, and the specific usage will be introduced later

```go
helper.Info("hello")
helper.Errorf("hello %s", "eric")
```

### Logger - adapts to various log output methods

[Logger](https://github.com/go-kratos/kratos/blob/main/log/log.go):This is the underlying log interface, which is used to quickly adapt various log libraries to the framework. Only the simplest Log method needs to be implemented.
