package services

import (
	"github.com/abu-sayed/go-rest/models"
)

func GetPost(id int) models.Post{
	posts := models.Post{ID:1, Title:"Hello", Content:"Hello World"}
	return posts
}