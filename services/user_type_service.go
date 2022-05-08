package services

import (
	"sl-api/models"
)

func InsertUserType(user_type *models.UserType) error {
	if Database.Create(user_type).Error != nil {
		return badParamsError
	}

	return nil
}

func GetUserType(id string) (models.UserType, error) {
	user_type := models.UserType{}

	Database.First(&user_type, id)

	if user_type.ID == 0 {
		return user_type, notFoundError
	}

	return user_type, nil
}

func GetAllUserTypes() ([]models.UserType, error) {
	var allTypes []models.UserType

	Database.Find(&allTypes)

	if len(allTypes) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allTypes, emptyDbError
	}

	return allTypes, nil
}

func DeleteUserType(id string) error {
	return Database.Delete(&models.UserType{}, id).Error
}
