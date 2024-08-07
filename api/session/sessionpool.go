package session

import (
	"github.com/google/uuid"
	"sync"
)

type Pool struct{}

var pool = sync.Pool{New: func() any {
	return new(Bag)
}}

func (r *Pool) Acquire() *Bag {
	bag, ok := pool.Get().(*Bag)
	if !ok {
		bag = new(Bag)
	}

	r.clean(bag)
	bag.ID = uuid.NewString()

	return bag
}

func (r *Pool) clean(bag *Bag) {
	bag.ID = ""
	bag.Authenticated.Store(false)
	bag.Clear()
}

func (r *Pool) Release(request *Bag) {
	r.clean(request)

	pool.Put(request)
}
