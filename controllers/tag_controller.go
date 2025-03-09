package controllers

import (
	"GinProject1/data/request"
	"GinProject1/data/response"
	"GinProject1/error_handling"
	"GinProject1/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagController struct {
	TagsService services.TagsService
}

func NewtagsController(service services.TagsService) *TagController {
	return &TagController {
		TagsService: service,
	}
}


func(controller *TagController) Create(ctx *gin.Context){
	CreateTagsRequest := request.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&CreateTagsRequest)
	error_handling.ErrorPanice(err)


	controller.TagsService.Create(CreateTagsRequest)
	webResponse := response.Response{
		Code: http.StatusOK,
		Status: "Ok",
		Data: nil,
	}
	ctx.Header("Context-Type", "application/json")

	ctx.JSON(http.StatusOK, webResponse)
}


func(controller *TagController) Update(ctx *gin.Context){
	updateTagsRequest := request.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagsRequest)
	error_handling.ErrorPanice(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	error_handling.ErrorPanice(err)

	updateTagsRequest.Id = id

	controller.TagsService.Update(updateTagsRequest)

	webResponse := response.Response{
		Code: http.StatusOK,
		Status: "Ok",
		Data: nil,
	}
	ctx.Header("Context-Type", "application/json")

	ctx.JSON(http.StatusOK, webResponse)
}


func(controller *TagController) Delete(ctx *gin.Context){
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	error_handling.ErrorPanice(err)

	controller.TagsService.Delete(id)

	webResponse := response.Response{
		Code: http.StatusOK,
		Status: "Ok",
		Data: nil,
	}
	ctx.Header("Context-Type", "application/json")

	ctx.JSON(http.StatusOK, webResponse)
}


func(controller *TagController) FindById(ctx *gin.Context){
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	
	error_handling.ErrorPanice(err)

	tagResponse := controller.TagsService.FindById(id)
	

	webResponse := response.Response{
		Code: http.StatusOK,
		Status: "Ok",
		Data: tagResponse,
	}
	ctx.Header("Context-Type", "application/json")

	ctx.JSON(http.StatusOK, webResponse)
}

func(controller *TagController) FindAll(ctx *gin.Context){
	tagResponse := controller.TagsService.FindAll()


	webResponse := response.Response{
		Code: http.StatusOK,
		Status: "Ok",
		Data: tagResponse,
	}
	ctx.Header("Context-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}



