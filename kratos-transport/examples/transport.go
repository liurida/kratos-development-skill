package transport

import (
	"context"
	"net/url"
)

// Server is transport server.
type Server interface {
	Start(context.Context) error
	Stop(context.Context) error
}

// Endpointer is registry endpoint.
type Endpointer interface {
	Endpoint() (*url.URL, error)
}

// Header is the storage medium used by a Header.
type Header interface {
	Get(key string) string
	Set(key string, value string)
	Add(key string, value string)
	Keys() []string
	Values(key string) []string
}

// Transporter is transport context value interface.
type Transporter interface {
	Kind() Kind
	Endpoint() string
	Operation() string
	RequestHeader() Header
	ReplyHeader() Header
}

// Kind defines the type of Transport
type Kind string

const (
	KindGRPC Kind = "grpc"
	KindHTTP Kind = "http"
)

type serverTransportKey struct{}

// NewServerContext returns a new Context that carries value.
func NewServerContext(ctx context.Context, tr Transporter) context.Context {
	return context.WithValue(ctx, serverTransportKey{}, tr)
}

// FromServerContext returns the Transport value stored in ctx, if any.
func FromServerContext(ctx context.Context) (tr Transporter, ok bool) {
	tr, ok = ctx.Value(serverTransportKey{}).(Transporter)
	return
}
