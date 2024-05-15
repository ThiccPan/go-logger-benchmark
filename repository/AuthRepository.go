package repository

import (
	"github.com/thiccpan/go-logger-benchmark/domain"
	"gorm.io/gorm"
)

type SQLiteAuthRepo struct {
	db     *gorm.DB
}

func NewSQLiteAuthRepo(db *gorm.DB) *SQLiteAuthRepo {
	repo := SQLiteAuthRepo{
		db:     db,
	}
	return &repo
}

func (sau *SQLiteAuthRepo) AddUser(user *domain.User) error {
	db := sau.db.Create(user)
	if err := db.Error; err != nil {
		return err
	}
	return nil
}

func (sau *SQLiteAuthRepo) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	db := sau.db.First(&user, "email = ?", email)
	if err := db.Error; err != nil {
		return user, err
	}
	return user, nil
}

func (sau *SQLiteAuthRepo) UpdateUser(email string, newUser *domain.User) error {
	return nil
}
