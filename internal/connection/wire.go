package connection

import (
	"github.com/google/wire"
	"github.com/panjf2000/gnet/v2"
	apiSocket "hc/api/connection"
	"hc/api/connection/request"
	"hc/api/packet"
	"hc/api/session"
	"sync"
)

// singletons
var (
	repository     *Repository
	repositoryOnce sync.Once

	requestPool     *request.Pool
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
	ProvideCreateSessionOnNewConnectionHandler,
	ProvideDeleteSessionOnConnectionDestroyed,
)

func ProvideGameSocket(repository *Repository, manager TrafficManager) *GameSocket {
	return &GameSocket{
		BuiltinEventEngine: gnet.BuiltinEventEngine{},
		Repository:         repository,
		TrafficManager:     manager,
	}
}

func ProvideSocketRepository(controller FrontController, createSessionHandler CreateSessionOnNewConnection, deleteSessionHandler DeleteSessionOnConnectionDestroyed) *Repository {
	repositoryOnce.Do(func() {
		connectionHandlers := []apiSocket.ConnFunc{
			SayHelloToClientHandler,
			createSessionHandler.Handle,
		}

		trafficHandlers := []apiSocket.TrafficHandlerFunc{
			controller.Handle,
		}

		shutdownHandlers := []apiSocket.ShutdownHandlerFunc{
			deleteSessionHandler.Handle,
		}

		repository = NewRepository(connectionHandlers, trafficHandlers, shutdownHandlers)
	})

	return repository
}

func ProvideTrafficManager(trafficRepository *Repository, trafficParser *TrafficParser, pool *request.Pool) TrafficManager {
	return TrafficManager{
		TrafficRepository: trafficRepository,
		TrafficParser:     trafficParser,
		RequestPool:       pool,
	}
}

func ProvideTrafficParser() *TrafficParser {
	return &TrafficParser{}
}

func ProvideRequestPool() *request.Pool {
	requestPoolOnce.Do(func() {
		requestPool = &request.Pool{}
	})

	return requestPool
}

func ProvideCreateSessionOnNewConnectionHandler(sessionStore session.Store, pool *session.Pool) CreateSessionOnNewConnection {
	return CreateSessionOnNewConnection{
		SessionStore: sessionStore,
		Pool:         pool,
	}
}

func ProvideDeleteSessionOnConnectionDestroyed(sessionStore session.Store, pool *session.Pool) DeleteSessionOnConnectionDestroyed {
	return DeleteSessionOnConnectionDestroyed{
		SessionStore: sessionStore,
		Pool:         pool,
	}
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
