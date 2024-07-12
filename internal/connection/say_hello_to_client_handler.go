package connection

import (
	"github.com/rs/zerolog/log"
	"io"
)

var (
	packetHello = []byte{'@', '@', 1}
)

// SayHelloToClientHandler takes care of sending the "HELLO" packet to the client when the client connects to the
// server.
func SayHelloToClientHandler(c io.Writer) error {
	if _, err := c.Write(packetHello); err != nil {
		return err
	}

	log.Info().Msg("Sent HELLO to the client")

	return nil
}
