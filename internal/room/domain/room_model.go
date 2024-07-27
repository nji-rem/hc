package domain

import (
	"fmt"
	"strings"
)

type RoomModel string

func NewRoomModel(value string) (RoomModel, error) {
	if !strings.HasPrefix(value, "model_") {
		return "", fmt.Errorf("room model does not contain required prefix 'model_'")
	}

	model := strings.Replace(value, "model_", "", 1)
	if len(model) > 1 {
		return "", fmt.Errorf("invalid room model %s", model)
	}

	if model[0] < 'a' || model[0] > 'n' {
		return "", fmt.Errorf("room model %s is invalid", model)
	}

	return RoomModel(value), nil
}
