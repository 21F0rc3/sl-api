package services

import (
	"sl-api/models"
	"sl-api/services"
)

func InsertSensorReading(reading *models.SensorReading) error {
	if services.Database.Create(reading).Error != nil {
		return services.BadParamsError
	}

	return nil
}

func GetSensorReading(id string) (models.SensorReading, error) {
	reading := models.SensorReading{}

	services.Database.First(&reading, id)

	if reading.ID == 0 {
		return reading, services.NotFoundError
	}

	return reading, nil
}

func GetAllSensorReadings() ([]models.SensorReading, error) {
	var allReadings []models.SensorReading

	services.Database.Find(&allReadings)

	if len(allReadings) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allReadings, services.EmptyDbError
	}

	return allReadings, nil
}

func DeleteSensorReading(id string) error {
	return services.Database.Delete(&models.SensorReading{}, id).Error
}
