package connection

import (
	apiSocket "hc/api/connection"
)

type Repository struct {
	connectionHandlers []apiSocket.ConnFunc
	trafficHandlers    []apiSocket.TrafficHandlerFunc
}

func (r *Repository) ConnectionHandlers() []apiSocket.ConnFunc {
	return r.connectionHandlers
}

func (r *Repository) TrafficHandlers() []apiSocket.TrafficHandlerFunc {
	return r.trafficHandlers
}

func NewRepository(connectionHandlers []apiSocket.ConnFunc, trafficHandlers []apiSocket.TrafficHandlerFunc) *Repository {
	return &Repository{
		connectionHandlers: connectionHandlers,
		trafficHandlers:    trafficHandlers,
	}
}
