package services

import (
	"GinProject1/data/request"
	"GinProject1/data/response"
	"GinProject1/error_handling"
	"GinProject1/models"
	"GinProject1/repositories"

	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository repositories.TagsRepository
	validate       *validator.Validate
}
func NewTagsServiceImpl(tagRepository repositories.TagsRepository, validate *validator.Validate) *TagsServiceImpl {
	return &TagsServiceImpl{
		TagsRepository: tagRepository,
		validate:       validate,
	}
}

func (service *TagsServiceImpl) Create(tags request.CreateTagsRequest)  {
	err := service.validate.Struct(tags)
	error_handling.ErrorPanice(err)
	tagModel := models.Tags{
		Name: tags.Name,
	}
	service.TagsRepository.Save(tagModel)
}

func (service *TagsServiceImpl) Delete(tagsId int) {
	service.TagsRepository.Delete(tagsId)
}


func (service *TagsServiceImpl) FindAll() []response.TagsResponse {
	result := service.TagsRepository.FindAll()

	var tags []response.TagsResponse
	for _, value := range result {
		tag := response.TagsResponse {
			Id: value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}

	return tags
}


func (service *TagsServiceImpl) FindById(tagsId int) response.TagsResponse {
	tagData, err := service.TagsRepository.FindById(tagsId)
	error_handling.ErrorPanice(err)

	tagResponse := response.TagsResponse{
		Id: tagData.Id,
		Name: tagData.Name,
	}
	return tagResponse
}


func (service *TagsServiceImpl) Update(tags request.UpdateTagsRequest) {
	tagData, err := service.TagsRepository.FindById(tags.Id)
	error_handling.ErrorPanice(err)

	tagData.Name = tags.Name
	service.TagsRepository.Update(tagData)
}

