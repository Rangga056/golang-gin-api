package repository

import (
	"errors"

	"github.com/Rangga056/golang-gin-api/data/request"
	"github.com/Rangga056/golang-gin-api/helper"
	"github.com/Rangga056/golang-gin-api/model"
	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	//Implements the TagsRepository interface and holds a pointer to a gorm.DB
	Db *gorm.DB
}

func NewTagsRepositoryImpl(Db *gorm.DB) TagsRepository {
	//Construct and return a new instance of TagsRepositoryImpl
	return &TagsRepositoryImpl{Db: Db}
}

// Delete implements TagsRepository.
func (t *TagsRepositoryImpl) Delete(tagsId int) {
	var tags model.Tags
	result := t.Db.Where("id = ?", tagsId).Delete(&tags)
	helper.ErrorPanic(result.Error)
}

// FindById implements TagsRepository.
func (t *TagsRepositoryImpl) FindById(tagsId int) (tags model.Tags, err error) {
	var tag model.Tags
	result := t.Db.Find(&tag, tagsId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

// Save implements TagsRepository.
func (t *TagsRepositoryImpl) Save(tags model.Tags) {
	result := t.Db.Create(&tags)
	helper.ErrorPanic(result.Error)
}

// Update implements TagsRepository.
func (t *TagsRepositoryImpl) Update(tags model.Tags) {
	var updateTag = request.UpdateTagRequest{
		Id:   tags.Id,
		Name: tags.Name,
	}
	result := t.Db.Model(&tags).Updates(updateTag)
	helper.ErrorPanic(result.Error)
}

// findAll implements TagsRepository.
func (t *TagsRepositoryImpl) findAll() []model.Tags {
	var tags []model.Tags
	result := t.Db.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}
