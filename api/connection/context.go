package connection

import "context"

type Context struct {
	context.Context

	sessionID string
}

func (c Context) SessionID() string {
	return c.sessionID
}

func NewContext(sessionID string) Context {
	return Context{
		Context:   context.Background(),
		sessionID: sessionID,
	}
}
