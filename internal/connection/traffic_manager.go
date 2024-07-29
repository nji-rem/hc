package connection

import (
	"errors"
	"github.com/panjf2000/gnet/v2"
	gameserverApi "hc/api/connection"
	"hc/api/connection/request"
)

type TrafficManager struct {
	TrafficRepository gameserverApi.Repository
	TrafficParser     gameserverApi.TrafficParser
	RequestPool       *request.Pool
}

var ErrContextNotFound = errors.New("connection context is undefined")

func (t TrafficManager) OrchestrateTraffic(c gnet.Conn) error {
	ctx, ok := c.Context().(gameserverApi.Context)
	if !ok {
		return ErrContextNotFound
	}

	// Acquire a new viewmodel object
	request := t.RequestPool.Acquire()

	// Release the viewmodel object when we're done with it.
	defer t.RequestPool.Release(request)

	// Parse the traffic into the viewmodel object
	if err := t.TrafficParser.Parse(c, request); err != nil {
		return err
	}

	// Get all traffic handlers and execute each one of them.
	trafficHandlers := t.TrafficRepository.TrafficHandlers()
	for i := 0; i < len(trafficHandlers); i++ {
		if err := trafficHandlers[i](ctx.SessionID(), request, c); err != nil {
			return err
		}
	}

	return nil
}
