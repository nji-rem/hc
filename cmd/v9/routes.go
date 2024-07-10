package main

import (
	"context"
	"fmt"
	"github.com/panjf2000/gnet/v2"
	"hc/api/routing"
	"hc/cmd/v9/messaging/incoming/handshake"
)

func CollectRoutes() []routing.Route {
	return []routing.Route{
		// Unauthenticated routes
		{
			Name:    "CJ",
			Handler: handshake.HandleSecretKey,

			Middleware: []routing.MiddlewareFunc{
				func(next routing.HandlerFunc) routing.HandlerFunc {
					return func(ctx context.Context, c gnet.Conn, packet any) error {
						fmt.Println("Executing middleware")
						return next(ctx, c, packet)
					}
				},
			},
		},
	}
}
