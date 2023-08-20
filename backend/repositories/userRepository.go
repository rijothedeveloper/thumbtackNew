package repositories

import "thumtack_category/models"

type userRepository interface {
	GetCategories() []models.Category
}