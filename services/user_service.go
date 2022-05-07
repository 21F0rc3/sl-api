package services

import (
	"sl-api/models"
)

func InsertUser(user *models.User) error {
	if Database.Create(user).Error != nil {
		return nil
	}

	return nil
}
