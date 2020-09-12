package event

import (
	"context"
	"encoding/json"

	"github.com/evoke365/webserver/store"
	"github.com/evoke365/webserver/store/data"
)

type Controller struct {
	store store.DB
}

func NewController(db store.DB) *Controller {
	return &Controller{
		store: db,
	}
}

// Save stores an event to the event store.
func (c *Controller) Save(ctx context.Context, aggregateID string, t data.EventType, payload interface{}) error {
	blob, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return c.store.InsertEvent(ctx, &data.Event{
		AggregateID: aggregateID,
		Topic:       t,
		Data:        blob,
	})
}

// Publish publishes an event to SNS.
func (c *Controller) Publish(ctx context.Context) {

}
