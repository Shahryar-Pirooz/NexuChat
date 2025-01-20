package types

type User struct {
	Base
	ID        string
	Username  string
	IP        string
	Connected bool
}
