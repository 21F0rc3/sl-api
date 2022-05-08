package services

import (
	"sl-api/models"
)

func InsertBottle(bottle *models.Bottle) error {
	if Database.Create(bottle).Error != nil {
		return badParamsError
	}

	return nil
}

func GetBottle(id string) (models.Bottle, error) {
	bottle := models.Bottle{}

	Database.First(&bottle, id)

	if bottle.ID == 0 {
		return bottle, notFoundError
	}

	return bottle, nil
}

func GetAllBottles() ([]models.Bottle, error) {
	var allBottles []models.Bottle

	Database.Find(&allBottles)

	if len(allBottles) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allBottles, emptyDbError
	}

	return allBottles, nil
}

func DeleteBottle(id string) error {
	return Database.Delete(&models.Bottle{}, id).Error
}
