package socket

import "github.com/panjf2000/gnet/v2"

type (
	ConnectionHandlerFunc func(c gnet.Conn) error
	TrafficHandlerFunc    func(c gnet.Conn, buf []byte) error

	ConfigFunc func(connectionHandlers *[]ConnectionHandlerFunc, trafficHandlers *[]TrafficHandlerFunc)
)

type Repository interface {
	ConnectionHandlers() []ConnectionHandlerFunc
	TrafficHandlers() []TrafficHandlerFunc
}

type Configurator interface {
	Configure(fn ConfigFunc)
}
