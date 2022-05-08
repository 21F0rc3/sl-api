package services

import (
	"sl-api/models"
)

func InsertSensorReading(reading *models.SensorReading) error {
	if Database.Create(reading).Error != nil {
		return badParamsError
	}

	return nil
}

func GetSensorReading(id string) (models.SensorReading, error) {
	reading := models.SensorReading{}

	Database.First(&reading, id)

	if reading.ID == 0 {
		return reading, notFoundError
	}

	return reading, nil
}

func GetAllSensorReadings() ([]models.SensorReading, error) {
	var allReadings []models.SensorReading

	Database.Find(&allReadings)

	if len(allReadings) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allReadings, emptyDbError
	}

	return allReadings, nil
}

func DeleteSensorReading(id string) error {
	return Database.Delete(&models.SensorReading{}, id).Error
}
