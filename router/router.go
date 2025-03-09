package router

import (
	"GinProject1/controllers"
	"net/http"
	"github.com/gin-gonic/gin"
)


func NewRouter(tagsController *controllers.TagController) *gin.Engine{
	router :=  gin.Default()

	router.GET("", func(ctx *gin.Context){
        ctx.JSON(http.StatusOK, "welcome")
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