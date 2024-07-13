package main

import (
	"hc/api/packet"
	"hc/cmd/v9/messaging/incoming/handshake/secretkey"
)

func CollectRoutes() []packet.Packet {
	return []packet.Packet{
		// Unauthenticated routes
		{
			Name:    "CJ",
			Handler: secretkey.HandleSecretKey,

			Middleware: []packet.MiddlewareFunc{
				secretkey.SendClothesPackMiddleware{}.Handle,
				secretkey.SessionDataMiddleware{}.Handle,
			},
		},
	}
}
