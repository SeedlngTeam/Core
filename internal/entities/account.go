package entity

import (
	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

func (a Account) GetID() uint {
	return a.ID
}

func (a Account) GetName() string {
	return a.Name
}

func (a Account) GetUser() uint {
	return a.UserID
}
