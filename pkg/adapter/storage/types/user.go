package types

type User struct {
	Base
	Username  string
	IP        string
	Connected bool
}
