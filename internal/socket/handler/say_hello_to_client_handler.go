package handler

import (
	"bytes"
	"github.com/panjf2000/gnet/v2"
	"github.com/rs/zerolog/log"
	"hc/internal/encoding/base64"
)

// SayHelloToClientHandler takes care of sending the "HELLO" packet to the client when the client connects to the
// server. It does nothing more than that.
func SayHelloToClientHandler(c gnet.Conn) error {
	var buf bytes.Buffer

	buf.WriteString(base64.Encode(0)) // @@
	buf.WriteByte(1)                  //
	buf.WriteByte(0)

	if _, err := c.Write(buf.Bytes()); err != nil {
		return err
	}

	log.Info().Msgf("Sent HELLO (%s) to the client", buf.Bytes())

	return nil
}
