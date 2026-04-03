```go
package metrics

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

// Option is metrics option.
type Option func(*options)

// WithRequests with requests counter.
func WithRequests(c metric.Int64Counter) Option {
	return func(o *options) {
		o.requests = c
	}
}

// WithSeconds with seconds histogram.
func WithSeconds(histogram metric.Float64Histogram) Option {
	return func(o *options) {
		o.seconds = histogram
	}
}

type options struct {
	requests metric.Int64Counter
	seconds  metric.Float64Histogram
}

// Server is middleware server-side metrics.
func Server(opts ...Option) middleware.Middleware {
	op := options{}
	for _, o := range opts {
		o(&op)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			startTime := time.Now()
			var ()

			reply, err := handler(ctx, req)

			if op.requests != nil {
				// ...
			}
			if op.seconds != nil {
				// ...
			}

			return reply, err
		}
	}
}
```