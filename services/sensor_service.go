package services

import (
	"sl-api/models"
)

func InsertSensor(sensor *models.Sensor) error {
	if Database.Create(sensor).Error != nil {
		return badParamsError
	}

	return nil
}

func GetSensor(id string) (models.Sensor, error) {
	sensor := models.Sensor{}

	Database.First(&sensor, id)

	if sensor.ID == 0 {
		return sensor, notFoundError
	}

	return sensor, nil
}

func GetAllSensors() ([]models.Sensor, error) {
	var allSensors []models.Sensor

	Database.Find(&allSensors)

	if len(allSensors) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allSensors, emptyDbError
	}

	return allSensors, nil
}

func DeleteSensor(id string) error {
	return Database.Delete(&models.Sensor{}, id).Error
}
