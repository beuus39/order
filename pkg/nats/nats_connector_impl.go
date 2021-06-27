package nats

import (
	"github.com/nats-io/nats.go"
)

type NatsConfig struct {
	Url string
}

func (n *NatsConfig) Connect() (*nats.EncodedConn, error) {
	nc, _ := nats.Connect(n.Url)
	c, err  := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	return c, err
}

func NewNatsConfig(url string) NatsConnector {
	return &NatsConfig{
		Url: url,
	}
}