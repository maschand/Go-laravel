package interfaces

import "github.com/create-go-app/fiber-go-template/app/models"

type (
	PostInterface interface {
		getPost() models.Post
		createPost() models.Post
		updatePost() models.Post
		deletePost() int
	}

	PostRepository interface {
		getPost() models.Post
		createPost() models.Post
		updatePost() models.Post
		deletePost() int
	}
)
