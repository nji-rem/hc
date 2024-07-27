package domain

import (
	"errors"
	"fmt"
	"strings"
)

type RoomName string

var ErrRoomNameEmpty = errors.New("room name is empty")

var allowedCharacters = map[uint8]bool{
	' ': true,
	'@': true,
	'!': true,
	'#': true,
	'(': true,
	')': true,
	'/': true,
	'*': true,
	'[': true,
	']': true,
}

func NewRoomName(value string) (RoomName, error) {
	trimmedValue := strings.TrimSpace(value)
	if len(trimmedValue) < 1 {
		return "", ErrRoomNameEmpty
	}

	for i := 0; i < len(value); i++ {
		v := value[i]
		switch {
		case v >= 'a' && v <= 'z',
			v >= 'A' && v <= 'Z',
			v >= '0' && v <= '9':
			continue
		}

		if _, isAllowedCharacter := allowedCharacters[v]; !isAllowedCharacter {
			return "", fmt.Errorf("character '%c' is considered an illegal character", v)
		}
	}

	return RoomName(value), nil
}
