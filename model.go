package auth

import "time"

type Model interface {
	GetUser(id string, user *User) error
	InsertUser(user *User) (string, error)
	VerifyUser(id, code string, user *User) error
	IsErrNotFound(err error) bool
}

type User struct {
	Email                string    `json:"email"`
	Password             string    `json:"password"`
	Timezone             int       `json:"timezone"`
	Token                string    `json:"token"`
	TokenExpiry          time.Time `json:"tokenExpiry"`
	ActivationCode       string    `json:"activationCode"`
	ActivationCodeExpiry time.Time `json:"activationCodeExpiry"`
	IsActive             bool      `json:"isActive"`
	Created              time.Time `json:"created"`
	Modified             time.Time `json:"modified"`
}
