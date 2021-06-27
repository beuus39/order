package queue

type ProductSubscriber interface {
	SubscribeHello(topic string)
	SubscriberProduct(topic string)
}