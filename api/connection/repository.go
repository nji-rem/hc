package connection

import (
	"io"
)

type (
	ConnFunc           func(writer io.Writer) error
	TrafficHandlerFunc func(request *Request, writer io.Writer) error

	ConfigFunc func(connectionHandlers *[]ConnFunc, trafficHandlers *[]TrafficHandlerFunc)
)

type Repository interface {
	ConnectionHandlers() []ConnFunc
	TrafficHandlers() []TrafficHandlerFunc
}

type Configurator interface {
	Configure(fn ConfigFunc)
}
