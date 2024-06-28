//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"hc/internal/socket"
	"hc/internal/socket/handler"
)

func NewSocketServer(addr string, packetHandler handler.PacketHandler) *socket.GameServer {
	return socket.NewGameServer(addr, packetHandler.Handle)
}

func NewPacketHandler() handler.PacketHandler {
	return handler.PacketHandler{}
}

func NewApp(gameServer *socket.GameServer, packetHandler handler.PacketHandler) *App {
	return &App{GameServer: gameServer}
}

func InitializeApp(addr string) (*App, error) {
	wire.Build(NewApp, NewSocketServer, NewPacketHandler)

	return &App{}, nil
}
