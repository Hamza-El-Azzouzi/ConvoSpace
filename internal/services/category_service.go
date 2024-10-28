package services

import (
	"forum/internal/models"
	"forum/internal/repositories"
)

type CategoryService struct {
	CategorieRepo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{CategorieRepo: repo}
}

// GetAllCategories adds business logic layer
func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	categories, err := s.CategorieRepo.GetAllCategories()
	if err != nil {
		return nil, err
	}

	// Here you can add business logic like:
	// - Filtering categories based on user permissions
	// - Adding additional computed fields
	// - Caching results
	// - Logging or metrics
	// - Transforming data for the front-end

	return categories, nil
}
