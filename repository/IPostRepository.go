package repository

import "github.com/thiccpan/go-logger-benchmark/domain"

type IPostRepo interface {
	AddPost(post *domain.Post) error
	GetPosts() ([]domain.Post, error)
	GetPost(id uint) (domain.Post, error)
	UpdatePost(id uint, newPost *domain.Post) (domain.Post, error)
	DeletePost(id uint) (domain.Post, error)
}
