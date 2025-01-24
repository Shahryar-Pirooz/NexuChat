package natsserver

import "github.com/nats-io/nats.go"

type natsOpt struct {
	ip string
}

func NewNatsService(ip string) *natsOpt {
	return &natsOpt{
		ip: ip,
	}
}

func (no *natsOpt) Connect() *nats.Conn {
	nc, err := nats.Connect(no.ip)
	if err != nil {
		panic(err)
	}
	return nc
}
