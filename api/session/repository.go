package session

type Repository interface {
	Add(ctx Context) error
}
