package packet

import (
	"fmt"
	"hc/api/packet"
)

type Resolver struct {
	packets map[string]packet.Packet
}

func (r *Resolver) Get(header string) (packet.Packet, error) {
	route, ok := r.packets[header]
	if !ok {
		return packet.Packet{}, fmt.Errorf("no route for header %s", header)
	}

	return route, nil
}

// SetPackets sets the packets map.
//
// Please note that this function is far from thread-safe. It's not intended to be thread-safe.
func (r *Resolver) SetPackets(packets []packet.Packet) {
	items := make(map[string]packet.Packet, len(packets))
	for _, item := range packets {
		items[item.Name] = item
	}

	r.packets = items
}
