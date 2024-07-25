package login

import "hc/pkg/packet"

type OKResponse struct{}

func (l OKResponse) Body() []byte {
	packetWriter := packet.AcquireWriter()
	defer packet.ReleaseWriter(packetWriter)

	_ = packetWriter.AppendHeader(3)

	b, _ := packetWriter.Bytes()
	return b
}
