package category

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"

	"github.com/projects/go-api-exemple/api/responses"
	"github.com/projects/go-api-exemple/models"

	"github.com/projects/go-api-exemple/business"
	dtos "github.com/projects/go-api-exemple/dtos/categories"
)

// FindAll : Buscar todas categorias
func FindAll(w http.ResponseWriter, r *http.Request) {
	response := responses.Response{W: w, R: r}
	categories := business.Find()

	categoriesDto := []dtos.CategoryDTO{}
	for _, category := range categories {
		categoriesDto = append(categoriesDto, mapModelToDto(category))
	}

	response.Ok(categoriesDto)
}

// FindByID : Buscar categoria por id
func FindByID(w http.ResponseWriter, r *http.Request) {
	response := responses.Response{W: w, R: r}
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		response.BadRequest("Invalid id")
		return
	}

	category := business.FindByID(id)

	if category == nil {
		response.NotFound("Category not found")
	} else {
		response.Ok(mapModelToDto(*category))
	}
}

// Create : Criar nova categoria
func Create(w http.ResponseWriter, r *http.Request) {
	response := responses.Response{W: w, R: r}
	var categoryDto dtos.CategoryDTO
	json.NewDecoder(r.Body).Decode(&categoryDto)

	_, err := govalidator.ValidateStruct(categoryDto)
	if err != nil {
		response.BadRequest(err.Error())
	}

	category := business.Create(mapDtoToModel(categoryDto))
	response.Ok(mapModelToDto(category))
}

func mapDtoToModel(category dtos.CategoryDTO) models.Category {
	return models.Category{
		Name: category.Name,
	}
}

func mapModelToDto(category models.Category) dtos.CategoryDTO {
	return dtos.CategoryDTO{
		Name: category.Name,
		ID:   category.ID,
	}
}
