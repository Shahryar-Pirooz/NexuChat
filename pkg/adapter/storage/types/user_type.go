package types

import "time"

type User struct {
	ID        string
	Username  string
	Password  string
	IP        string
	Role      uint8
	Createdat time.Time
	Updatedat time.Time
}
