package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/evoke365/webserver/store/data"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertEvent implements query logic that inserts a event to the event store.
func (m *MongoDB) InsertEvent(ctx context.Context, event *data.Event) error {
	if event == nil {
		return fmt.Errorf("event data cannot be empty")
	}

	event.ID = primitive.NewObjectID()
	event.Timestamp = time.Now()
	event.Version = 1

	if _, err := m.event.InsertOne(ctx, event); err != nil {
		return err
	}
	return nil
}
