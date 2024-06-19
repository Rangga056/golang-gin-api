package main

import (
	"fmt"
	"net/http"

	"github.com/Rangga056/golang-gin-api/config"
	"github.com/Rangga056/golang-gin-api/controller"
	"github.com/Rangga056/golang-gin-api/helper"
	"github.com/Rangga056/golang-gin-api/model"
	"github.com/Rangga056/golang-gin-api/repository"
	"github.com/Rangga056/golang-gin-api/router"
	"github.com/Rangga056/golang-gin-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Started server") //Logs information that the server has started
	//Database
	db := config.DatabaseConnection()
	fmt.Println("Database connect successfully", db != nil)
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	//Repository
	tagsRepository := repository.NewTagsRepositoryImpl(db)

	//Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	//Controller
	tagsController := controller.NewTagsController(tagsService)

	//Router
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
