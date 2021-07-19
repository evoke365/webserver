package fakedb

import (
	"encoding/json"

	"github.com/evoke365/webserver/store"
	"github.com/evoke365/webserver/store/data"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FakeDB struct {
	store.DB
	userDB map[string][]byte
}

func NewFakeDB() *FakeDB {
	return &FakeDB{
		userDB: make(map[string][]byte),
	}
}

func (db *FakeDB) GetUser(id string, user *data.User) error {
	return json.Unmarshal(db.userDB[id], user)
}

func (db *FakeDB) InsertUser(user *data.User) (string, error) {
	id := primitive.NewObjectID().String()
	b, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	db.userDB[id] = b
	return id, nil
}
