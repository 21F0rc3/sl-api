package models

import "github.com/jinzhu/gorm"

type SensorReading struct {
	gorm.Model
	SensorID uint
	Value    float64
	SampleID uint
}
