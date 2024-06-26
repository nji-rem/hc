package socket

import (
	"github.com/panjf2000/gnet/v2"
	"github.com/rs/zerolog/log"
)

type TrafficHandlerFunc func(c gnet.Conn, buf []byte) error

type GameServer struct {
	gnet.BuiltinEventEngine

	Addr            string
	trafficHandlers []TrafficHandlerFunc
}

func (g *GameServer) OnBoot(engine gnet.Engine) gnet.Action {
	log.Info().Msgf("Game server started on address %s", g.Addr)

	return gnet.None
}

func (g *GameServer) OnTraffic(c gnet.Conn) gnet.Action {
	buf, err := c.Next(-1)
	if err != nil {
		log.Error().Msgf("unable to read packet, closing connection. err: %s", err.Error())
		return gnet.Close
	}

	for _, handler := range g.trafficHandlers {
		if err := handler(c, buf); err != nil {
			log.Error().Msgf("an error occurred while handling packet: %s", err.Error())
			return gnet.Close
		}
	}

	return gnet.None
}

func NewGameServer(addr string, trafficHandlers ...TrafficHandlerFunc) *GameServer {
	return &GameServer{
		BuiltinEventEngine: gnet.BuiltinEventEngine{},
		Addr:               addr,
		trafficHandlers:    trafficHandlers,
	}
}
