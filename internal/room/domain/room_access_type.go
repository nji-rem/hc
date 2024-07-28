package domain

type RoomAccessType int

const (
	Open RoomAccessType = iota
	Closed
	Password
)

func NewRoomAccessType(str string) RoomAccessType {
	if str == "closed" {
		return Closed
	}

	if str == "password" {
		return Password
	}

	return Open
}
