package repositories

import (
	"fmt"
	"thumtack_category/config"
	"thumtack_category/models"
)

func GetCategories() []models.Category {
	fmt.Println("config.DB is  ")
	fmt.Println(config.DB)
	sql := "SELECT services.id, services.name "+
	"FROM services "+
	"INNER JOIN popular_services ON services.id=popular_services.service_id;"
	rows, err := config.DB.Query(sql)
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