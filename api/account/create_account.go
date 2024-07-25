package account

type CreateAccount interface {
	Create(name, password, figure, gender string) (bool, error)
}
