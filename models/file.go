package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Name   string `gorm:"unique;not null"`
	Url    string `gorm:"not null"`
	UserId uint   `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserId"`
}
