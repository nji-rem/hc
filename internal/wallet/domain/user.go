package domain

type User struct {
	ID        int
	AccountID int `db:"account_id"`
	Credits   int
}
