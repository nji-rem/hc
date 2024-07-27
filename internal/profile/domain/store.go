package domain

type Store interface {
	Update(entity Profile) error
	Add(entity Profile) error
	FindByAccountID(id int) (Profile, error)
}
