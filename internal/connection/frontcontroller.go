package connection

import (
	"github.com/rs/zerolog/log"
	"hc/api/connection"
	"hc/api/connection/request"
	"hc/api/packet"
	"io"
)

type FrontController struct {
	Resolver       packet.Resolver
	WrapMiddleware packet.WrapFunc
}

func (f FrontController) Handle(request *request.Bag, writer io.Writer) error {
	// Resolve route
	route, err := f.Resolver.Get(request.Header)
	if err != nil {
		return err
	}

	// Wrap middleware
	handler := f.WrapMiddleware(route.Handler, route.Middleware)

	// Create a goroutine that writes to the writer
	ch := make(chan connection.Response)
	defer close(ch)

	go func(writer io.Writer) {
	MainLoop:
		for {
			select {
			case data, ok := <-ch:
				if !ok {
					log.Debug().Msg("Response channel closed for incoming request")
					break MainLoop
				}

				if _, err := writer.Write(data.Body()); err != nil {
					log.Error().Msgf("unable to send message response: %s", err.Error())
					break MainLoop
				}

				log.Info().Msgf("[S->C] Sent %x", data)
			}
		}
	}(writer)

	// Execute the handler
	return handler(request, ch)
}
