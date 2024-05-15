package service

import (
	"github.com/thiccpan/go-logger-benchmark/domain"
	"github.com/thiccpan/go-logger-benchmark/repository"
	"go.uber.org/zap"
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
	logger *zap.Logger
}

// AddItem implements IItemService.

func NewItemService(repo repository.IItemRepo, logger *zap.Logger) *itemService {
	return &itemService{
		repo:   repo,
		logger: logger,
	}
}

func (is *itemService) AddItem(newItem domain.Item) (domain.Item, error) {
	item, err := is.repo.AddItem(newItem)
	// is.logger.LogDebug("adding item on repository successfull")
	is.logger.Debug("adding item on repository successfull")
	if err != nil {
		// is.logger.LogDebug("error on adding item on repository")
		is.logger.Debug("error on adding item on repository")
		return item, err
	}
	return item, nil
}

func (is *itemService) GetItems() ([]domain.Item, error) {
	items, err := is.repo.GetItems()
	// is.logger.LogDebug("fetching items on repository successfull")
	is.logger.Debug("fetching items on repository successfull")
	if err != nil {
		// is.logger.LogDebug("error on fetching items on repository")
		is.logger.Debug("error on fetching items on repository")
		return nil, err
	}
	return items, nil
}

func (is *itemService) GetItem(id uint) (domain.Item, error) {
	item, err := is.repo.GetItem(uint(id))
	// is.logger.LogDebug("fetching item on repository successfull")
	is.logger.Debug("fetching item on repository successfull")
	if err != nil {
		// is.logger.LogDebug("error fetching item on repository")
		is.logger.Debug("error fetching item on repository")
		return item, err
	}
	return item, nil
}

func (is *itemService) UpdateItem(id uint, updateToItem domain.Item) (domain.Item, error) {
	item, err := is.repo.UpdateItem(uint(id), &updateToItem)
	// is.logger.LogDebug("updating item on repository successfull")
	is.logger.Debug("updating item on repository successfull")
	if err != nil {
		// is.logger.LogDebug("error updating item on repository")
		is.logger.Debug("error updating item on repository")
		return item, err
	}
	return item, nil
}

func (is *itemService) DeleteItem(id uint) (domain.Item, error) {
	item, err := is.repo.DeleteItem(id)
	// is.logger.LogDebug("deleting item on reporsitory successfull")
	is.logger.Debug("deleting item on reporsitory successfull")
	if err != nil {
		// is.logger.LogDebug("error deleting item on repository")
		is.logger.Debug("error deleting item on repository")
		return item, err
	}
	return item, nil
}
