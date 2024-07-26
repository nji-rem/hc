package credits

import (
	"hc/pkg/packet"
	"strconv"
)

type BalanceResponse struct {
	Credits int
}

func (b BalanceResponse) Body() []byte {
	packetWriter := packet.AcquireWriter()
	defer packet.ReleaseWriter(packetWriter)

	packetWriter.AppendHeader(6)
	packetWriter.AppendString(strconv.Itoa(b.Credits) + ".0")

	buf, _ := packetWriter.Bytes()
	return buf
}
