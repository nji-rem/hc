package routing

type Route struct {
	Name       string
	Handler    HandlerFunc
	Middleware map[string]MiddlewareFunc
}
