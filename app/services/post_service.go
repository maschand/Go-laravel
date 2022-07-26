package services

import (
	"github.com/chand19-af/digitels-template/app/models"
	"github.com/chand19-af/digitels-template/app/repositories"
)

func getPost() models.Post {
	repo := repositories.NewPost()
	return
}

//func (service PostService) getPost() models.Post {
//	return models.Post{}
//}
