//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	socket2 "hc/pkg/socket"
)

var AppSet = wire.NewSet(
	socket2.GameServerSet,
)

func NewApp(gameServer *socket2.GameServer, repository *socket2.Repository) *App {
	return &App{
		GameServer:       gameServer,
		GameConfigurator: repository,
	}
}

func InitializeApp() (*App, error) {
	wire.Build(NewApp, AppSet)

	return &App{}, nil
}
