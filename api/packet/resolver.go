package packet

type Resolver interface {
	Get(header string) (Packet, error)
}
