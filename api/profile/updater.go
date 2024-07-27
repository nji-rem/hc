package profile

type Updatable struct {
	Motto  string
	Figure string
	Sex    string
}

type Updater interface {
	Update(accountId int, updatable Updatable) error
}
