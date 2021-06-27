package queue

import (
	"fmt"
	"github.com/beuus39/order/internal/shared/dtos"
	"github.com/beuus39/order/pkg/nats"
)

type NatsProductConfig struct {
	nats nats.NatsConfig
}

func (n *NatsProductConfig) SubscriberProduct(topic string) {
	c, err := n.nats.Connect()
	defer c.Close()
	_, err = c.Subscribe(topic, func(subj, reply string, product *dtos.ProductDto) {
		fmt.Printf("Received a product on subject %s! %+v\n", subj, product)
	})
	if err != nil {
		return
	}
}

func (n *NatsProductConfig) SubscribeHello(topic string) {
	c, err := n.nats.Connect()
	defer c.Close()
	if err != nil {
		return
	}

	_, err = c.Subscribe(topic, func(s string) {
		fmt.Printf("Received a message: %s\n", s)
	})
	if err != nil {
		return
	}
}

func NewProductSubscriber(nats nats.NatsConfig) ProductSubscriber {
	return &NatsProductConfig{
		nats: nats,
	}
}
