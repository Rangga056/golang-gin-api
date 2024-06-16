package model

// Create a model for Tags
type Tags struct {
	Id   int    `gorm:"type:int;primary_key"` //add a struct tag from gorm
	Name string `gorm:"type:varchar(255)"`
}
