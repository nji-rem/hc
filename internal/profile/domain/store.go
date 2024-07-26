package domain

type Store interface {
	Add(entity Profile) error
}
