package account

type VerifyCredentials interface {
	Verify(username, password string) (bool, int, error)
}
