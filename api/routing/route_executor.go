package routing

type Executor interface {
	ExecutePacket(header string, packet []byte) error
}
