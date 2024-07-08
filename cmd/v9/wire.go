//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	routing2 "hc/api/routing"
	apiSocket "hc/api/socket"
	"hc/cmd/v9/connection"
	"hc/pkg/routing"
	"hc/pkg/socket"
)

var AppSet = wire.NewSet(
	socket.GameServerSet,
	routing.RouteSet,
)

func NewApp(gameConfigurator apiSocket.Configurator, server *socket.GameServer, routeRepository *routing.Repository) *App {
	// Configure game server
	gameConfigurator.Configure(func(connectionHandlers *[]apiSocket.ConnectionHandlerFunc, trafficHandlers *[]apiSocket.TrafficHandlerFunc) {
		*connectionHandlers = append(*connectionHandlers, connection.SayHelloToClientHandler)

		*trafficHandlers = append(*trafficHandlers, connection.PacketHandler{}.Handle)
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
