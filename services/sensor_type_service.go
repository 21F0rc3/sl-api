package services

import (
	"sl-api/models"
)

func InsertSensorType(sensor_type *models.SensorType) error {
	if Database.Create(sensor_type).Error != nil {
		return badParamsError
	}

	return nil
}

func GetSensorType(id string) (models.SensorType, error) {
	sensor_type := models.SensorType{}

	Database.First(&sensor_type, id)

	if sensor_type.ID == 0 {
		return sensor_type, notFoundError
	}

	return sensor_type, nil
}

func GetAllSensorTypes() ([]models.SensorType, error) {
	var allTypes []models.SensorType

	Database.Find(&allTypes)

	if len(allTypes) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allTypes, emptyDbError
	}

	return allTypes, nil
}

func DeleteSensorType(id string) error {
	return Database.Delete(&models.SensorType{}, id).Error
}
