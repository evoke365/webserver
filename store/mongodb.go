package store

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const defaultTokenExpirySec = 60 * 60 * 24

type MongoConfig struct {
	// MongoDB connection string
	URI string
	// name of the DB
	DBName string
	// name of the collection for storing user info
	CollectionName string
}

func (mc MongoConfig) Validate() error {
	if len(mc.URI) == 0 {
		return errors.New("missing URI")
	}
	if len(mc.DBName) == 0 {
		return errors.New("missing DBName")
	}
	if len(mc.CollectionName) == 0 {
		return errors.New("missing CollectionName")
	}
	return nil
}

// MongoDB defines the structure of the instance.
type MongoDB struct {
	Collection *mongo.Collection
}

// NewMongoDB returns a new instande of MongoDB implementation of the DB interface.
func NewMongoDB(c MongoConfig) (*MongoDB, error) {
	client, err := newMongoClient(c.URI)
	if err != nil {
		return nil, err
	}
	return &MongoDB{
		Collection: client.Database(c.DBName).Collection(c.CollectionName),
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

func (m *MongoDB) GetUser(id string, user *User) error {
	return m.Collection.FindOne(context.Background(), bson.M{"email": id}).Decode(&user)
}

func (m *MongoDB) InsertUser(user *User) (string, error) {
	now := time.Now()
	tok := newToken()
	id := primitive.NewObjectID()
	if _, err := m.Collection.InsertOne(context.Background(),
		bson.M{
			"_id":                    id,
			"email":                  user.Email,
			"password":               user.Password,
			"token":                  tok,
			"token_expiry":           now.Add(time.Second * defaultTokenExpirySec),
			"timezone":               user.Timezone,
			"activation_code":        user.ActivationCode,
			"activation_code_expiry": user.ActivationCodeExpiry,
			"is_active":              user.IsActive,
			"created":                now,
			"modified":               now,
		}); err != nil {
		return "", err
	}
	return id.String(), nil
}

func (m *MongoDB) ActivateUser(id string) error {
	if _, err := m.Collection.UpdateOne(context.Background(), bson.M{"email": id}, bson.M{"$set": bson.M{"is_active": true}}); err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) VerifyUser(id, code string, user *User) error {
	return m.Collection.FindOne(
		context.Background(),
		bson.M{
			"email":           id,
			"activation_code": code,
			"activation_code_expiry": bson.M{
				"$gte": time.Now(),
			},
		},
	).Decode(&user)
}

func (m *MongoDB) UpdateActiveCode(id, code string, exp time.Time) (*User, error) {
	user := &User{}
	opts := options.FindOneAndUpdate()
	opts.SetReturnDocument(options.After)
	if err := m.Collection.FindOneAndUpdate(
		context.Background(),
		bson.M{"email": id},
		bson.M{"$set": bson.M{
			"activation_code":        code,
			"activation_code_expiry": exp,
		}},
		opts,
	).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (m *MongoDB) FindUserByTok(tok string, user *User) error {
	return m.Collection.FindOne(context.Background(), bson.M{"token": tok}).Decode(&user)
}

func (m *MongoDB) TouchTok(id string) error {
	if _, err := m.Collection.UpdateOne(
		context.Background(),
		bson.M{"email": id},
		bson.M{
			"$set": bson.M{
				"token_expiry": time.Now().Add(time.Second * defaultTokenExpirySec),
			},
		},
	); err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) UpdatePassword(id, tok, pwd string) error {
	if _, err := m.Collection.UpdateOne(
		context.Background(),
		bson.M{"email": id, "token": tok},
		bson.M{"$set": bson.M{"password": pwd}},
	); err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) IsErrNotFound(err error) bool {
	return err == mongo.ErrNoDocuments
}

func newToken() string {
	return uuid.New().String()
}
