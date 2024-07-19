package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateAccountsTable, downCreateAccountsTable)
}

func upCreateAccountsTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func downCreateAccountsTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE 'accounts'")
	if err != nil {
		return err
	}

	return nil
}
