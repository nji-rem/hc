package routing

import (
	"context"
	"fmt"
	"github.com/panjf2000/gnet/v2"
	"github.com/rs/zerolog/log"
	routingContract "hc/api/routing"
	"hc/api/routing/request"
)

type RouteExecutor struct {
	Repository routingContract.Repository
}

func (r *RouteExecutor) ExecutePacket(header string, c gnet.Conn, data []byte) error {
	route, err := r.Repository.Get(header)
	if err != nil {
		return fmt.Errorf("unable to execute route: %s", err.Error())
	}

	log.Debug().Msgf("About to execute route %s for header %s", route.Name, header)

	ctx := request.NewContext(context.Background())

	handler := route.Handler
	for _, middleware := range route.Middleware {
		handler = middleware(handler)
	}

	return handler(ctx, c, data)
}
