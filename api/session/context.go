package session

import (
	"context"
	"github.com/google/uuid"
)

// Context contains the session context.
//
// The context in package session differs quite heavily from the context struct provided with each request.
// Each connected user has its own unique UUID, and each user has its own unique request id. They contain different
// things. Middlewares use it for different things, and so on. Don't get them mixed up, or you'll end up in trouble.
type Context struct {
	context.Context

	// UUID contains the session's unique identifier. We can identify a user by its uuid.
	UUID string

	// networkCh contains a channel to the writer.
	networkCh <-chan []byte
}

func NewContext(networkCh <-chan []byte, ctx context.Context) *Context {
	return &Context{
		Context:   ctx,
		UUID:      uuid.NewString(),
		networkCh: networkCh,
	}
}
