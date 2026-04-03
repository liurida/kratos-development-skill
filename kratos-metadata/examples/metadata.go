```go
package metadata

import (
	"context"
	"fmt"
	"strings"
)

// Metadata is our way of representing request headers internally.
type Metadata map[string][]string

// New creates an MD from a given key-values map.
func New(mds ...map[string][]string) Metadata {
	md := Metadata{}
	for _, m := range mds {
		for k, vList := range m {
			for _, v := range vList {
				md.Add(k, v)
			}
		}
	}
	return md
}

// Add adds the key, value pair to the header.
func (m Metadata) Add(key, value string) {
	if key == "" {
		return
	}

	lowerKey := strings.ToLower(key)
	m[lowerKey] = append(m[lowerKey], value)
}

// Get returns the value associated with the passed key.
func (m Metadata) Get(key string) string {
	v := m[strings.ToLower(key)]
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

// Set stores the key-value pair.
func (m Metadata) Set(key string, value string) {
	if key == "" || value == "" {
		return
	}
	m[strings.ToLower(key)] = []string{value}
}

type serverMetadataKey struct{}

// NewServerContext creates a new context with client md attached.
func NewServerContext(ctx context.Context, md Metadata) context.Context {
	return context.WithValue(ctx, serverMetadataKey{}, md)
}

// FromServerContext returns the server metadata in ctx if it exists.
func FromServerContext(ctx context.Context) (Metadata, bool) {
	md, ok := ctx.Value(serverMetadataKey{}).(Metadata)
	return md, ok
}

type clientMetadataKey struct{}

// NewClientContext creates a new context with client md attached.
func NewClientContext(ctx context.Context, md Metadata) context.Context {
	return context.WithValue(ctx, clientMetadataKey{}, md)
}

// FromClientContext returns the client metadata in ctx if it exists.
func FromClientContext(ctx context.Context) (Metadata, bool) {
	md, ok := ctx.Value(clientMetadataKey{}).(Metadata)
	return md, ok
}

```