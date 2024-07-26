package profile

type CreateProfile interface {
	Create(accountID int, motto, figure, sex string) error
}
