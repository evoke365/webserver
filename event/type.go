package event

import (
	"github.com/evoke365/webserver/store/data"
)

const (
	// Note defines a topic string that is picked up by event consumers.
	Note data.AggregateType = "note"

	NoteAdded           data.EventType = "noteAdded"
	NoteDeleted         data.EventType = "noteDeleted"
	NoteReminderUpdated data.EventType = "noteReminderUpdated"
)
