package routing

import "github.com/panjf2000/gnet/v2"

type Executor interface {
	ExecutePacket(header string, c gnet.Conn, packet []byte) error
}
