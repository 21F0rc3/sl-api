package services

import (
	"sl-api/models"
	"sl-api/services"
)

func InsertUserType(userType *models.UserType) error {
	if services.Database.Create(userType).Error != nil {
		return services.BadParamsError
	}

	return nil
}

func GetUserType(id string) (models.UserType, error) {
	user_type := models.UserType{}

	services.Database.First(&user_type, id)

	if user_type.ID == 0 {
		return user_type, services.NotFoundError
	}

	return user_type, nil
}

func GetAllUserTypes() ([]models.UserType, error) {
	var allTypes []models.UserType

	services.Database.Find(&allTypes)

	if len(allTypes) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allTypes, services.EmptyDbError
	}

	return allTypes, nil
}

func DeleteUserType(id string) error {
	return services.Database.Delete(&models.UserType{}, id).Error
}
