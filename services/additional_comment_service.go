package services

import (
	"sl-api/models"
)

func InsertAdditionalComment(comment *models.AdditionalComment) error {
	if Database.Create(comment).Error != nil {
		return badParamsError
	}

	return nil
}

func GetAdditionalComment(id string) (models.AdditionalComment, error) {
	comment := models.AdditionalComment{}

	Database.First(&comment, id)

	if comment.ID == 0 {
		return comment, notFoundError
	}

	return comment, nil
}

func GetAllAdditionalComments() ([]models.AdditionalComment, error) {
	var allComments []models.AdditionalComment

	Database.Find(&allComments)

	if len(allComments) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allComments, emptyDbError
	}

	return allComments, nil
}

func DeleteAdditionalComment(id string) error {
	return Database.Delete(&models.AdditionalComment{}, id).Error
}
