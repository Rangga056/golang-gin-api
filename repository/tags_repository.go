package repository

import "github.com/Rangga056/golang-gin-api/model"

type TagsRepository interface { //Create interface for Tags
	Save(tags model.Tags)                             //Saves a tags object & take model.tags as an argument
	Update(tags model.Tags)                           //update tags object
	Delete(tagsId int)                                //Delete tags object identified by tagsId
	FindById(tagsId int) (tags model.Tags, err error) //Find a tags by id & return it or return err
	FindAll() []model.Tags                            //Find all tags & return a slice containing all the tags
}
