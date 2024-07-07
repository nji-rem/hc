//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	apiSocket "hc/api/socket"
	"hc/cmd/v9/onconnection"
	"hc/pkg/socket"
)

var AppSet = wire.NewSet(
	socket.GameServerSet,
)

func NewApp(gameConfigurator apiSocket.Configurator, server *socket2.GameServer) *App {
	// Configure game server
	gameConfigurator.Configure(func(connectionHandlers *[]apiSocket.ConnectionHandlerFunc, trafficHandlers *[]apiSocket.TrafficHandlerFunc) {
		*connectionHandlers = append(*connectionHandlers, onconnection.SayHelloToClientHandler)
	})

	// Load routes
	CollectRoutes()
	return &App{GameServer: server}
}

func InitializeApp() *App {
	wire.Build(NewApp, AppSet)

	return &App{}
}
