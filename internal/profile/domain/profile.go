package domain

type Profile struct {
	ID        int
	AccountID int `db:"account_id"`
	Look      string
	Motto     string
	Sex       string
}
