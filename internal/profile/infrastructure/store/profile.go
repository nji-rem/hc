package store

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"hc/internal/profile/domain"
)

type Profile struct {
	DB *sqlx.DB
}

func (p *Profile) Add(entity domain.Profile) error {
	query := "INSERT INTO profilesvc_profiles (account_id, look, motto, sex) VALUES (:account_id, :look, :motto, :sex)"
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
