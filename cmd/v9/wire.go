//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	routing2 "hc/api/routing"
	apiSocket "hc/api/socket"
	"hc/cmd/v9/connection"
	"hc/internal/routing"
	socket2 "hc/internal/socket"
)

var AppSet = wire.NewSet(
	socket2.GameServerSet,
	routing.RouteSet,
	ProvidePacketHandler,
)

func ProvidePacketHandler(router *routing.RouteExecutor) connection.PacketHandler {
	return connection.PacketHandler{
		Router: router,
	}
}

func NewApp(gameConfigurator apiSocket.Configurator, server *socket2.GameServer, routeRepository *routing.Repository, handler connection.PacketHandler) *App {
	// Configure game server
	gameConfigurator.Configure(func(connectionHandlers *[]apiSocket.ConnectionHandlerFunc, trafficHandlers *[]apiSocket.TrafficHandlerFunc) {
		*connectionHandlers = append(*connectionHandlers, connection.SayHelloToClientHandler)

		*trafficHandlers = append(*trafficHandlers, handler.Handle)
	})

	// Configure routes
	collectedRoutes := CollectRoutes()
	routeMap := make(map[string]routing2.Route, len(collectedRoutes))
	for _, route := range CollectRoutes() {
		routeMap[route.Name] = route
	}

	routeRepository.Routes = routeMap

	return &App{GameServer: server}
}

func InitializeApp() *App {
	wire.Build(NewApp, AppSet)

	return &App{}
}
