package registration

import (
	"hc/pkg/packet"
)

type ApproveNameReply struct {
	NameCheckCode int
}

func (a ApproveNameReply) Body() []byte {
	packetWriter := packet.AcquireWriter()
	defer packet.ReleaseWriter(packetWriter)

	packetWriter.AppendHeader(36)
	packetWriter.AppendInt(a.NameCheckCode)

	b, _ := packetWriter.Bytes()
	return b
}
