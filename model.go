package auth

import "time"

type Model interface {
	GetUser(id string, user *User) error
	InsertUser(user *User) (string, error)
	VerifyUser(id, code string, user *User) error
	IsErrNotFound(err error) bool
}

type Mailer interface {
	Send(to string, data interface{}) error
}

type User struct {
	Email                string    `json:"email,omitempty,omitempty"`
	Password             string    `json:"password,omitempty"`
	Timezone             int       `json:"timezone,omitempty"`
	Token                string    `json:"token,omitempty"`
	TokenExpiry          time.Time `json:"token_expiry,omitempty"`
	ActivationCode       string    `json:"activation_code,omitempty"`
	ActivationCodeExpiry time.Time `json:"activation_code_expiry,omitempty"`
	IsActive             bool      `json:"is_active,omitempty"`
	Created              time.Time `json:"created,omitempty"`
	Modified             time.Time `json:"modified_date,omitempty"`
}
