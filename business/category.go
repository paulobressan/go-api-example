package business

import (
	"github.com/projects/go-api-exemple/models"
	"github.com/projects/go-api-exemple/utils"
)

// Find : Retorna todas categorias
func Find() []models.Category {
	var categories []models.Category
	utils.DB().Find(&categories)
	return categories
}

// FindByID : Retorna todas categorias
func FindByID(id int) *models.Category {
	var category models.Category
	err := utils.DB().Where("id = ?", id).Find(&category).Error
	if err != nil {
		return nil
	}
	return &category
}

// Create : Criar e retornar nova categoria
func Create(category models.Category) models.Category {
	utils.DB().Create(&category)
	return category
}
