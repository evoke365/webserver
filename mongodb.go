package auth

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const defaultTokenExpirySec = 60 * 60 * 24

type MongoDB struct {
	session    *mgo.Session
	name       string
	collection string
}

func NewMongoDB(s *mgo.Session, name, collection string) (*MongoDB, error) {
	return &MongoDB{
		session:    s,
		name:       name,
		collection: collection,
	}, nil
}

// Ensure closure of mongo db session.
func (db *MongoDB) withCollection(f func(c *mgo.Collection) error) error {
	s := db.session.Copy()
	defer s.Close()
	c := s.DB(db.name).C(db.collection)
	return f(c)
}

func (db *MongoDB) GetUser(email string, user *User) error {
	if err := db.withCollection(func(c *mgo.Collection) error {
		if err := c.Find(bson.M{"email": email}).One(&user); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (db *MongoDB) InsertUser(user *User) (string, error) {
	now := time.Now()
	tok := newToken()
	if err := db.withCollection(func(c *mgo.Collection) error {
		if err := c.Insert(
			bson.M{
				"_id":         bson.NewObjectId(),
				"email":       user.Email,
				"password":    user.Password,
				"token":       tok,
				"tokenExpiry": now.Add(time.Second * defaultTokenExpirySec),
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
