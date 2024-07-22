package handshake

import (
	"hc/pkg/packet"
)

type SessionParametersComposer struct{}

func (s SessionParametersComposer) Body() []byte {
	packetWriter := packet.AcquireWriter()
	defer packet.ReleaseWriter(packetWriter)

	_ = packetWriter.AppendHeader(257)

	b, _ := packetWriter.Bytes()
	return b
}

func NewSessionParametersComposer() SessionParametersComposer {
	return SessionParametersComposer{}
}
