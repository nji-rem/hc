package password

type Hasher interface {
	Hash(plaintext string) (string, error)
}
