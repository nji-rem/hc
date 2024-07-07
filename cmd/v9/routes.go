package main

import (
	"fmt"
	"hc/api/routing"
	"hc/api/routing/request"
)

func CollectRoutes() func() []routing.Route {
	return func() []routing.Route {
		return []routing.Route{
			// Unauthenticated routes
			{
				Name: "CU",
				Handler: func(ctx request.Context, packet any) error {
					fmt.Println("handle secret key request")
					return nil
				},
				Middleware: []routing.MiddlewareFunc{},
			},
		}
	}
}
