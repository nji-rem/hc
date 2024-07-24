package request

import (
	"github.com/google/uuid"
	"sync"
)

type Pool struct{}

var (
	pool = sync.Pool{New: func() any {
		return new(Bag)
	}}

	requestBodyPool = sync.Pool{New: func() any {
		return new(Body)
	}}
)

func (r *Pool) Acquire() *Bag {
	request, ok := pool.Get().(*Bag)
	if !ok {
		return new(Bag)
	}

	request.ID = uuid.NewString()
	request.Body = requestBodyPool.Get().(*Body)

	return request
}

func (r *Pool) Release(request *Bag) {
	request.Body.SetParsedBody(nil)
	request.Body.SetRaw(nil)

	requestBodyPool.Put(request.Body)

	request.ID = ""
	request.Header = ""
	request.Body = nil

	pool.Put(request)
}
