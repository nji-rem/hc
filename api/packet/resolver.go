package packet

type Resolver interface {
	SetPackets(packets []Packet)
	Get(header string) (Packet, error)
}
