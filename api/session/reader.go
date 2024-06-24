package session

type Reader interface {
	FindById(id string)
}
