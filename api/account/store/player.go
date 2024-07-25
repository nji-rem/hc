package store

import "hc/internal/account/domain/accountaggregate"

type Player interface {
	FindByUsername(username string) (bool, accountaggregate.Entity, error)
	Add(entity accountaggregate.Entity) error
	NameTaken(username accountaggregate.Username) (bool, error)
}
