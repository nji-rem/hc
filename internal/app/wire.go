//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"hc/internal/socket"
)

func NewSocketServer(addr string, packetHandler socket.PacketHandler) *socket.GameServer {
	return socket.NewGameServer(addr, packetHandler.Handle)
}

func NewPacketHandler() socket.PacketHandler {
	return socket.PacketHandler{}
}

func NewApp(gameServer *socket.GameServer, packetHandler socket.PacketHandler) *App {
	return &App{GameServer: gameServer}
}

func InitializeApp(addr string) (*App, error) {
	wire.Build(NewApp, NewSocketServer, NewPacketHandler)

	return &App{}, nil
}
