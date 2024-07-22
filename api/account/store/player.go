package store

import "hc/internal/account/domain/accountaggregate"

type Player interface {
	NameTaken(username accountaggregate.Username) (bool, error)
}
