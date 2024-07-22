package registration

import "hc/pkg/packet"

type PasswordApprovedResponse struct {
	StatusCode int
}

func (p PasswordApprovedResponse) Body() []byte {
	packetWriter := packet.AcquireWriter()
	defer packet.ReleaseWriter(packetWriter)

	packetWriter.AppendHeader(282) // DZ
	packetWriter.AppendInt(p.StatusCode)

	b, _ := packetWriter.Bytes()
	return b
}
