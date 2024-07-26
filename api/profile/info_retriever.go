package profile

type Info struct {
	Motto  string
	Figure string
	Sex    string
}

type InfoRetriever interface {
	Retrieve(accountId int) (Info, error)
}
