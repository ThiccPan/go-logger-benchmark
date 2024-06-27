package service

import (
	"github.com/sirupsen/logrus"
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
	repo   repository.IItemRepo
	logger *logrus.Logger
}

func NewItemService(repo repository.IItemRepo, logger *logrus.Logger) *itemService {
	return &itemService{
		repo:   repo,
		logger: logger,
	}
}

func (is *itemService) AddItem(newItem domain.Item) (domain.Item, error) {
	item, err := is.repo.AddItem(newItem)
	if err != nil {
		is.logger.Debug("error on adding item on repository")
		return item, err
	}
	is.logger.Debug("adding item on repository successfull")
	return item, nil
}

func (is *itemService) GetItems() ([]domain.Item, error) {
	items, err := is.repo.GetItems()
	if err != nil {
		is.logger.Debug("error on fetching items on repository")
		return nil, err
	}
	is.logger.Debug("fetching items on repository successfull")
	return items, nil
}

func (is *itemService) GetItem(id uint) (domain.Item, error) {
	item, err := is.repo.GetItem(uint(id))
	if err != nil {
		is.logger.Debug("error fetching item on repository")
		return item, err
	}
	is.logger.Debug("fetching item on repository successfull")
	return item, nil
}

func (is *itemService) UpdateItem(id uint, updateToItem domain.Item) (domain.Item, error) {
	item, err := is.repo.UpdateItem(uint(id), &updateToItem)
	if err != nil {
		is.logger.Debug("error updating item on repository")
		return item, err
	}
	is.logger.Debug("updating item on repository successfull")
	return item, nil
}

func (is *itemService) DeleteItem(id uint) (domain.Item, error) {
	item, err := is.repo.DeleteItem(id)
	if err != nil {
		is.logger.Debug("error deleting item on repository")
		return item, err
	}
	is.logger.Debug("deleting item on reporsitory successfull")
	return item, nil
}
