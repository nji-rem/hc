package domain

type Room struct {
	ID               int
	AccountID        int `db:"account_id"`
	Name             RoomName
	Model            RoomModel
	Description      string
	RoomAccessType   RoomAccessType `db:"room_access_type"`
	RoomOwnerVisible bool           `db:"room_owner_visible"`
}
