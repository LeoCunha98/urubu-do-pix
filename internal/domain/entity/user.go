package entity

import "github.com/google/uuid"

type UserRepository interface {
	Create(user *User) error
	Find(id string) (*User, error)
	FindAll() ([]*User, error)
	Update(user *User) error
}

type User struct {
	ID   string
	Name string
}

func NewUser(name string) *User {
	return &User{
		ID:   uuid.New().String(),
		Name: name,
	}
}
