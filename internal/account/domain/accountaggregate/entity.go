package accountaggregate

import "database/sql"

type Entity struct {
	ID        int
	Username  Username
	Password  string
	Gender    string
	Look      string
	Motto     string
	CreatedAt sql.NullString `db:"created_at"`
	UpdatedAt sql.NullString `db:"updated_at"`
}
