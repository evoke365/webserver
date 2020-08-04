package profile

import "time"

type Profile struct {
	Email    string    `bson:"email"`
	Timezone int32     `bson:"timezone"`
	IsActive bool      `bson:"is_active"`
	Created  time.Time `bson:"created,omitempty"`
	Modified time.Time `bson:"modified,omitempty"`
}
