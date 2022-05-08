package models

import "github.com/jinzhu/gorm"

type SensorType struct {
	gorm.Model
	Sensor string `gorm:"unique"`
}
