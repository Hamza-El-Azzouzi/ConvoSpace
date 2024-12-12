package services

import (
	"forum/internal/models"
	"forum/internal/repositories"
)

type CategoryService struct {
	CategorieRepo *repositories.CategoryRepository
}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	categories, err := s.CategorieRepo.GetAllCategories()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
