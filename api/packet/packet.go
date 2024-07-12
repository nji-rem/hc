package packet

type Packet struct {
	Name       string
	Handler    HandlerFunc
	Middleware []MiddlewareFunc
}
