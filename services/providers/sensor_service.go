package services

import (
	"sl-api/controllers"
	"sl-api/models"
	"sl-api/services"
	"strconv"
)

func InsertSensor(sensor *models.Sensor) error {
	if services.Database.Create(sensor).Error != nil {
		return services.BadParamsError
	}

	return nil
}

func GetSensor(id string) (models.Sensor, error) {
	sensor := models.Sensor{}

	services.Database.First(&sensor, id)

	if sensor.ID == 0 {
		return sensor, services.NotFoundError
	}

	return sensor, nil
}

func GetSensorByType(sType uint) (models.Sensor, error) {
	sensor := models.Sensor{}

	services.Database.First(&sensor, controllers.SENSOR_TYPE_ID_ATTR_NAME+" = "+strconv.Itoa(int(sType)))

	if sensor.ID == 0 {
		return sensor, services.NotFoundError
	}

	return sensor, nil
}

func GetAllSensors() ([]models.Sensor, error) {
	var allSensors []models.Sensor

	services.Database.Find(&allSensors)

	if len(allSensors) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allSensors, services.EmptyDbError
	}

	return allSensors, nil
}

func DeleteSensor(id string) error {
	return services.Database.Delete(&models.Sensor{}, id).Error
}
