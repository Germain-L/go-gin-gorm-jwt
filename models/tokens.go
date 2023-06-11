package models

import "gorm.io/gorm"

type RefreshToken struct {
	gorm.Model
	Token  string `gorm:"unique;not null"`
	UserId uint   `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserId"`
	Valid  bool   `gorm:"not null"`
}

func (r *RefreshToken) IsValid() bool {
	return r.Valid
}

func (r *RefreshToken) Invalidate() {
	r.Valid = false
}
