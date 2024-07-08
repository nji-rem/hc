package routing

type Repository interface {
	Get(header string) (Route, error)
}
