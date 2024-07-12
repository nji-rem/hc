package main

import (
	"hc/api/packet"
	"hc/cmd/v9/messaging/incoming/handshake"
)

func CollectRoutes() []packet.Packet {
	return []packet.Packet{
		// Unauthenticated routes
		{
			Name:    "CJ",
			Handler: handshake.HandleSecretKey,

			Middleware: []packet.MiddlewareFunc{},
		},
	}
}
