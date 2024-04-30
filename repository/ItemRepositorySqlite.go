package repository

import (
	"github.com/thiccpan/go-logger-benchmark/domain"
	"github.com/thiccpan/go-logger-benchmark/logger"
	"gorm.io/gorm"
)

type SQLiteItemRepo struct {
	logger logger.Ilogger
	db     *gorm.DB
}

func NewSQLiteItemRepo(logger logger.Ilogger, db *gorm.DB) *SQLiteItemRepo {
	repo := &SQLiteItemRepo{
		db:     db,
		logger: logger,
	}
	return repo
}

func (pr *SQLiteItemRepo) AddItem(item domain.Item) error {
	res := pr.db.Create(&item)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (pr *SQLiteItemRepo) GetItems() ([]domain.Item, error) {
	items := []domain.Item{}
	res := pr.db.Find(&items)
	if res.Error != nil {
		return items, res.Error
	}

	return items, nil
}

func (pr *SQLiteItemRepo) GetItem(id uint) (domain.Item, error) {
	item := domain.Item{}
	res := pr.db.First(&item, "id = ?", id)
	if res.Error != nil {
		return item, res.Error
	}

	return item, nil
}

func (pr *SQLiteItemRepo) UpdateItem(id uint, newItem *domain.Item) (domain.Item, error) {
	newItem.ID = id
	res := pr.db.Updates(newItem)
	if res.Error != nil {
		return domain.Item{}, res.Error
	}
	return *newItem, nil
}

func (pr *SQLiteItemRepo) DeleteItem(id uint) (domain.Item, error) {
	res := pr.db.Delete(&domain.Item{}, id)
	if res.Error != nil {
		return domain.Item{}, res.Error
	}
	return domain.Item{}, nil
}
