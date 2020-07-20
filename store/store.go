package store

import "time"

// DB defines a common interface of available database operations.
type DB interface {
	GetUser(id string, user *User) error
	InsertUser(user *User) (string, error)
	ActivateUser(id string) error
	VerifyUser(id, code string, user *User) error
	UpdateActiveCode(id, code string, exp time.Time) (*User, error)
	FindUserByTok(tok string, user *User) error
	TouchTok(id string) error
	UpdatePassword(id, tok, pwd string) error
	IsErrNotFound(err error) bool
}
