package services

import (
	"sl-api/models"
)

func InsertReviewState(review_state *models.ReviewState) error {
	if Database.Create(review_state).Error != nil {
		return badParamsError
	}

	return nil
}

func GetReviewState(id string) (models.ReviewState, error) {
	review_state := models.ReviewState{}

	Database.First(&review_state, id)

	if review_state.ID == 0 {
		return review_state, notFoundError
	}

	return review_state, nil
}

func GetAllReviewStates() ([]models.ReviewState, error) {
	var allReviewStates []models.ReviewState

	Database.Find(&allReviewStates)

	if len(allReviewStates) == 0 { // Se o array estiver vazio, não existem utilizadores
		return allReviewStates, emptyDbError
	}

	return allReviewStates, nil
}

func DeleteReviewState(id string) error {
	return Database.Delete(&models.ReviewState{}, id).Error
}
