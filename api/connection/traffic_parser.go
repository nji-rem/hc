package connection

import "io"

type TrafficParser interface {
	Parse(reader io.Reader, request *Request) error
}
