package application

import apiStore "hc/api/account/store"

type CheckNameAvailabilityHandler struct {
	Store apiStore.Player
}

func (c CheckNameAvailabilityHandler) Handle(name string) (available bool, err error) {
	// WIP
	return
}
