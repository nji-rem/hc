package routing

import (
	"github.com/google/wire"
	"sync"
)

var RouteSet = wire.NewSet(ProvideRouteExecutor, ProvideRepository)

var repository *Repository
var repositoryOnce sync.Once

func ProvideRepository() *Repository {
	repositoryOnce.Do(func() {
		repository = new(Repository)
	})

	return repository
}

func ProvideRouteExecutor(repository *Repository) RouteExecutor {
	return RouteExecutor{
		Repository: repository,
	}
}
