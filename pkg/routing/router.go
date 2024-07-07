package routing

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"hc/api/routing"
	"hc/api/routing/request"
)

type RouteExecutor struct {
	routes map[string]routing.Route
}

func (r RouteExecutor) ExecutePacket(header string, data []byte) error {
	// Retrieve route from map (use repository for different drivers in the future).
	route, ok := r.routes[header]
	if !ok {
		return fmt.Errorf("packet with header id %s is unregistered", header)
	}

	log.Debug().Msgf("About to execute route %s for header %s", route.Name, header)

	ctx := request.NewContext(context.Background())

	handler := route.Handler
	for _, middleware := range route.Middleware {
		handler = middleware(handler)
	}

	return handler(ctx, data)
}

func NewRouteExecutor(routeCollector func() ([]routing.Route, error)) (RouteExecutor, error) {
	actualRoutes, err := routeCollector()
	if err != nil {
		return RouteExecutor{}, fmt.Errorf("unable to create route executor instance, err: %s", err.Error())
	}

	var routes map[string]routing.Route
	for _, actualRoute := range actualRoutes {
		routes[actualRoute.Name] = actualRoute
	}

	log.Info().Msgf("Registered %d packet handlers", len(routes))

	return RouteExecutor{routes: routes}, nil
}
