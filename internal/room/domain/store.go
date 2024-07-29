package domain

type Store interface {
	Add(room Room) error
}
