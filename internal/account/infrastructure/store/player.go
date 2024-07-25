package store

import (
	"database/sql"
	"errors"
	"fmt"
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

func (p *Player) FindByUsername(username string) (bool, accountaggregate.Entity, error) {
	var entity accountaggregate.Entity
	err := p.DB.Get(&entity, "SELECT * FROM accounts WHERE username = ?", username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, accountaggregate.Entity{}, nil
		}

		return false, accountaggregate.Entity{}, err
	}

	return true, entity, nil
}

func (p *Player) Add(entity accountaggregate.Entity) error {
	query := "INSERT INTO accounts (username, password, look, gender, motto) VALUES (:username, :password, :look, :gender, :motto)"
	result, err := p.DB.NamedExec(query, entity)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("expected 1 row to be affected, but %d rows are affected", rowsAffected)
	}

	return nil
}
