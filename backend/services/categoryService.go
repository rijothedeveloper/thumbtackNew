package services

import (
	"thumtack_category/models"
	"thumtack_category/repositories"
)

func GetCategories() []models.Category {
	return repositories.GetCategories()
}