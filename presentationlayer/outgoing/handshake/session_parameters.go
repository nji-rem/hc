package handshake

import (
	"hc/internal/packet"
	"io"
)

type SessionParametersComposer struct {
	PacketAcquirer func() *packet.Writer
	PacketReleaser func(writer *packet.Writer)
}

func (s SessionParametersComposer) WriteTo(writer io.Writer) (n int64, err error) {
	packetWriter := s.PacketAcquirer()
	defer s.PacketReleaser(packetWriter)

	if err := packetWriter.AppendHeader(257); err != nil {
		return 0, err
	}

	return packetWriter.WriteTo(writer)
}

func NewSessionParametersComposer() SessionParametersComposer {
	return SessionParametersComposer{
		PacketAcquirer: packet.AcquireWriter,
		PacketReleaser: packet.ReleaseWriter,
	}
}
