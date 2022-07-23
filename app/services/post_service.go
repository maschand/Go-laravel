package services

import (
	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/create-go-app/fiber-go-template/app/repositories"
)

func getPost() models.Post {
	repo := repositories.NewPost()
	return
}

//func (service PostService) getPost() models.Post {
//	return models.Post{}
//}
