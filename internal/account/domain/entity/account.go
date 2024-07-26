package entity

import (
	"database/sql"
)

type Account struct {
	ID        int
	Username  Username
	Password  string
	CreatedAt sql.NullString `db:"created_at"`
	UpdatedAt sql.NullString `db:"updated_at"`
}
