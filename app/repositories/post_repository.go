package repositories

import (
	"github.com/chand19-af/digitels-template/app/interfaces"
	"github.com/chand19-af/digitels-template/app/models"
)

type Post struct {
	PostRepository interfaces.PostRepository
}

func NewPost(postInterface interfaces.PostRepository) Post {
	return Post{
		PostRepository: postInterface,
	}
}

func (repository Post) getPost() models.Post {
	body := repository.getPost()
	return body
}

func (repository Post) createPost() models.Post {
	body := repository.createPost()
	return body
}

func (repository Post) updatePost() models.Post {
	body := models.Post{}
	return body
}

func (repository Post) deletePost() models.Post {
	body := models.Post{}
	return body
}
