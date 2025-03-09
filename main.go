package main

import (
	"GinProject1/config"
	"GinProject1/controllers"
	"GinProject1/error_handling"
	"GinProject1/models"
	"GinProject1/repositories"
	"GinProject1/router"
	"GinProject1/services"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)


func main() {
    log.Info().Msg("Server started")

    // for database
    db := config.DatabaseConnection()
    validate := validator.New()

    db.Table("tags").AutoMigrate(&models.Tags{})


    //Repository
    tagsRepository := repositories.NewTagsRepository(db)

    // service

    tagsService := services.NewTagsServiceImpl(tagsRepository, validate)

    //Controller 

    tagsController := controllers.NewtagsController(tagsService)


    // router
    router := router.NewRouter(tagsController)



    server := &http.Server{
        Addr: ":8080",
        Handler: router,
    }

    err := server.ListenAndServe()
    error_handling.ErrorPanice(err)



}