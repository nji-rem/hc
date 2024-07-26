package store

import (
	"hc/internal/account/domain/entity"
)

type Player interface {
	FindByUsername(username string) (bool, entity.Account, error)
	Add(entity entity.Account) (int, error)
	NameTaken(username entity.Username) (bool, error)
}
