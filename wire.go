//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	apiPacket "hc/api/packet"
	"hc/internal/connection"
	"hc/internal/packet"
	"sync"
)

var AppSet = wire.NewSet(
	connection.GameServerSet,
	ProvideRouteResolver,
	wire.Bind(new(apiPacket.Resolver), new(*packet.Resolver)),
	ProvideConfig,
)

var (
	routeResolver     *packet.Resolver
	routeResolverOnce sync.Once

	config     *viper.Viper
	configOnce sync.Once
)

func ProvideRouteResolver() *packet.Resolver {
	routeResolverOnce.Do(func() {
		routeResolver = packet.NewResolver(CollectRoutes())
	})

	return routeResolver
}

func ProvideConfig() *viper.Viper {
	configOnce.Do(func() {
		v, err := config.Build(
			config.WithConfigDirectory("config/"),
			config.WithEnvFile(".env"))

		if err != nil {
			log.Fatal().Msgf("unable to initialize config: %s", err.Error())
		}

		config = v
	})
	
	return v
}

func NewApp(server *connection.GameSocket) *App {
	return &App{GameServer: server}
}

func InitializeApp() *App {
	wire.Build(NewApp, AppSet)

	return &App{}
}
