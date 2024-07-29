package store

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"hc/internal/room/domain"
)

type Room struct {
	DB *sqlx.DB
}

func (r Room) Add(room domain.Room) error {
	query := `
		INSERT INTO roomsvc_rooms (name, model, description, room_access_type, room_owner_visible) 
		VALUES (:name, :model, :description, :room_access_type, :room_owner_visible)
	`

	result, err := r.DB.NamedExec(query, &room)
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
