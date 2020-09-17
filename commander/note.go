package commander

import (
	"context"
	"log"

	"github.com/evoke365/webserver/event"
	"github.com/evoke365/webserver/store"
	"github.com/evoke365/webserver/store/data"
)

type NoteCommander struct {
	db store.DB
}

func NewNoteCommander(db store.DB) *NoteCommander {
	return &NoteCommander{
		db: db,
	}
}

// Execute defines logic of the topic listener.
func (c *NoteCommander) Execute(ctx context.Context, msg event.Message) {
	switch eventType := msg.Topic; data.EventType(eventType) {
	case event.NoteAdded:
		if err := c.db.InsertNote(ctx, msg.Data); err != nil {
			log.Println(err.Error())
		}
		break
	}
}
