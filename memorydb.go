package auth

import (
	"time"

	"github.com/google/uuid"
	"github.com/jacygao/crud"
)

const memoryDBExpirySec = 60 * 60 * 24

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