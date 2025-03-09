package repositories

import (
	"GinProject1/data/request"
	"GinProject1/error_handling"
	"GinProject1/models"
	"errors"

	"gorm.io/gorm"
)


type TagsRepositoryImpl struct {
	Db *gorm.DB
}



func NewTagsRepository(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}


func (r *TagsRepositoryImpl) Delete(tagsId int) {
	var tags models.Tags

	result := r.Db.Where("id = ?", tagsId).Delete(&tags)
	error_handling.ErrorPanice(result.Error)
}

func (r *TagsRepositoryImpl) FindAll() []models.Tags {

	var tags []models.Tags

	result := r.Db.Find(&tags)
	error_handling.ErrorPanice(result.Error)

	return tags
}


func (r *TagsRepositoryImpl) FindById(tagsId int) (tegs models.Tags, err error) {
	var tag models.Tags
	result := r.Db.Find(&tag, tagsId)

	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag in not found")
	}
}

func (r *TagsRepositoryImpl) Save(tags models.Tags) {
	result := r.Db.Create(&tags)
	error_handling.ErrorPanice(result.Error)
}

func (r *TagsRepositoryImpl) Update(tags models.Tags){
	var updateTag = request.UpdateTagsRequest{
		Id : tags.Id,
		Name : tags.Name,
	}

	result := r.Db.Model(&tags).Updates(updateTag)
	error_handling.ErrorPanice(result.Error)
}