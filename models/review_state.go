package models

import "github.com/jinzhu/gorm"

type ReviewState struct {
	gorm.Model
	State   string `gorm:"unique"`
	Message string `gorm:"not null"`
	Points  int
}
