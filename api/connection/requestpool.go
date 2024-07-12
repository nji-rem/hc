package connection

import (
	"context"
	"sync"
)

type RequestPool struct{}

var pool = sync.Pool{New: func() any {
	return new(Request)
}}

func (r *RequestPool) Acquire() *Request {
	request, ok := pool.Get().(*Request)
	if !ok {
		return new(Request)
	}

	request.Context = context.Background()

	return request
}

func (r *RequestPool) Release(request *Request) {
	request.Header = ""
	request.Body = nil
	request.Context = nil

	pool.Put(request)
}
