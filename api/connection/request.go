package connection

import (
	"context"
)

type Request struct {
	// Header contains the packet id.
	Header string

	// Body contains the byte buffer.
	Body []byte

	context.Context
}
