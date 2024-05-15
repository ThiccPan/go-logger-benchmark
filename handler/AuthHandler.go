package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thiccpan/go-logger-benchmark/domain"
	"github.com/thiccpan/go-logger-benchmark/helper"
	"github.com/thiccpan/go-logger-benchmark/service"
)

type AuthHandler struct {
	service service.AuthService
	jwt     helper.AuthJWT
}

func NewAuthHandler(service service.AuthService, jwt helper.AuthJWT) *AuthHandler {
	return &AuthHandler{
		service: service,
		jwt:     jwt,
	}
}

func (ah *AuthHandler) RegisterHandler(c echo.Context) error {
	loginReq := domain.RegisterRequest{}
	if err := c.Bind(&loginReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status": 400,
		})
	}

	if err := ah.service.Register(loginReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status": 400,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status": 200,
	})
}

func (ah *AuthHandler) LoginHandler(c echo.Context) error {
	var user domain.LoginRequest
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "invalid login req",
		})
	}

	userCred, err := ah.service.Login(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "login service err",
		})
	}

	token, err := ah.jwt.GenerateToken(userCred.ID, userCred.Username, userCred.Email, userCred.Password)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status": 200,
		"user":   user,
		"token":  token,
	})
}
