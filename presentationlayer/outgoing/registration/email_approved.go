package registration

import "hc/pkg/packet"

type EmailApprovedResponse struct{}

func (e EmailApprovedResponse) Body() []byte {
	packetWriter := packet.AcquireWriter()
	defer packet.ReleaseWriter(packetWriter)

	packetWriter.AppendHeader(271)

	b, _ := packetWriter.Bytes()
	return b
}
