package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title   string
	Status  string
	Done    bool
	GroupID uint
}

type Group struct {
	gorm.Model
	Name  string
	Tasks []Task `gorm:"foreignKey:GroupID"`
	Users []User `gorm:"many2many:user_groups;"`
}
