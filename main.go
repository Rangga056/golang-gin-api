package main

import (
	"net/http"

	"github.com/Rangga056/golang-gin-api/helper"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Started server") //Logs information that the server has started
	//Initiate a new Gin router with default middleware(logger and recovery middleware)
	routes := gin.Default()

	//Set a route for HTTP GET request on "/"
	routes.GET("", func(ctx *gin.Context) { //Takes a gin context as a parameter
		ctx.JSON(http.StatusOK, "Welcome")
	})

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
