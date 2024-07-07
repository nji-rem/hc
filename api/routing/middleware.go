package routing

import (
	"hc/api/routing/request"
)

type (
	// HandlerFunc contains the function signature that is required for packet handlers.
	HandlerFunc func(ctx request.Context, packet any) error

	// MiddlewareFunc contains the function signature that is required for middleware. Each middleware is wrapped around
	// a packet handler, HandlerFunc.
	MiddlewareFunc func(next HandlerFunc) HandlerFunc
)
