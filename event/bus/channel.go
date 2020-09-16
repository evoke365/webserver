package bus

import (
	"context"
	"time"

	"github.com/evoke365/webserver/event"
)

// Config defines configuration fields for running the event bus.
type Config struct {
	Topics                  []string
	QueueRetryDelayMilliSec int
}

type Event struct {
	Topic   string
	Message []byte
}

// ChannelBus is an implementation of the Bus interface backed by go channels.
type ChannelBus struct {
	conf   Config
	bus    chan Event
	queues map[string](chan []byte)
}

// NewChannelBus returns an instance of the ChannelBus struct.
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

func (c *ChannelBus) Publish(ctx context.Context, topic string, message []byte) error {
	c.bus <- Event{
		Topic:   string(topic),
		Message: message,
	}
	return nil
}

func (c *ChannelBus) Subscribe(ctx context.Context, topic string, command event.Command) {
	for {
		event := <-c.queues[topic]
		if event != nil {
			command(ctx, event)
		}
		time.Sleep(time.Duration(c.conf.QueueRetryDelayMilliSec) * time.Millisecond)
	}
}
