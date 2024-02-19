package repository

import (
	"gin-socmed/model"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *model.Post) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) Create(post *model.Post) error {
	err := r.db.Create(&post).Error
	return err
}
