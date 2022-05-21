package services

import (
	"sl-api/models"
	"sl-api/services"
)

func InsertBottle(bottle *models.Bottle) error {
	if services.Database.Create(bottle).Error != nil {
		return services.BadParamsError
	}

	return nil
}

func GetBottle(id string) (models.Bottle, error) {
	bottle := models.Bottle{}

	services.Database.First(&bottle, id)

	if bottle.ID == 0 {
		return bottle, services.NotFoundError
	}

	return bottle, nil
}

func GetAllBottles() ([]models.Bottle, error) {
	var allBottles []models.Bottle

	services.Database.Find(&allBottles)

	if len(allBottles) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allBottles, services.EmptyDbError
	}

	return allBottles, nil
}

func DeleteBottle(id string) error {
	return services.Database.Delete(&models.Bottle{}, id).Error
}
