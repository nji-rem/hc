package connection

import (
	"github.com/panjf2000/gnet/v2"
	gameserverApi "hc/api/connection"
)

type TrafficManager struct {
	TrafficRepository gameserverApi.Repository
	TrafficParser     gameserverApi.TrafficParser
	RequestPool       *gameserverApi.RequestPool
}

func (t TrafficManager) OrchestrateTraffic(c gnet.Conn) error {
	// Acquire a new request object
	request := t.RequestPool.Acquire()

	// Release the request object when we're done with it.
	defer t.RequestPool.Release(request)

	// Parse the traffic into the request object
	if err := t.TrafficParser.Parse(c, request); err != nil {
		return err
	}

	// Get all traffic handlers and execute each one of them.
	trafficHandlers := t.TrafficRepository.TrafficHandlers()
	for i := 0; i < len(trafficHandlers); i++ {
		if err := trafficHandlers[i](request, c); err != nil {
			return err
		}
	}

	return nil
}
