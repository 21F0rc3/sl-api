package services

import (
	"sl-api/models"
)

func InsertLiquidType(liquid_type *models.LiquidType) error {
	if Database.Create(liquid_type).Error != nil {
		return badParamsError
	}

	return nil
}

func GetLiquidType(id string) (models.LiquidType, error) {
	liquid_type := models.LiquidType{}

	Database.First(&liquid_type, id)

	if liquid_type.ID == 0 {
		return liquid_type, notFoundError
	}

	return liquid_type, nil
}

func GetAllLiquidTypes() ([]models.LiquidType, error) {
	var allTypes []models.LiquidType

	Database.Find(&allTypes)

	if len(allTypes) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allTypes, emptyDbError
	}

	return allTypes, nil
}

func DeleteLiquidType(id string) error {
	return Database.Delete(&models.LiquidType{}, id).Error
}
