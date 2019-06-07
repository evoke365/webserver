package auth

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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
