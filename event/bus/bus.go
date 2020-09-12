package bus

import "github.com/evoke365/webserver/store/data"

type Publisher interface {
	Publish(topic data.EventType, message []byte) error
}

type Subscriber interface {
	Subscribe(topic string) ([]byte, error)
}
