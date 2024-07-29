package message

import "hc/pkg/packet"

type AlertResponse struct {
	Msg string
}

func (a AlertResponse) Body() []byte {
	packetWriter := packet.AcquireWriter()
	defer packet.ReleaseWriter(packetWriter)

	packetWriter.AppendHeader(139)
	packetWriter.AppendString(a.Msg)

	b, _ := packetWriter.Bytes()
	return b
}
