package password

type Hasher interface {
	Hash(plaintext string) (string, error)
}

type Verifier interface {
	Verify(plaintext, ciphertext []byte) (bool, error)
}
