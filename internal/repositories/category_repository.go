package repositories

import (
	"database/sql"
	"fmt"

	"forum/internal/models"
)

type CategoryRepository struct {
	DB *sql.DB
}

func (r *CategoryRepository) GetAllCategories() ([]models.Category, error) {
	rows, err := r.DB.Query("SELECT * FROM categories")
	if err != nil {
		return nil, fmt.Errorf("error querying categories: %v", err)
	}
	defer rows.Close()
	var categories []models.Category
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("error iterating categories: %v", err)
	}
	for rows.Next() {
		var cat models.Category
		err := rows.Scan(&cat.ID, &cat.Name)
		if err != nil {
			return nil, fmt.Errorf("error scanning category: %v", err)
		}
		categories = append(categories, cat)
	}

	return categories, nil
}
