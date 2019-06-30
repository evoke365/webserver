package auth

import "time"

type Model interface {
	GetUser(id string, user *User) error
	InsertUser(user *User) (string, error)
	VerifyUser(id, code string, user *User) error
	IsErrNotFound(err error) bool
}

type User struct {
	Email                string
	Password             string
	Timezone             int
	Token                string
	TokenExpiry          time.Time
	ActivationCode       string
	ActivationCodeExpiry time.Time
	isActive             bool
	Created              time.Time
	Modified             time.Time
}
