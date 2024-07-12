package packet

import (
	"hc/api/connection"
	"io"
)

type (
	// HandlerFunc contains the function signature that is required for packet handlers.
	//
	// Argument request contains a reference to the current request. It contains data such as the packet header, the
	// packet body, and some other information.
	//
	// Argument writer contains a reference to a writer, most likely with a network stream. Using an interface makes it
	// easier to mock.
	HandlerFunc func(request *connection.Request, writer io.Writer) error

	// MiddlewareFunc contains the function signature that is required for middleware. Each middleware is wrapped around
	// a packet handler, HandlerFunc.
	MiddlewareFunc func(next HandlerFunc) HandlerFunc

	// WrapFunc is the function signature responsible for wrapping the HandlerFunc and a slice of MiddlewareFunc together
	// into one HandlerFunc, making it runnable.
	WrapFunc func(handler HandlerFunc, middleware []MiddlewareFunc) HandlerFunc
)
