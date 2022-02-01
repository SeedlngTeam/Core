package entity

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Phone    string
}
