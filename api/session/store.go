package session

type Store interface {
	Add(session *Bag) error
	Get(id string) (*Bag, error)
	Delete(id string) (*Bag, error)
}
