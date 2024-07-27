package main

import (
	"github.com/rs/zerolog/log"
	"hc/api/connection"
	"hc/api/connection/request"
	"hc/api/packet"
	"hc/presentationlayer/event/incoming"
	secretkey2 "hc/presentationlayer/event/incoming/handshake/secretkey"
	"hc/presentationlayer/event/incoming/login"
	"hc/presentationlayer/event/incoming/registration/register/middleware"
	"hc/presentationlayer/outgoing/registration"
	"hc/presentationlayer/outgoing/user/credits"
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
			Handler: secretkey2.HandleSecretKey,

			// tbh I have no idea, todo: check lingo code.
			Middleware: []packet.MiddlewareFunc{
				secretkey2.SendClothesPackMiddleware{}.Handle,
				secretkey2.SessionDataMiddleware{}.Handle,
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
				InitializeLoginAfterRegistrationMiddleware().Handle,
				InitializeValidateUsernameMiddleware().Handle,
				middleware.ParseRegisterRequest,
			},
		},
		{
			Name:    incoming.TryLogin,
			Handler: InitializeTryLoginHandler().Handle,
			Middleware: []packet.MiddlewareFunc{
				login.ParseTryLoginMiddleware,
			},
		},
		{
			Name:    incoming.UserInfo,
			Handler: InitializeInfoRetrieveHandler().Handle,
		},
		{
			Name: incoming.CreditsBalance,
			Handler: func(sessionId string, request *request.Bag, response chan<- connection.Response) error {
				response <- credits.BalanceResponse{Credits: 666}
				return nil
			},
		},
		{
			Name:    incoming.Update,
			Handler: InitializeUpdateUserHandler().Handle,
			Middleware: []packet.MiddlewareFunc{
				// we can re-use this for update lol
				middleware.ParseRegisterRequest,
			},
		},
		{
			Name: incoming.BadgeData,
			Handler: func(sessionId string, request *request.Bag, response chan<- connection.Response) error {
				return nil
			},
		},
		{
			Name: "BV", // navigator packet, WIP -   tCmds.setaProp("NAVIGATE", 150)
			Handler: func(sessionId string, request *request.Bag, response chan<- connection.Response) error {
				return nil
			},
		},
		{
			Name: "BW", // navigator packet, WIP -   tCmds.setaProp("GETUSERFLATCATS", 151)
			Handler: func(sessionId string, request *request.Bag, response chan<- connection.Response) error {
				return nil
			},
		},
		{
			Name: "@P",
			Handler: func(sessionId string, request *request.Bag, response chan<- connection.Response) error {
				return nil
			},
		},
		{
			Name: "@]",
			Handler: func(sessionId string, request *request.Bag, response chan<- connection.Response) error {
				return nil
			},
		},
	}
}
