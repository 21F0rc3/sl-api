package models

import "github.com/jinzhu/gorm"

type Sample struct {
	gorm.Model
	ReviewStateID int
}
