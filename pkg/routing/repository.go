package routing

import (
	"fmt"
	"hc/api/routing"
)

type Repository struct {
	Routes map[string]routing.Route
}

func (r Repository) Get(header string) (routing.Route, error) {
	route, ok := r.Routes[header]
	if !ok {
		return routing.Route{}, fmt.Errorf("no route for header %s", header)
	}

	return route, nil
}
