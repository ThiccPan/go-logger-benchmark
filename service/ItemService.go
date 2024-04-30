package service

import (
	"github.com/thiccpan/go-logger-benchmark/domain"
	"github.com/thiccpan/go-logger-benchmark/repository"
)

type IItemService interface {
	AddItem(newItem domain.Item) error
}

type itemService struct {
	repo repository.IItemRepo
}

// AddItem implements IItemService.

func NewItemService(repo repository.IItemRepo) *itemService {
	return &itemService{
		repo: repo,
	}
}

func (is *itemService) AddItem(newItem domain.Item) error {
	if err := is.repo.AddItem(newItem); err != nil {
		return err
	}
	return nil
}
