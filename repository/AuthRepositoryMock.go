package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/thiccpan/go-logger-benchmark/domain"
	"gorm.io/gorm"
)

type MockAuthRepo struct{
	users []domain.User
}

type Users struct {
    Users []domain.User `json:"users"`
}

func NewMockAuthRepo(db *gorm.DB) *MockAuthRepo {
	usersFile, err := os.OpenFile("usersdata.json", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("err opening user file:", err)
		os.Exit(1)
	}
	defer usersFile.Close()
	
	byteValue, _ := io.ReadAll(usersFile)
	users := Users{}
	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	repo := MockAuthRepo{
		users: users.Users,
	}
	return &repo
}

func (sau *MockAuthRepo) AddUser(user *domain.User) error {
	return nil
}

func (sau *MockAuthRepo) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	isFound := false 
	for _, v := range sau.users {
		if v.Email == email {
			isFound = true
			user = v
		}
	}
	if !isFound {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (sau *MockAuthRepo) UpdateUser(email string, newUser *domain.User) error {
	return nil
}
