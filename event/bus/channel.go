package bus

import (
	"time"
)

type Config struct {
	Topics                  []string
	QueueRetryDelayMilliSec int
}

type Event struct {
	Topic   string
	Message []byte
}

type ChannelBus struct {
	conf   Config
	bus    chan Event
	queues map[string](chan []byte)
}

func NewChannelBus(c Config) *ChannelBus {
	qs := make(map[string](chan []byte))
	for _, topic := range c.Topics {
		qs[topic] = make(chan []byte)
	}

	bus := &ChannelBus{
		bus:    make(chan Event),
		queues: qs,
	}

	go bus.start()

	// Waiting for channel to set up
	time.Sleep(10 * time.Millisecond)

	return bus
}

// TODO: implementing stop signal
func (c *ChannelBus) start() {
	for {
		event := <-c.bus
		c.queues[event.Topic] <- event.Message
		time.Sleep(10 * time.Millisecond)
	}
}

func (c *ChannelBus) Publish(topic string, message []byte) error {
	c.bus <- Event{
		Topic:   string(topic),
		Message: message,
	}
	return nil
}

func (c *ChannelBus) Subscribe(topic string, command func([]byte)) {
	for {
		event := <-c.queues[topic]
		if event != nil {
			command(event)
		}
		time.Sleep(time.Duration(c.conf.QueueRetryDelayMilliSec) * time.Millisecond)
	}
}
