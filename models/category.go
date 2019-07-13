package models

import (
	"github.com/jinzhu/gorm"
)

// Category : Modelo de categorias
type Category struct {
	gorm.Model
	Name string `gorm:"not null"`
}
