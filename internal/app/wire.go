//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"hc/internal/socket"
)

func NewSocketServer(addr string) *socket.GameServer {
	return &socket.GameServer{Addr: addr}
}

func NewApp(gameServer *socket.GameServer) *App {
	return &App{GameServer: gameServer}
}

func InitializeApp(addr string) (*App, error) {
	wire.Build(NewApp, NewSocketServer)

	return &App{}, nil
}
