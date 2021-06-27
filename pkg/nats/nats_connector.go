package nats

import "github.com/nats-io/nats.go"

type NatsConnector interface {
	Connect() (*nats.EncodedConn, error)
}
