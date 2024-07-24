package packet

import (
	"hc/api/connection"
	"hc/api/connection/request"
)

type (
	// HandlerFunc contains the function signature that is required for packet handlers.
	//
	// Argument request contains a reference to the current request. It contains data such as the packet header, the
	// packet body, and some other information.
	HandlerFunc func(sessionId string, request *request.Bag, response chan<- connection.Response) error

	// MiddlewareFunc contains the function signature that is required for middleware. Each middleware is wrapped around
	// a packet handler, HandlerFunc.
	MiddlewareFunc func(next HandlerFunc) HandlerFunc

	// WrapFunc is the function signature responsible for wrapping the HandlerFunc and a slice of MiddlewareFunc together
	// into one HandlerFunc, making it runnable.
	WrapFunc func(handler HandlerFunc, middleware []MiddlewareFunc) HandlerFunc
)
