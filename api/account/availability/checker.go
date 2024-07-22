package availability

type Status int

const (
	Available Status = iota
	UsernameTooLong
	UsernameTooShort
	UsernameContainsIllegalCharacters
	UsernameTaken
)

// UsernameAvailableFunc contains the handler that's responsible for checking if a username is currently available
// to use.
type UsernameAvailableFunc func(name string) (Status, error)
