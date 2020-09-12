package event

import (
	"github.com/evoke365/webserver/store/data"
)

const (
	NoteAdded           data.EventType = "noteAdded"
	NoteDeleted         data.EventType = "noteDeleted"
	NoteReminderUpdated data.EventType = "noteReminderUpdated"
)
