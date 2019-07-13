package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/projects/go-api-exemple/models"
	"github.com/projects/go-api-exemple/utils"

	"github.com/projects/go-api-exemple/config"

	"github.com/projects/go-api-exemple/api/router"
)

func main() {
	config.LoadEnvVars()
	router := router.ConfigureRouter()

	utils.DB().AutoMigrate(&models.Category{})

	log.Printf("Api started at port :%s", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Port), router))
}
