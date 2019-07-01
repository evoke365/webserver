package auth

import "time"

type Model interface {
	GetUser(id string, user *User) error
	InsertUser(user *User) (string, error)
	VerifyUser(id, code string, user *User) error
	IsErrNotFound(err error) bool
}

type User struct {
	Email                string    `json:"email,omitempty,omitempty"`
	Password             string    `json:"password,omitempty"`
	Timezone             int       `json:"timezone,omitempty"`
	Token                string    `json:"token,omitempty"`
	TokenExpiry          time.Time `json:"tokenExpiry,omitempty"`
	ActivationCode       string
	ActivationCodeExpiry time.Time `json:"activationCodeExpiry,omitempty"`
	IsActive             bool      `json:"isActive,omitempty"`
	Created              time.Time `json:"created,omitempty"`
	Modified             time.Time `json:"modified,omitempty"`
}
