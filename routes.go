package main

import (
	"github.com/rs/zerolog/log"
	"hc/api/connection"
	"hc/api/connection/request"
	"hc/api/packet"
	"hc/presentationlayer/incoming"
	"hc/presentationlayer/incoming/handshake/secretkey"
	"hc/presentationlayer/incoming/registration/register/middleware"
	"hc/presentationlayer/outgoing/registration"
)

func CollectRoutes() []packet.Packet {
	return []packet.Packet{
		{
			Name: "@q",
			Handler: func(sessionId string, request *request.Bag, response chan<- connection.Response) error {
				log.Warn().Msg("@q not implemented yet")
				return nil
			},

			Middleware: []packet.MiddlewareFunc{},
		},
		{
			Name:    incoming.SecretKey,
			Handler: secretkey.HandleSecretKey,

			// tbh I have no idea, todo: check lingo code.
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
			Handler: func(sessionId string, request *request.Bag, response chan<- connection.Response) error {
				// We won't be implementing the e-mail check.
				response <- registration.EmailApprovedResponse{}

				return nil
			},

			Middleware: []packet.MiddlewareFunc{},
		},
		{
			Name:    incoming.Register,
			Handler: InitializeRegisterHandler().Handle,
			Middleware: []packet.MiddlewareFunc{
				InitializeValidateUsernameMiddleware().Handle,
				middleware.ParseRequestMiddleware,
			},
		},
	}
}
