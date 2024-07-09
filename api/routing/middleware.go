package routing

import (
	"context"
	"github.com/panjf2000/gnet/v2"
)

type (
	// HandlerFunc contains the function signature that is required for packet handlers.
	HandlerFunc func(ctx context.Context, c gnet.Conn, packet any) error

	// MiddlewareFunc contains the function signature that is required for middleware. Each middleware is wrapped around
	// a packet handler, HandlerFunc.
	MiddlewareFunc func(next HandlerFunc) HandlerFunc
)
