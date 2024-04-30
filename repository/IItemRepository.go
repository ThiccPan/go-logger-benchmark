package repository

import "github.com/thiccpan/go-logger-benchmark/domain"

type IItemRepo interface {
	AddItem(post domain.Item) error
	GetItems() ([]domain.Item, error)
	GetItem(id uint) (domain.Item, error)
	UpdateItem(id uint, newPost *domain.Item) (domain.Item, error)
	DeleteItem(id uint) (domain.Item, error)
}
