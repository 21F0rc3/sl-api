package services

import (
	"sl-api/models"
)

func InsertUser(user *models.User) error {
	if Database.Create(user).Error != nil {
		return badParamsError
	}

	return nil
}

func GetUser(id string) (models.User, error) {
	user := models.User{}

	// Vai buscar o primeiro user com o correspondente id
	Database.First(&user, id)

	if user.ID == 0 { // Caso não encontre nenhum utilizador retorna um erro
		return user, notFoundError
	}

	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	var allUsers []models.User

	Database.Find(&allUsers)

	if len(allUsers) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allUsers, emptyDbError
	}

	return allUsers, nil
}

func DeleteUser(id string) error {
	return Database.Delete(&models.User{}, id).Error
}
