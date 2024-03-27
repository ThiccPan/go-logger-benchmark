package repository

import "github.com/thiccpan/go-logger-benchmark/domain"

type IAuthRepository interface {
	AddUser(user *domain.User) error
	GetUserByEmail(email string) domain.User
	UpdateUser(email string, newUser *domain.User) error
}
