package main

import (
	"fmt"
	"hc/api/connection"
	"hc/api/packet"
	"hc/cmd/v9/messaging/incoming/handshake/secretkey"
	"io"
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
		{
			Name: "@j", // name check
			Handler: func(request *connection.Request, writer io.Writer) error {
				fmt.Println("name check not implemented yet")
				return nil
			},
			Middleware: []packet.MiddlewareFunc{},
		},
	}
}
