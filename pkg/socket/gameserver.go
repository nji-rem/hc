package socket

import (
	"github.com/panjf2000/gnet/v2"
	"github.com/rs/zerolog/log"
	apiSocket "hc/api/socket"
)

type GameServer struct {
	gnet.BuiltinEventEngine

	Repository apiSocket.Repository
}

func (g *GameServer) OnBoot(engine gnet.Engine) gnet.Action {
	log.Info().Msg("Game server started successfully")

	return gnet.None
}

func (g *GameServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	for _, handler := range g.Repository.ConnectionHandlers() {
		if err := handler(c); err != nil {
			return nil, gnet.Close
		}
	}

	return
}

func (g *GameServer) OnTraffic(c gnet.Conn) gnet.Action {
	buf, err := c.Next(-1)
	if err != nil {
		log.Error().Msgf("unable to read packet, closing onconnection. err: %s", err.Error())
		return gnet.Close
	}

	for _, handler := range g.Repository.TrafficHandlers() {
		if err := handler(c, buf); err != nil {
			log.Error().Msgf("an error occurred while handling packet: %s", err.Error())
			return gnet.Close
		}
	}

	return gnet.None
}
