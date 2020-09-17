package event

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/evoke365/webserver/store"
	"github.com/evoke365/webserver/store/data"
)

type Bus interface {
	Publish(ctx context.Context, topic string, message []byte) error
	Subscribe(ctx context.Context, topic string, command Commander)
}

type Commander interface {
	Execute(ctx context.Context, data Message)
}

type Message struct {
	Topic string
	Data  []byte
}

type Controller struct {
	store store.DB
	bus   Bus
}

func NewController(db store.DB, b Bus) *Controller {
	return &Controller{
		store: db,
		bus:   b,
	}
}

// Save stores an event to the event store.
func (c *Controller) Save(ctx context.Context, aggregateID string, at data.AggregateType, et data.EventType, payload interface{}) error {
	blob, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	if err := c.store.InsertEvent(ctx, &data.Event{
		AggregateID:   aggregateID,
		AggregateType: at,
		Type:          et,
		Data:          blob,
	}); err != nil {
		return err
	}

	return c.bus.Publish(ctx, string(et), blob)
}

func (c *Controller) StartSubscribers(ctx context.Context, commands map[string]Commander) error {
	if len(commands) == 0 {
		return fmt.Errorf("no topics found to subscribe")
	}

	for topic, command := range commands {
		go c.bus.Subscribe(ctx, topic, command)
	}
	return nil
}
