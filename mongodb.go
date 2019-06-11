package auth

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const defaultExpirySec = 60 * 60 * 24

type MongoDb struct {
	session    *mgo.Session
	name       string
	collection string
}

func NewMongoDb(s *mgo.Session, name, collection string) (*MongoDb, error) {
	return &MongoDb{
		session:    s,
		name:       name,
		collection: collection,
	}, nil
}

// Ensure closure of mongo db session.
func (db *MongoDb) withCollection(f func(c *mgo.Collection) error) error {
	s := db.session.Copy()
	defer s.Close()
	c := s.DB(db.name).C(db.collection)
	return f(c)
}

func (db *MongoDb) Get(key string, valuePtr interface{}) (uint64, error) {
	db.withCollection(func(c *mgo.Collection) error {
		if err := c.Find(bson.M{"_id": bson.ObjectIdHex(key)}).One(valuePtr); err != nil {
			return err
		}
		return nil
	})
	return 0, nil
}

func (db *MongoDb) Upsert(key string, value interface{}, expiry uint32) (uint64, error) {
	db.withCollection(func(c *mgo.Collection) error {
		if _, err := c.UpsertId(bson.M{"_id": bson.ObjectIdHex(key)}, value); err != nil {
			return err
		}
		return nil
	})
	return 0, nil
}

func (db *MongoDb) InsertUser(user *User) (string, error) {
	now := time.Now()
	tok := newToken()
	if err := db.withCollection(func(c *mgo.Collection) error {
		if err := c.Insert(
			bson.M{
				"_id":         bson.NewObjectId(),
				"email":       user.Email,
				"password":    user.Password,
				"token":       tok,
				"tokenExpiry": now.Add(time.Second * defaultExpirySec),
				"timezone":    user.Timezone,
				"created":     now,
				"modified":    now,
			}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return "", err
	}
	return tok, nil
}

func newToken() string {
	return uuid.New().String()
}
