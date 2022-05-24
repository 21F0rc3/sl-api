package models

import "gorm.io/gorm"

type SampleLiquid struct {
	gorm.Model
	SampleID     uint
	LiquidTypeID int
	Quantity     int
}
