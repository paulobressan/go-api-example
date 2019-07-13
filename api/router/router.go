package router

import (
	"github.com/gorilla/mux"
	"github.com/projects/go-api-exemple/api/category"
)

// ConfigureRouter : Configurar roteamento da API
func ConfigureRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/categories", category.FindAll).Methods("GET")
	r.HandleFunc("/categories/{id}", category.FindByID).Methods("GET")
	r.HandleFunc("/categories", category.Create).Methods("POST")

	return r
}
