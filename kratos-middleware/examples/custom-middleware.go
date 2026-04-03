package main

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

func MyMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				// Do something on entering
				tr.RequestHeader().Set("X-Middleware", "MyMiddleware")
				defer func() {
					// Do something on exiting
				}()
			}
			return handler(ctx, req)
		}
	}
}
