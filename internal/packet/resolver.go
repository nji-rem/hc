package packet

import (
	"fmt"
	"hc/api/packet"
)

type Resolver struct {
	packets map[string]packet.Packet
}

func (r Resolver) Get(header string) (packet.Packet, error) {
	route, ok := r.packets[header]
	if !ok {
		return packet.Packet{}, fmt.Errorf("no route for header %s", header)
	}

	return route, nil
}

func NewResolver(packets []packet.Packet) *Resolver {
	items := make(map[string]packet.Packet, len(packets))
	for _, item := range packets {
		items[item.Name] = item
	}

	return &Resolver{
		packets: items,
	}
}
