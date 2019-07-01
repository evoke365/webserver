package auth

import (
	"time"
	"errors"

	"github.com/google/uuid"
	"github.com/JacyGao/crud"
)

const memoryDBExpirySec = 60 * 60 * 24

var errNotFound = errors.New("not found")

// MemoryDB defines instance structure and dependencies.
type MemoryDB struct {
	store *crud.CRUD
}

// NewMemoryDB initialises a new MemoryDB instance.
func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		store: crud.New(),
	}
}

// GetUser gets user of given id from memory db.
func (db *MemoryDB) GetUser(id string, user *User) error {
	if _, err := db.store.Get(id, user); err != nil {
		return err
	}
	if user == nil {
		return errNotFound
	}
	return nil
}

// InsertUser inserts a new user to memory db.
func (db *MemoryDB) InsertUser(user *User) (string, error) {
	now := time.Now()
	id := uuid.New().String()
	tok := id
	user.Token = tok
	user.TokenExpiry = now.Add(time.Second * memoryDBExpirySec)
	user.Created = now
	user.Modified = now
	if _, err := db.store.Insert(id, user, memoryDBExpirySec); err != nil {
		return "", err
	}
	return tok, nil
}

func (db *MemoryDB) VerifyUser(id, code string, user *User) error {
	if _, err := db.store.Get(id, user); err != nil {
		return err
	}

	if user == nil {
		return errNotFound
	}

	if user.ActivationCode != code {
		user = nil
		return errNotFound
	}

	return nil
}

func (db *MemoryDB) IsErrNotFound(err error) bool {
	return err == errNotFound
}