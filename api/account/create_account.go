package account

type CreateAccount interface {
	Create(name, password string) (int, error)
}
