package connection

import (
	apiSocket "hc/api/connection"
)

type Repository struct {
	connectionHandlers []apiSocket.ConnFunc
	trafficHandlers    []apiSocket.TrafficHandlerFunc
	shutdownHandlers   []apiSocket.ShutdownHandlerFunc
}

func (r *Repository) ConnectionHandlers() []apiSocket.ConnFunc {
	return r.connectionHandlers
}

func (r *Repository) TrafficHandlers() []apiSocket.TrafficHandlerFunc {
	return r.trafficHandlers
}

func (r *Repository) ShutdownHandlers() []apiSocket.ShutdownHandlerFunc {
	return r.shutdownHandlers
}

func NewRepository(connectionHandlers []apiSocket.ConnFunc, trafficHandlers []apiSocket.TrafficHandlerFunc, shutdownHandlers []apiSocket.ShutdownHandlerFunc) *Repository {
	return &Repository{
		connectionHandlers: connectionHandlers,
		trafficHandlers:    trafficHandlers,
		shutdownHandlers:   shutdownHandlers,
	}
}
