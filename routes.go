package main

import (
	"fmt"
	"hc/api/connection"
	"hc/api/packet"
	secretkey2 "hc/messaging/incoming/handshake/secretkey"
	"io"
)

func CollectRoutes() []packet.Packet {
	return []packet.Packet{
		// Unauthenticated routes
		{
			Name:    "CJ",
			Handler: secretkey2.HandleSecretKey,

			Middleware: []packet.MiddlewareFunc{
				secretkey2.SendClothesPackMiddleware{}.Handle,
				secretkey2.SessionDataMiddleware{}.Handle,
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
