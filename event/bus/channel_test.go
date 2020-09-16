package bus

import (
	"bytes"
	"context"
	"testing"
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
	if bytes.Compare(mockMsg, msg) != 0 {
		t.Fatalf("expected %+v but got %+v", mockMsg, msg)
	}
}
