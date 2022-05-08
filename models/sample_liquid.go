package models

import "gorm.io/gorm"

type SampleLiquid struct {
	gorm.Model
	SampleID     int
	LiquidTypeID int
	Quantity     int
}
