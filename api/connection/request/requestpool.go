package request

import (
	"context"
	"github.com/google/uuid"
	"sync"
)

type Pool struct{}

var pool = sync.Pool{New: func() any {
	return new(Bag)
}}

func (r *Pool) Acquire(sessionId string) *Bag {
	request, ok := pool.Get().(*Bag)
	if !ok {
		return new(Bag)
	}

	request.ID = uuid.NewString()
	request.SessionID = sessionId
	request.Context = context.Background()

	return request
}

func (r *Pool) Release(request *Bag) {
	request.ID = ""
	request.SessionID = ""
	request.Header = ""
	request.Body = nil
	request.Context = nil

	pool.Put(request)
}
