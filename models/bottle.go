package models

import "github.com/jinzhu/gorm"

type Bottle struct {
	gorm.Model
	QrCode       string `json:"qr_code" gorm:"unique"`
	Firebase_UID string `json:"firebase_uid"`
}
