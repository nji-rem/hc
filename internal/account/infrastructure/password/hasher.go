package password

import "golang.org/x/crypto/bcrypt"

type HashService struct{}

// Cost contains the bcrypt cost. Number should be high enough to be computationally expensive to crack. This is
// currently not configurable.
const Cost = 13

func (h HashService) Hash(plaintext string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plaintext), Cost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
