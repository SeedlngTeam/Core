package entity

import "github.com/google/uuid"

type Account struct {
	ID   uuid.UUID
	Name string
	User *User
}

func (a Account) GetID() uuid.UUID {
	return a.ID
}

func (a Account) GetName() string {
	return a.Name
}

func (a Account) GetUser() uuid.UUID {
	return a.User.ID
}
