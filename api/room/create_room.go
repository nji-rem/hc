package room

type CreateRoom interface {
	Create(accountId int, name, model, description, accessType string, roomOwnerVisible bool) error
}
