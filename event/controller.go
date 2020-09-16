package event

import (
	"context"
	"encoding/json"

	"github.com/evoke365/webserver/event/bus"
	"github.com/evoke365/webserver/store"
	"github.com/evoke365/webserver/store/data"
)

type Controller struct {
	store     store.DB
	publisher bus.PubSub
}

func NewController(db store.DB, pub bus.Publisher) *Controller {
	return &Controller{
		store:     db,
		publisher: pub,
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

	return c.publisher.Publish(string(et), blob)
}
