package connection

import (
	"fmt"
	"github.com/panjf2000/gnet/v2"
	"github.com/rs/zerolog/log"
	apiRouting "hc/api/routing"
)

type PacketHandler struct {
	Router apiRouting.Executor
}

func (p PacketHandler) Handle(c gnet.Conn, buf []byte) error {
	if len(buf) < 5 {
		return fmt.Errorf("len(packetHello) is too small, got %d", len(buf))
	}

	header := string(buf[3:5])
	packet := buf[5:]

	log.Debug().Msgf("Got packet header %s with contents %x", header, packet)

	if err := p.Router.ExecutePacket(header, c, packet); err != nil {
		log.Warn().Msgf("unable to execute packet: %s", err.Error())
		return nil
	}

	return nil
}
