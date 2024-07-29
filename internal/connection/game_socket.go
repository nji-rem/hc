package connection

import (
	"github.com/panjf2000/gnet/v2"
	"github.com/rs/zerolog/log"
	apiSocket "hc/api/connection"
)

type GameSocket struct {
	gnet.BuiltinEventEngine

	Repository     apiSocket.Repository
	TrafficManager TrafficManager
}

func (g *GameSocket) OnBoot(engine gnet.Engine) gnet.Action {
	log.Info().Msg("Game server started successfully")

	return gnet.None
}

func (g *GameSocket) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	for _, handler := range g.Repository.ConnectionHandlers() {
		if err := handler(c); err != nil {
			return nil, gnet.Close
		}
	}

	return
}

func (g *GameSocket) OnClose(conn gnet.Conn, err error) (action gnet.Action) {
	if err != nil {
		log.Error().Msgf("Connection %s closed due to an message: %s", conn.RemoteAddr().String(), err.Error())
	}

	log.Info().Msgf("Connection %s closed", conn.RemoteAddr().String())

	for _, handler := range g.Repository.ShutdownHandlers() {
		if err := handler(conn); err != nil {
			log.Error().Msgf("an message occurred while running a shutdown handler: %s", err.Error())
		}
	}

	return
}

func (g *GameSocket) OnTraffic(c gnet.Conn) gnet.Action {
	if err := g.TrafficManager.OrchestrateTraffic(c); err != nil {
		log.Error().Msgf("Unable to handle game server viewmodel: %s", err.Error())
		return gnet.Close
	}

	return gnet.None
}
