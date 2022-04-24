package domain

import "github.com/jinzhu/gorm"

type OilBin struct {
	gorm.Model
	CoordinateX float32 `json:"coordinateX" gorm:"coordinate_x"`
	CoordinateY float32 `json:"coordinateY" gorm:"coordinate_y"`
	Address     string
}
