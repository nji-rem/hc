package room

import "errors"

var (
	ErrInvalidRoomName   = errors.New("provided room name contains illegal characters")
	ErrInvalidRoomModel  = errors.New("invalid room model")
	ErrInvalidAccessType = errors.New("invalid access type")
)
