package room

type CreateRoom interface {
	Create(accountId int, name, model, accessType string, roomOwnerVisible bool) error
}
