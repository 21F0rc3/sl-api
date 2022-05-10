package models

import "github.com/jinzhu/gorm"

type OilBin struct {
	gorm.Model          // { ID, CreatedAt, UpdatedAt, DeletedAt, Name }
	CoordinateX float32 `json:"coordinate_x" gorm:"coordinate_x"`
	CoordinateY float32 `json:"coordinate_y" gorm:"coordinate_y"`
	Address     string  `json:"address"`
}
