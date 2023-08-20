package services

import (
	"thumtack_category/config"
	"thumtack_category/models"
	"thumtack_category/repositories"
)

func GetCategories() []models.Category {
	repository := repositories.NewUserRepositoryImpl(config.DB)
	return repository.GetCategories()
}