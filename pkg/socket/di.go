package socket

import (
	"github.com/google/wire"
	"github.com/panjf2000/gnet/v2"
	apiSocket "hc/api/socket"
	"sync"
)

var repository *Repository
var repositoryOnce sync.Once

var GameServerSet = wire.NewSet(
	ProvideSocketServer,
	ProvideSocketRepository,
	wire.Bind(new(apiSocket.Configurator), new(*Repository)),
)

func ProvideSocketServer(repository *Repository) *GameServer {
	return &GameServer{
		BuiltinEventEngine: gnet.BuiltinEventEngine{},
		Repository:         repository,
	}
}

func ProvideSocketRepository() *Repository {
	repositoryOnce.Do(func() {
		repository = new(Repository)
	})

	return repository
}
