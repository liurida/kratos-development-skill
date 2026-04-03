
---
id: overview
title: Overview
---

Kratos has a series of built-in middleware to deal with common purpose such as logging or metrics. You could also implement **Middleware** interface to develop your custom middleware to process common business such as the user authentication etc.

## Built-in Middleware

Their codes are located in `middleware` directory.

- `logging`: This middleware is for logging the request.
- `metrics`: This middleware is for enabling metric.
- `recovery`: This middleware is for panic recovery.
- `tracing`: This middleware is for enabling trace.
- `validate`: This middleware is for parameter validation.
- `metadata`: This middleware is for enabling metadata transmission.
- `auth`: This middleware is for authority check using JWT.
- `ratelimit`: This middleware is for traffic control in server side.
- `circuitbreaker`: This middleware is for breaker control in client side.
