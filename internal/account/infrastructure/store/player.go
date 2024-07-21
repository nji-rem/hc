package store

import "github.com/jmoiron/sqlx"

type Player struct {
	DB *sqlx.DB
}
