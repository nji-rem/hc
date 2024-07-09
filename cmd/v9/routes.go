package main

import (
	"hc/api/routing"
	"hc/cmd/v9/messaging/incoming/handshake"
)

func CollectRoutes() []routing.Route {
	return []routing.Route{
		// Unauthenticated routes
		{
			Name:    "CJ",
			Handler: handshake.HandleSecretKey,

			Middleware: []routing.MiddlewareFunc{},
		},
	}
}
