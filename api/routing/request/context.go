package request

import "context"

type Context struct {
	context.Context
}

func NewContext(context context.Context) Context {
	return Context{
		Context: context,
	}
}
