package main

import (
	"fmt"
	"hc/api/connection"
	"hc/api/packet"
	"hc/presentationlayer/incoming"
	"hc/presentationlayer/incoming/handshake/secretkey"
	"io"
)

func CollectRoutes() []packet.Packet {
	return []packet.Packet{
		// Unauthenticated routes
		{
			Name:    incoming.SecretKey,
			Handler: secretkey.HandleSecretKey,

			Middleware: []packet.MiddlewareFunc{
				secretkey.SendClothesPackMiddleware{}.Handle,
				secretkey.SessionDataMiddleware{}.Handle,
			},
		},
		{
			Name: incoming.NameCheck,
			Handler: func(request *connection.Request, writer io.Writer) error {
				fmt.Println("name check not implemented yet")
				return nil
			},
			
			Middleware: []packet.MiddlewareFunc{},
		},
	}
}
