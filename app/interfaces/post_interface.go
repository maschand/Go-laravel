package interfaces

import "gitlab.com/d6825/golang_template/app/models"

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
