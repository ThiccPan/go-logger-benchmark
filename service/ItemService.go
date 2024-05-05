package service

import (
	"github.com/thiccpan/go-logger-benchmark/domain"
	"github.com/thiccpan/go-logger-benchmark/repository"
)

type IItemService interface {
	AddItem(newItem domain.Item) (domain.Item, error)
	GetItems() ([]domain.Item, error)
	GetItem(id uint) (domain.Item, error)
	UpdateItem(id uint, updateToItem domain.Item) (domain.Item, error)
	DeleteItem(id uint) (domain.Item, error)
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

func (is *itemService) AddItem(newItem domain.Item) (domain.Item, error) {
	item, err := is.repo.AddItem(newItem)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (is *itemService) GetItems() ([]domain.Item, error) {
	items, err := is.repo.GetItems()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (is *itemService) GetItem(id uint) (domain.Item, error) {
	item, err := is.repo.GetItem(uint(id))
	if err != nil {
		return item, err
	}
	return item, nil
}

func (is *itemService) UpdateItem(id uint, updateToItem domain.Item) (domain.Item, error) {
	item, err := is.repo.UpdateItem(uint(id), &updateToItem)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (is *itemService) DeleteItem(id uint) (domain.Item, error) {
	item, err := is.repo.DeleteItem(id)
	if err != nil {
		return item, err
	}
	return item, nil
}
