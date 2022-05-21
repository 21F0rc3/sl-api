package services

import (
	"sl-api/models"
	"sl-api/services"
)

func InsertSensorType(sensor_type *models.SensorType) error {
	if services.Database.Create(sensor_type).Error != nil {
		return services.BadParamsError
	}

	return nil
}

func GetSensorType(id string) (models.SensorType, error) {
	sensor_type := models.SensorType{}

	services.Database.First(&sensor_type, id)

	if sensor_type.ID == 0 {
		return sensor_type, services.NotFoundError
	}

	return sensor_type, nil
}

func GetAllSensorTypes() ([]models.SensorType, error) {
	var allTypes []models.SensorType

	services.Database.Find(&allTypes)

	if len(allTypes) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allTypes, services.EmptyDbError
	}

	return allTypes, nil
}

func DeleteSensorType(id string) error {
	return services.Database.Delete(&models.SensorType{}, id).Error
}
