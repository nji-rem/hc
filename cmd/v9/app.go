package main

import (
	"github.com/panjf2000/gnet/v2"
	"github.com/rs/zerolog/log"
	apiSocket "hc/api/socket"
	"hc/cmd/v9/onconnection"
)

type App struct {
	// GameServer contains an instance to the game server socket runner.
	GameServer gnet.EventHandler

	// GameConfigurator makes it possible to configure connection and traffic handlers.
	GameConfigurator apiSocket.Configurator
}

func (a *App) Bootstrap(addr string) error {
	// Register the game server socket.
	a.GameConfigurator.Configure(func(connectionHandlers *[]apiSocket.ConnectionHandlerFunc, trafficHandlers *[]apiSocket.TrafficHandlerFunc) {
		// Register connection handlers
		*connectionHandlers = append(*connectionHandlers, onconnection.SayHelloToClientHandler)
	})

	log.Info().Msgf("Starting game server on address %s", addr)

	if err := gnet.Run(a.GameServer, addr, gnet.WithMulticore(true)); err != nil {
		log.Error().Msgf("unable to execute game server: %s", err.Error())
		return err
	}

	return nil
}
