package auth

import "time"

type Model struct {
	s Store
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

func NewModel(s Store) *Model {
	return &Model{
		s: s,
	}
}

func (m *Model) GetUser(id string) (*User, error) {
	user := &User{}
	_, err := m.s.Get(id, user)
	return user, err
}

func (m *Model) SetUser(id string, user *User) error {
	return nil
}
