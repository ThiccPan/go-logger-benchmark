package service

import (
	"errors"

	"github.com/thiccpan/go-logger-benchmark/domain"
	"github.com/thiccpan/go-logger-benchmark/logger"
	"github.com/thiccpan/go-logger-benchmark/repository"
)

type AuthService interface {
	Register(user domain.RegisterRequest) error
	Login(user domain.LoginRequest) (domain.User, error)
}

type authService struct {
	logger logger.Ilogger
	repo   repository.IAuthRepository
}

func NewAuthService(logger logger.Ilogger, repo repository.IAuthRepository) *authService {
	return &authService{
		logger: logger,
		repo:   repo,
	}
}

func (as *authService) Register(user domain.RegisterRequest) error {
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

func (as *authService) Login(user domain.LoginRequest) (domain.User, error) {
	foundUser, err := as.repo.GetUserByEmail(user.Email)
	if err != nil {
		return foundUser, err
	}
	return foundUser, nil
}
