package main

import (
	"hc/api/connection"
	"hc/api/packet"
	"hc/presentationlayer/incoming"
	"hc/presentationlayer/incoming/handshake/secretkey"
	"hc/presentationlayer/outgoing/registration"
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
		{
			Name: incoming.MailCheck,
			Handler: func(request *connection.Request, response chan<- connection.Response) error {
				// We won't be implementing the e-mail check.
				response <- registration.EmailApprovedResponse{}

				return nil
			},

			Middleware: []packet.MiddlewareFunc{},
		},
	}
}
