package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateAccountsTable, downCreateAccountsTable)
}

const Sql = `
	CREATE TABLE accounts (
	    id INTEGER PRIMARY KEY,
	    username VARCHAR(255) NOT NULL,
	    password VARCHAR(255) NOT NULL,
	    look VARCHAR(255) NOT NULL,
	    motto VARCHAR(255)
	);
`

func upCreateAccountsTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(Sql)
	if err != nil {
		return err
	}

	return nil
}

func downCreateAccountsTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE 'accounts'")
	if err != nil {
		return err
	}

	return nil
}
