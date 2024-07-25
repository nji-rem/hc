package store

import "hc/internal/account/domain/accountaggregate"

type Player interface {
	Add(entity accountaggregate.Entity) error
	NameTaken(username accountaggregate.Username) (bool, error)
}
