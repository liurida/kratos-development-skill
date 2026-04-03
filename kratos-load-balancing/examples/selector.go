package selector

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
)

// ErrNoAvailable is no available node.
var ErrNoAvailable = errors.ServiceUnavailable("no_available_node", "")

// Selector is node pick balancer.
type Selector interface {
	Rebalancer

	// Select nodes
	Select(ctx context.Context, opts ...SelectOption) (selected Node, done DoneFunc, err error)
}

// Rebalancer is nodes rebalancer.
type Rebalancer interface {
	// Apply is apply all nodes when any changes happen
	Apply(nodes []Node)
}

// Builder build selector
type Builder interface {
	Build() Selector
}

// Node is node interface.
type Node interface {
	Scheme() string
	Address() string
	ServiceName() string
	InitialWeight() *int64
	Version() string
	Metadata() map[string]string
}

// DoneFunc is callback function when RPC invoke done.
type DoneFunc func(ctx context.Context, di DoneInfo)
