package store

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"hc/internal/account/domain/entity"
)

type Player struct {
	DB *sqlx.DB
}

func (p *Player) NameTaken(username entity.Username) (bool, error) {
	var count int
	if err := p.DB.Get(&count, "SELECT COUNT(*) FROM accountsvc_accounts WHERE username = ?", username); err != nil {
		return false, err
	}

	return count > 0, nil
}

func (p *Player) FindByUsername(username string) (bool, entity.Account, error) {
	var e entity.Account
	err := p.DB.Get(&e, "SELECT * FROM accountsvc_accounts WHERE username = ?", username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, entity.Account{}, nil
		}

		return false, entity.Account{}, err
	}

	return true, e, nil
}

func (p *Player) Add(entity entity.Account) (int, error) {
	query := "INSERT INTO accountsvc_accounts (username, password, created_at, updated_at) VALUES (:username, :password, now(), now())"
	result, err := p.DB.NamedExec(query, entity)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected != 1 {
		return 0, fmt.Errorf("expected 1 row to be affected, but %d rows are affected", rowsAffected)
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastInsertedId), nil
}
