package router

import (
	"net/http"

	"github.com/Rangga056/golang-gin-api/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(tagsController *controller.TagsController) *gin.Engine {

	//Initiate a new Gin router with default middleware(logger and recovery middleware)
	router := gin.Default()

	//Set a route for HTTP GET request on "/"
	router.GET("", func(ctx *gin.Context) { //Takes a gin context as a parameter
		ctx.JSON(http.StatusOK, "Welcome")
	})

	baseRouter := router.Group("/api")
	tagsRouter := baseRouter.Group("/tags")
	tagsRouter.GET("", tagsController.FindAll)
	tagsRouter.GET("/:tagId", tagsController.FindById)
	tagsRouter.POST("", tagsController.Create)
	tagsRouter.PATCH("/:tagId", tagsController.Update)
	tagsRouter.DELETE("/:tagId", tagsController.Delete)

	return router
}
