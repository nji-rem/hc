package connection

import (
	"github.com/google/wire"
	"github.com/panjf2000/gnet/v2"
	apiSocket "hc/api/connection"
	"hc/api/packet"
	"sync"
)

// singletons
var (
	repository     *Repository
	repositoryOnce sync.Once

	requestPool     *apiSocket.RequestPool
	requestPoolOnce sync.Once
)

var GameServerSet = wire.NewSet(
	ProvideGameSocket,
	ProvideSocketRepository,
	ProvideTrafficManager,
	ProvideTrafficParser,
	ProvideRequestPool,
	ProvideFrontController,
	ProvideMiddlewareWrapper,
)

func ProvideGameSocket(repository *Repository, manager TrafficManager) *GameSocket {
	return &GameSocket{
		BuiltinEventEngine: gnet.BuiltinEventEngine{},
		Repository:         repository,
		TrafficManager:     manager,
	}
}

func ProvideSocketRepository(controller FrontController) *Repository {
	repositoryOnce.Do(func() {
		connectionHandlers := []apiSocket.ConnFunc{
			SayHelloToClientHandler,
		}

		trafficHandlers := []apiSocket.TrafficHandlerFunc{
			controller.Handle,
		}

		repository = NewRepository(connectionHandlers, trafficHandlers)
	})

	return repository
}

func ProvideTrafficManager(trafficRepository *Repository, trafficParser *TrafficParser, pool *apiSocket.RequestPool) TrafficManager {
	return TrafficManager{
		TrafficRepository: trafficRepository,
		TrafficParser:     trafficParser,
		RequestPool:       pool,
	}
}

func ProvideTrafficParser() *TrafficParser {
	return &TrafficParser{}
}

func ProvideRequestPool() *apiSocket.RequestPool {
	requestPoolOnce.Do(func() {
		requestPool = &apiSocket.RequestPool{}
	})

	return requestPool
}

func ProvideFrontController(resolver packet.Resolver, middlewareWrapper packet.WrapFunc) FrontController {
	return FrontController{
		Resolver:       resolver,
		WrapMiddleware: middlewareWrapper,
	}
}

func ProvideMiddlewareWrapper() packet.WrapFunc {
	return func(handler packet.HandlerFunc, middleware []packet.MiddlewareFunc) packet.HandlerFunc {
		for _, item := range middleware {
			handler = item(handler)
		}

		return handler
	}
}
