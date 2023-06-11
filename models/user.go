package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string         `gorm:"unique;not null"`
	Password string         `gorm:"not null"`
	Files    []File         `gorm:"foreignKey:UserId;AssociationForeignKey:ID"`
	Tokens   []RefreshToken `gorm:"foreignKey:UserId;AssociationForeignKey:ID"`
}

func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(bytes)

	return nil
}
