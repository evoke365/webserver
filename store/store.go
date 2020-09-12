package store

import (
	"context"
	"time"

	"github.com/evoke365/webserver/store/data"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DB defines a common interface of available database operations.
type DB interface {
	GetUser(id string, user *data.User) error
	InsertUser(user *data.User) (string, error)
	ActivateUser(id string) error
	VerifyUser(id, code string, user *data.User) error
	UpdateActiveCode(id, code string, exp time.Time) (*data.User, error)
	FindUserByTok(tok string, user *data.User) error
	TouchTok(id string) error
	UpdatePassword(id, tok, pwd string) error
	IsErrNotFound(err error) bool

	GetNoteById(ctx context.Context, objID primitive.ObjectID) (*data.Note, error)
	GetNotes(ctx context.Context, userID string) ([]*data.Note, error)
	InsertNote(ctx context.Context, notes ...interface{}) error
	DeleteNote(ctx context.Context, userID, noteID string) error
	UpdateNote(ctx context.Context, userID, noteID string, newData map[string]interface{}) (*data.Note, error)
	GetUpdatedNotes(ctx context.Context, fromDate time.Time, toDate time.Time) ([]*data.Note, error)

	InsertEvent(ctx context.Context, event *data.Event) error
}
