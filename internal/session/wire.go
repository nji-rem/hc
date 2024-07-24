package session

import (
	"github.com/google/wire"
	apiSession "hc/api/session"
	"sync"
)

var Set = wire.NewSet(
	ProvideSessionStore,
	ProvidePool,
	wire.Bind(new(apiSession.Store), new(*Store)),
)

var storeOnce sync.Once
var store *Store

var poolOnce sync.Once
var pool *apiSession.Pool

func ProvideSessionStore() *Store {
	storeOnce.Do(func() {
		store = new(Store)
	})

	return store
}

func ProvidePool() *apiSession.Pool {
	poolOnce.Do(func() {
		pool = new(apiSession.Pool)
	})

	return pool
}
