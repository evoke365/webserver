package auth

import "time"

type Model interface {
	GetUser(id string, user *User) error
	InsertUser(user *User) (string, error)
	ActivateUser(id string) error
	VerifyUser(id, code string, user *User) error
	FindUserByTok(tok string, user *User) error
	TouchTok(id string) error
	IsErrNotFound(err error) bool
}

type Mailer interface {
	Send(to string, data interface{}) error
}

type User struct {
	Email                string    `bson:"email,omitempty,omitempty"`
	Password             string    `bson:"password,omitempty"`
	Timezone             int       `bson:"timezone,omitempty"`
	Token                string    `bson:"token,omitempty"`
	TokenExpiry          time.Time `bson:"token_expiry,omitempty"`
	ActivationCode       string    `bson:"activation_code,omitempty"`
	ActivationCodeExpiry time.Time `bson:"activation_code_expiry,omitempty"`
	IsActive             bool      `bson:"is_active,omitempty"`
	Created              time.Time `bson:"created,omitempty"`
	Modified             time.Time `bson:"modified,omitempty"`
}
