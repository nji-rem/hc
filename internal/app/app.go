package app

import (
	"github.com/panjf2000/gnet/v2"
	"github.com/rs/zerolog/log"
	"hc/internal/socket"
)

type App struct {
	// GameServer contains the game server socket. It's already instantiated with a game server
	// address and should be accessible.
	GameServer *socket.GameServer
}

func (a *App) Start() error {
	// This is a blocking call
	if err := gnet.Run(a.GameServer, a.GameServer.Addr, gnet.WithMulticore(true)); err != nil {
		log.Error().Msgf("unable to execute game server: %s", err.Error())
		return err
	}

	return nil
}
