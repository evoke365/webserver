package bus

import (
	"bytes"
	"context"
	"testing"

	"github.com/evoke365/webserver/event"
)

func TestPublish(t *testing.T) {
	mockConf := Config{
		Topics:                  []string{"mock1", "mock2"},
		QueueRetryDelayMilliSec: 100,
	}
	bus := NewChannelBus(mockConf)
	mockMsg := []byte("123")
	if err := bus.Publish(context.Background(), "mock1", mockMsg); err != nil {
		t.Fatal(err)
	}

	msg := <-bus.queues["mock1"]
	if bytes.Compare(mockMsg, msg.Data) != 0 {
		t.Fatalf("expected %+v but got %+v", mockMsg, msg)
	}
}

func TestSubscribe(t *testing.T) {
	mockConf := Config{
		Topics:                  []string{"mock1", "mock2"},
		QueueRetryDelayMilliSec: 100,
	}
	bus := NewChannelBus(mockConf)
	signal := make(chan struct{}, 1)
	ctx := context.Background()
	go bus.Subscribe(ctx, "mock1", &TestCommander{signal})

	if err := bus.Publish(ctx, "mock1", []byte("123")); err != nil {
		t.Fatal(err)
	}
	<-signal
}

type TestCommander struct {
	signal chan struct{}
}

func (c *TestCommander) Execute(ctx context.Context, data event.Message) {
	c.signal <- struct{}{}
}
