package socket

import (
	apiSocket "hc/api/socket"
	"sync"
)

type Repository struct {
	connectionHandlers []apiSocket.ConnectionHandlerFunc
	trafficHandlers    []apiSocket.TrafficHandlerFunc

	once sync.Once
}

func (r *Repository) Configure(fn apiSocket.ConfigFunc) {
	r.once.Do(func() {
		fn(&r.connectionHandlers, &r.trafficHandlers)
	})
}

func (r *Repository) ConnectionHandlers() []apiSocket.ConnectionHandlerFunc {
	return r.connectionHandlers
}

func (r *Repository) TrafficHandlers() []apiSocket.TrafficHandlerFunc {
	return r.trafficHandlers
}
