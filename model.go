package auth

import "time"

type Model interface {
	GetUser(id string, user *User) error
	GetUserByCredentials(email, password string, user *User) error
	InsertUser(user *User) error
	PutUser(id string, user *User) error
}

type User struct {
	Email    string
	Password string
	Token    string
	Location string
	Timezone int
	Ip       string
	Created  time.Time
	Modified time.Time
}
