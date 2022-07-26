package interfaces

import "github.com/chand19-af/digitels-template/app/models"

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
