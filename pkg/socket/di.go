package socket

import (
	"github.com/google/wire"
	"github.com/panjf2000/gnet/v2"
	"sync"
)

var GameServerSet = wire.NewSet(
	ProvideSocketServer,
	ProvideSocketRepository,
)

var repository *Repository
var repositoryOnce sync.Once

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
