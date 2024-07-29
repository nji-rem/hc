package request

// Bag contains information about the incoming viewmodel.
type Bag struct {
	ID     string
	Header string
	Body   *Body
}
