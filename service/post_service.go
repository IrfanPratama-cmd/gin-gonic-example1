package service

import (
	"gin-socmed/lib"
	"gin-socmed/model"
	"gin-socmed/repository"
)

type PostService interface {
	Create(req *model.PostRequest) error
}

type postService struct {
	repository repository.PostRepository
}

func NewPostService(r repository.PostRepository) *postService {
	return &postService{
		repository: r,
	}
}

func (s *postService) Create(req *model.PostRequest) error {
	post := model.Post{
		PostAPI: model.PostAPI{
			UserID: req.UserID,
			Tweet:  req.Tweet,
		},
	}

	if req.Picture != nil {
		post.PictureUrl = &req.Picture.Filename
	}

	if err := s.repository.Create(&post); err != nil {
		return &lib.InternalServerError{Message: err.Error()}
	}

	return nil
}
