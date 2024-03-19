package repository

import (
	"github.com/thiccpan/go-logger-benchmark/domain"
	"github.com/thiccpan/go-logger-benchmark/logger"
	"gorm.io/gorm"
)

type SQLitePostRepo struct {
	logger logger.Ilogger
	db     *gorm.DB
}

func NewSQLitePostRepo(logger logger.Ilogger, db *gorm.DB) *SQLitePostRepo {
	repo := SQLitePostRepo{
		db:     db,
		logger: logger,
	}
	return &repo
}

func (pr *SQLitePostRepo) AddPost(post *domain.Post) error {
	res := pr.db.Create(post)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (pr *SQLitePostRepo) GetPosts() ([]domain.Post, error) {
	posts := []domain.Post{}
	res := pr.db.Find(&posts)
	if res.Error != nil {
		return posts, res.Error
	}

	return posts, nil
}

func (pr *SQLitePostRepo) GetPost(id uint) (domain.Post, error) {
	post := domain.Post{}
	res := pr.db.First(&post, "id = ?", id)
	if res.Error != nil {
		return post, res.Error
	}
	
	return post, nil
}

func (pr *SQLitePostRepo) UpdatePost(id uint, newPost *domain.Post) (domain.Post, error) {
	pr.db.Updates(newPost)
	return domain.Post{}, nil
}

func (pr *SQLitePostRepo) DeletePost(id uint) (domain.Post, error) {
	pr.db.Delete(&domain.Post{}, id)
	return domain.Post{}, nil
}
