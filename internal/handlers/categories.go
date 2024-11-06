package handlers

import (
	"encoding/json"
	"net/http"

	"forum/internal/services"
	"forum/internal/utils"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAllCategories()
	if err != nil {
		utils.Error(w, 500)
	}

	json.NewEncoder(w).Encode(categories)
}
