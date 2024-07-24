package request

// Bag contains information about the incoming request.
type Bag struct {
	ID     string
	Header string
	Body   *Body
}
