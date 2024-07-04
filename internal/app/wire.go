//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"hc/internal/socket"
	"hc/internal/socket/handler"
)

func NewSocketServer(addr string, newConnectionHandlers []socket.ConnectionHandlerFunc, packetHandler handler.PacketHandler) *socket.GameServer {
	return socket.NewGameServer(addr, newConnectionHandlers, packetHandler.Handle)
}

func NewConnectionHandler() []socket.ConnectionHandlerFunc {
	return []socket.ConnectionHandlerFunc{handler.SayHelloToClientHandler}
}

func NewPacketHandler() handler.PacketHandler {
	return handler.PacketHandler{}
}

func NewApp(gameServer *socket.GameServer, packetHandler handler.PacketHandler) *App {
	return &App{GameServer: gameServer}
}

func InitializeApp(addr string) (*App, error) {
	wire.Build(NewApp, NewSocketServer, NewConnectionHandler, NewPacketHandler)

	return &App{}, nil
}
