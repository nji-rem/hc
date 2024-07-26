package wallet

type GiftCredits interface {
	Gift(accountID int, amount int, reason string) error
}
