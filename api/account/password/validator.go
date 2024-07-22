package password

type ValidationStatus int

const (
	Valid ValidationStatus = iota
	PasswordTooShort
	PasswordTooLong
)

type ValidationFunc func(password string) (ValidationStatus, error)
