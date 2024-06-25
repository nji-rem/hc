package socket

import (
	"fmt"
	"github.com/panjf2000/gnet/v2"
)

type PacketHandler struct {
}

func (p PacketHandler) Handle(c gnet.Conn, buf []byte) error {
	fmt.Println("Reached PacketHandler")

	return nil
}
