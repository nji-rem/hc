package connection

import (
	"hc/api/connection/request"
	"io"
)

type TrafficParser interface {
	Parse(reader io.Reader, request *request.Bag) error
}
