package domain

type Store interface {
	Add(entity Profile) error
	FindByAccountID(id int) (Profile, error)
}
