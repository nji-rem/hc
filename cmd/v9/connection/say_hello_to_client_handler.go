package connection

import (
	"github.com/panjf2000/gnet/v2"
	"github.com/rs/zerolog/log"
)

var (
	packetHello = []byte{'@', '@', 1, 0}
)

// SayHelloToClientHandler takes care of sending the "HELLO" packet to the client when the client connects to the
// server.
func SayHelloToClientHandler(c gnet.Conn) error {
	if _, err := c.Write(packetHello); err != nil {
		return err
	}

	log.Info().Msgf("Sent HELLO to %s", c.RemoteAddr().String())

	return nil
}
