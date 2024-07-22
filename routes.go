package main

import (
	"hc/api/packet"
	"hc/presentationlayer/incoming"
	"hc/presentationlayer/incoming/handshake/secretkey"
)

func CollectRoutes() []packet.Packet {
	return []packet.Packet{
		{
			Name:    incoming.SecretKey,
			Handler: secretkey.HandleSecretKey,

			Middleware: []packet.MiddlewareFunc{
				secretkey.SendClothesPackMiddleware{}.Handle,
				secretkey.SessionDataMiddleware{}.Handle,
			},
		},
		{
			Name:    incoming.NameCheck,
			Handler: InitializeNameCheckHandler().Handle,

			Middleware: []packet.MiddlewareFunc{},
		},
		{
			Name:       incoming.PasswordCheck,
			Handler:    InitializePasswordVerifyHandler().Handle,
			Middleware: []packet.MiddlewareFunc{},
		},
	}
}
