//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	apiPacket "hc/api/packet"
	"hc/internal/connection"
	"hc/internal/packet"
	"sync"
)

var AppSet = wire.NewSet(
	connection.GameServerSet,
	ProvideRouteResolver,
	wire.Bind(new(apiPacket.Resolver), new(*packet.Resolver)),
)

var routeResolver *packet.Resolver
var routeResolverOnce sync.Once

func ProvideRouteResolver() *packet.Resolver {
	routeResolverOnce.Do(func() {
		routeResolver = packet.NewResolver(CollectRoutes())
	})

	return routeResolver
}

func NewApp(server *connection.GameSocket) *App {
	return &App{GameServer: server}
}

func InitializeApp() *App {
	wire.Build(NewApp, AppSet)

	return &App{}
}
