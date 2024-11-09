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
	if r.Method != http.MethodGet{
		utils.Error(w,405)
	}
	categories, err := h.service.GetAllCategories()
	if err != nil {
		utils.Error(w, 500)
	}

	json.NewEncoder(w).Encode(categories)
}
