package store

import (
	"github.com/jmoiron/sqlx"
	"hc/internal/account/domain/accountaggregate"
)

type Player struct {
	DB *sqlx.DB
}

func (p *Player) NameTaken(username accountaggregate.Username) (bool, error) {
	var count int
	if err := p.DB.Get(&count, "SELECT COUNT(*) FROM accounts WHERE username = ?", username); err != nil {
		return false, err
	}

	return count > 0, nil
}
