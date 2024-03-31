package service

import (
	"errors"
	"fmt"

	"github.com/thiccpan/go-logger-benchmark/domain"
	"github.com/thiccpan/go-logger-benchmark/logger"
	"github.com/thiccpan/go-logger-benchmark/repository"
)

type AuthService struct {
	logger logger.Ilogger
	repo   repository.IAuthRepository
}

func NewAuthService(logger logger.Ilogger, repo repository.IAuthRepository) *AuthService {
	return &AuthService{
		logger: logger,
		repo:   repo,
	}
}

func (as *AuthService) Register(user domain.RegisterRequest) error {
	foundUser, _ := as.repo.GetUserByEmail(user.Email)
	if foundUser.Email != "" {
		return errors.New("email is not unique")
	}

	err := as.repo.AddUser(&domain.User{
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
	})
	if err != nil {
		return errors.New("error while registering user")
	}

	return nil
}

func (as *AuthService) Login(user domain.LoginRequest) (domain.User, error) {
	fmt.Println(user.Email)
	foundUser, err := as.repo.GetUserByEmail(user.Email)
	fmt.Println(foundUser)
	if err != nil {
		return foundUser, err
	}
	return foundUser, nil
}