package session

import (
	"sync"
	"sync/atomic"
)

type Bag struct {
	// ID contains the session id.
	ID string

	// Authenticated tells the caller if this session is authenticated, e.g. to perform certain requests.
	Authenticated atomic.Bool

	// kvStore contains key-value data related to the current session.
	kvStore sync.Map
}

func (b *Bag) Set(key string, value any) {
	b.kvStore.Store(key, value)
}

func (b *Bag) Clear() {
	b.kvStore.Range(func(key, value any) bool {
		b.kvStore.Delete(key)
		return true
	})
}

func (b *Bag) Get(key string) any {
	v, ok := b.kvStore.Load(key)
	if !ok {
		return nil
	}

	return v
}
