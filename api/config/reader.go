package config

type Reader interface {
	Get(key string) any
	GetString(key string) string
}
