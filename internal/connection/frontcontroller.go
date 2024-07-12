package connection

import (
	"hc/api/connection"
	"hc/api/packet"
	"io"
)

type FrontController struct {
	Resolver       packet.Resolver
	WrapMiddleware packet.WrapFunc
}

func (f FrontController) Handle(request *connection.Request, writer io.Writer) error {
	// Resolve route
	route, err := f.Resolver.Get(request.Header)
	if err != nil {
		return err
	}

	// Wrap middleware
	handler := f.WrapMiddleware(route.Handler, route.Middleware)

	// Execute the handler
	return handler(request, writer)
}
