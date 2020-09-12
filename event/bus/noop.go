package bus

import "github.com/evoke365/webserver/store/data"

type NoopPublisher struct{}

func (n *NoopPublisher) Publish(topic data.EventType, message []byte) error { return nil }

type NoopSubscriber struct{}

func (n *NoopSubscriber) Subscribe(topic data.EventType) []byte { return nil, nil }
