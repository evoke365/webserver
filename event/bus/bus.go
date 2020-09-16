package bus

type PubSub interface {
	Publish(topic string, message []byte) error
	Subscribe(topic string, command func([]byte))
}
