package request

import "context"

type Bag struct {
	context.Context

	ID        string
	SessionID string
	Header    string
	Body      []byte
}
