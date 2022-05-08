package models

import "github.com/jinzhu/gorm"

type Bottle struct {
	gorm.Model
	QrCode string `gorm:"unique"`
	UserID int
}
