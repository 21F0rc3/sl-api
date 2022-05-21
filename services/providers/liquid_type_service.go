package services

import (
	"sl-api/models"
	"sl-api/services"
)

func InsertLiquidType(liquid_type *models.LiquidType) error {
	if services.Database.Create(liquid_type).Error != nil {
		return services.BadParamsError
	}

	return nil
}

func GetLiquidType(id string) (models.LiquidType, error) {
	liquid_type := models.LiquidType{}

	services.Database.First(&liquid_type, id)

	if liquid_type.ID == 0 {
		return liquid_type, services.NotFoundError
	}

	return liquid_type, nil
}

func GetAllLiquidTypes() ([]models.LiquidType, error) {
	var allTypes []models.LiquidType

	services.Database.Find(&allTypes)

	if len(allTypes) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allTypes, services.EmptyDbError
	}

	return allTypes, nil
}

func DeleteLiquidType(id string) error {
	return services.Database.Delete(&models.LiquidType{}, id).Error
}
