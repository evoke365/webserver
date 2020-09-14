package bus

type Publisher interface {
	Publish(topic string, message []byte) error
}

type Subscriber interface {
	Subscribe(topic string) ([]byte, error)
}
