package connection

import (
	"hc/api/connection/request"
	"io"
)

type (
	ConnFunc            func(writer io.Writer) error
	ShutdownHandlerFunc func(writer io.Writer) error
	TrafficHandlerFunc  func(request *request.Bag, writer io.Writer) error

	ConfigFunc func(connectionHandlers *[]ConnFunc, trafficHandlers *[]TrafficHandlerFunc)
)

type Repository interface {
	ConnectionHandlers() []ConnFunc
	ShutdownHandlers() []ShutdownHandlerFunc
	TrafficHandlers() []TrafficHandlerFunc
}

type Configurator interface {
	Configure(fn ConfigFunc)
}
