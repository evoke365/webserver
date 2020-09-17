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

// ChannelBus is an implementation of the Bus interface backed by go channels.
type ChannelBus struct {
	conf   Config
	bus    chan event.Message
	queues map[string](chan event.Message)
}

// NewChannelBus returns an instance of the ChannelBus struct.
func NewChannelBus(c Config) *ChannelBus {
	qs := make(map[string](chan event.Message))
	for _, topic := range c.Topics {
		qs[topic] = make(chan event.Message)
	}

	bus := &ChannelBus{
		bus:    make(chan event.Message),
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
		c.queues[event.Topic] <- event
		time.Sleep(10 * time.Millisecond)
	}
}

func (c *ChannelBus) Publish(ctx context.Context, topic string, message []byte) error {
	c.bus <- event.Message{
		Topic: string(topic),
		Data:  message,
	}
	return nil
}

func (c *ChannelBus) Subscribe(ctx context.Context, topic string, com event.Commander) {
	for {
		event := <-c.queues[topic]
		com.Execute(ctx, event)
		time.Sleep(time.Duration(c.conf.QueueRetryDelayMilliSec) * time.Millisecond)
	}
}
