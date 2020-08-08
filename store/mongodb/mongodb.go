package mongodb

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jacygao/auth/store/data"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultTokenExpirySec = 60 * 60 * 24
)

type MongoConfig struct {
	// MongoDB connection string
	URI string
	// name of the DB
	DBName string
}

func (mc MongoConfig) Validate() error {
	if len(mc.URI) == 0 {
		return errors.New("missing URI")
	}
	if len(mc.DBName) == 0 {
		return errors.New("missing DBName")
	}
	return nil
}

// MongoDB defines the structure of the instance.
type MongoDB struct {
	user, note *mongo.Collection
}

// NewMongoDB returns a new instande of MongoDB implementation of the DB interface.
func NewMongoDB(c MongoConfig) (*MongoDB, error) {
	client, err := newMongoClient(c.URI)
	if err != nil {
		return nil, err
	}
	db := client.Database(c.DBName)
	return &MongoDB{
		user: db.Collection(data.UserCollectionLabel),
		note: db.Collection(data.NoteCollectionLabel),
	}, nil
}

func newMongoClient(uri string) (*mongo.Client, error) {
	mongoCli, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if err := mongoCli.Connect(ctx); err != nil {
		log.Fatalf("failed to connect to mongodb. error: %v", err)
	}
	return mongoCli, nil
}

func (m *MongoDB) IsErrNotFound(err error) bool {
	return err == mongo.ErrNoDocuments
}

func newToken() string {
	return uuid.New().String()
}
