package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string `gorm:"unique"`
	Password string
	Groups   []Group `gorm:"many2many:user_groups;"`
}
