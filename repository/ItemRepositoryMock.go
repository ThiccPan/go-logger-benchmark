package repository

import (
	"errors"

	"github.com/thiccpan/go-logger-benchmark/domain"
)

type MockItemRepo struct{}

var items = map[uint]domain.Item{
	1: {
		ID:    1,
		Name:  "item 1",
		Stock: 10,
	},
	2: {
		ID:    2,
		Name:  "item 2",
		Stock: 10,
	},
	3: {
		ID:    3,
		Name:  "item 3",
		Stock: 10,
	},
	4: {
		ID:    4,
		Name:  "item 4",
		Stock: 10,
	},
	5: {
		ID:    5,
		Name:  "item 5",
		Stock: 10,
	},
}

func NewMockItemRepo() *MockItemRepo {
	repo := &MockItemRepo{}
	return repo
}

func (pr *MockItemRepo) AddItem(item domain.Item) (domain.Item, error) {
	return item, nil
}

func (pr *MockItemRepo) GetItems() ([]domain.Item, error) {
	itemsData := []domain.Item{}
	for _, v := range items {
		itemsData = append(itemsData, v)
	}

	return itemsData, nil
}

func (pr *MockItemRepo) GetItem(id uint) (domain.Item, error) {
	item, ok := items[id]
	if !ok {
		return domain.Item{}, errors.New("item not found!")
	}

	return item, nil
}

func (pr *MockItemRepo) UpdateItem(id uint, newItem *domain.Item) (domain.Item, error) {
	return *newItem, nil
}

func (pr *MockItemRepo) DeleteItem(id uint) (domain.Item, error) {
	return domain.Item{}, nil
}
