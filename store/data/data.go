package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	UserCollectionLabel  = "user"
	NoteCollectionLabel  = "note"
	EventCollectionLabel = "event"
)

// User defines the user data structure for database operations.
type User struct {
	Email                string    `bson:"email,omitempty,omitempty"`
	Password             string    `bson:"password,omitempty"`
	Timezone             int32     `bson:"timezone,omitempty"`
	Token                string    `bson:"token,omitempty"`
	TokenExpiry          time.Time `bson:"token_expiry,omitempty"`
	ActivationCode       string    `bson:"activation_code,omitempty"`
	ActivationCodeExpiry time.Time `bson:"activation_code_expiry,omitempty"`
	IsActive             bool      `bson:"is_active,omitempty"`
	Created              time.Time `bson:"created,omitempty"`
	Modified             time.Time `bson:"modified,omitempty"`
}

// Note defines the note data structure for database operations.
type Note struct {
	ID        primitive.ObjectID `validate:"nonzero" bson:"_id,omitempty"`
	UserID    string             `validate:"nonzero" bson:"userId,omitempty"`
	Keyword   string             `validate:"nonzero" bson:"keyword,omitempty"`
	Answer    string             `validate:"nonzero" bson:"answer,omitempty"`
	Important bool               `bson:"important"`
	Created   time.Time          `bson:"created,omitempty"`
	Modified  time.Time          `bson:"modified,omitempty"`
	Deleted   bool               `bson:"deleted"`
}

type Event struct {
	ID          primitive.ObjectID
	AggregateID string
	Topic       EventType
	Data        []byte
	Timestamp   time.Time
	Version     int
}

type EventType string
