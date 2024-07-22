package connection

type Response interface {
	Body() []byte
}
