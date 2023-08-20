package repositories

import (
	"database/sql"
	"fmt"
	"thumtack_category/models"
)

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) userRepositoryImpl{
	return userRepositoryImpl {
		db,
	}
}

func (uRepository *userRepositoryImpl) GetCategories() []models.Category {
	fmt.Println("DB is  ")
	fmt.Println(uRepository.db)
	sql := "SELECT services.id, services.name "+
	"FROM services "+
	"INNER JOIN popular_services ON services.id=popular_services.service_id;"
	rows, err := uRepository.db.Query(sql)
	if err != nil {
		// handle this error better than this
		fmt.Println("error in DB.Query")
		panic(err)
	}
	defer rows.Close()
	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Id, &category.Name)
		if err != nil {
		  // handle this error
		  panic(err)
		}
		fmt.Println(category)
		categories = append(categories, category)
	  }
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
	  panic(err)
	}
	return categories
}