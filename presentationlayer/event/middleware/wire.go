package middleware

import (
	"hc/api/session"
	"sync"
)

var (
	mustBeAuthenticatedMiddlewareOnce sync.Once
	mustBeAuthenticatedMiddleware     *MustBeAuthenticated
)

func ProvideAuthentication(store session.Store) *MustBeAuthenticated {
	mustBeAuthenticatedMiddlewareOnce.Do(func() {
		mustBeAuthenticatedMiddleware = &MustBeAuthenticated{SessionStore: store}
	})

	return mustBeAuthenticatedMiddleware
}
