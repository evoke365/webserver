package store

import "time"

// User defines the user data structure for database operations.
type User struct {
	Email                string    `bson:"email,omitempty,omitempty"`
	Password             string    `bson:"password,omitempty"`
	Timezone             int32     `bson:"timezone,omitempty"`
	Token                string    `bson:"token,omitempty"`
	TokenExpiry          time.Time `bson:"token_expiry,omitempty"`
	ActivationCode       string    `bson:"activation_code,omitempty"`
	ActivationCodeExpiry time.Time `bson:"activation_code_expiry,omitempty"`
	IsActive             bool      `bson:"is_active,omitempty"`
	Created              time.Time `bson:"created,omitempty"`
	Modified             time.Time `bson:"modified,omitempty"`
}
