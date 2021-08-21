package fakedb

import (
	"encoding/json"
	"errors"
	"time"
	"fmt"

	"github.com/evoke365/webserver/store"
	"github.com/evoke365/webserver/store/data"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	defaultTokenExpirySec = 5
)

var ErrNoDocument = errors.New("no document")

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
	val, ok := db.userDB[id]
	if !ok {
		return ErrNoDocument
	}
	return json.Unmarshal(val, user)
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

func (db *FakeDB) UpSertUser(id string, user *data.User) error {
	b, err := json.Marshal(user)
	if err != nil {
		return err
	}
	db.userDB[id] = b
	return nil
}

func (db *FakeDB) ActivateUser(id string) error {
	user := &data.User{}
	if err := db.GetUser(id, user); err != nil {
		return err
	}

	user.IsActive = true
	if err := db.UpSertUser(id, user); err != nil {
		return err
	}

	return nil
}

func (db *FakeDB) UpdateActiveCode(id, code string, exp time.Time) (*data.User, error) {
	user := &data.User{}
	if err := db.GetUser(id, user); err != nil {
		return nil, err
	}
	user.ActivationCode = code
	user.ActivationCodeExpiry = exp
	return user, db.UpSertUser(id, user)
}

// FindUserByTok performs a db scan. It is not ideal for prod usage but ok for unit testing here.
func (db *FakeDB) FindUserByTok(tok string, user *data.User) error {
	for _, val := range db.userDB {
		if err := json.Unmarshal(val, user); err != nil {
			return err
		}
		if user.Token == tok {
			return nil
		}
	}
	return ErrNoDocument
}

func (db *FakeDB) TouchTok(id string) error {
	user := &data.User{}
	if err := db.GetUser(id, user); err != nil {
		return err
	}
	user.TokenExpiry = time.Now().Add(time.Second * defaultTokenExpirySec)

	if err := db.UpSertUser(id, user); err != nil {
		return err
	}
	return nil
} 

func(db *FakeDB) UpdatePassword(id, tok, pwd string) error {
	user := &data.User{}
	if err := db.GetUser(id, user); err != nil {
		return err
	}

	if user.Token != tok {
		return ErrNoDocument
	}

	user.Password = pwd
	return nil
}

func (m *FakeDB) IsErrNotFound(err error) bool {
	return err == ErrNoDocument
}

func (m *FakeDB) isTokenValid(id string) bool {
	user := &data.User{}
	if err := m.GetUser(id, user); err != nil {
		return false
	}
	fmt.Printf("%+v", user)
	return user.TokenExpiry.After(time.Now())
}