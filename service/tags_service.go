package service

import (
	"github.com/Rangga056/golang-gin-api/data/request"
	"github.com/Rangga056/golang-gin-api/data/response"
)

type TagsService interface {
	//NOTE:Create Service methods for CRUD operations
	Create(tags request.CreateTagsRequests)
	Update(tags request.UpdateTagRequest)
	Delete(tagsId int)
	FindById(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}
