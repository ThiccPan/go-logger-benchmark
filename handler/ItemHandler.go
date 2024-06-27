package handler

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/thiccpan/go-logger-benchmark/domain"
	"github.com/thiccpan/go-logger-benchmark/helper"
	"github.com/thiccpan/go-logger-benchmark/repository"
	"github.com/thiccpan/go-logger-benchmark/service"
	"go.uber.org/zap"
)

type ItemHandler struct {
	Repo    repository.IItemRepo
	service service.IItemService
	logger  *zap.Logger
}

func NewItemHandler(repo repository.IItemRepo, service service.IItemService, logger *zap.Logger) *ItemHandler {
	return &ItemHandler{Repo: repo, service: service, logger: logger}
}

type addItemRequest struct {
	Name  string `json:"name"`
	Stock uint   `json:"stock"`
}

type updateItemRequest struct {
	Name  string `json:"name"`
	Stock uint   `json:"stock"`
}

func (ph *ItemHandler) GetItemsHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*helper.JwtCustomClaims)

	items, err := ph.service.GetItems()
	if err != nil {
		ph.logger.Info("error retrieving items")
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}

	// cetak pesan log dengan field email,
	ph.logger.Info("fetching all items successfull", zap.String("email", claims.Email))

	return c.JSON(http.StatusOK, map[string]any{
		"item": items,
	})
}
func (ph *ItemHandler) GetItemHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*helper.JwtCustomClaims)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ph.logger.Info("failed to convert id")
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "invalid id, use integer id",
		})
	}

	item, err := ph.service.GetItem(uint(id))
	if err != nil {
		ph.logger.Info("error retrieving items")
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	// cetak pesan log dengan field email, item_id
	ph.logger.Info(
		"getting item successfull",
		zap.String("email", claims.Email),
		zap.Int("item_id", id),
	)

	return c.JSON(http.StatusOK, map[string]any{
		"item": item,
	})
}

func (ph *ItemHandler) UpdateItemHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*helper.JwtCustomClaims)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ph.logger.Info("failed to convert id")
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "invalid id, use integer id",
		})
	}

	request := &updateItemRequest{}
	err = c.Bind(request)
	if err != nil {
		ph.logger.Info("bad request")
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	newItem := domain.Item{
		Name:  request.Name,
		Stock: request.Stock,
	}
	item, err := ph.service.UpdateItem(uint(id), newItem)
	if err != nil {
		ph.logger.Info("error updating items")
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	// cetak pesan log dengan field email, item_id, item_name
	ph.logger.Info(
		"updating item successfull",
		zap.String("email", claims.Email),
		zap.Int("item_id", id),
		zap.String("item_name", newItem.Name),
	)

	return c.JSON(http.StatusOK, map[string]any{
		"item": item,
	})
}

func (ph *ItemHandler) DeleteItemHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*helper.JwtCustomClaims)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ph.logger.Info("failed to convert id")
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "invalid id, use integer id",
		})
	}

	item, err := ph.service.DeleteItem(uint(id))
	if err != nil {
		ph.logger.Info("error deleting items")
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	// cetak pesan log dengan field email, item_id
	ph.logger.Info(
		"deleting item successfull",
		zap.String("email", claims.Email),
		zap.Int("item_id", id),
	)

	return c.JSON(http.StatusOK, map[string]any{
		"item": item,
	})
}

func (ph *ItemHandler) AddItemHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*helper.JwtCustomClaims)

	productRequest := &addItemRequest{}
	err := c.Bind(productRequest)
	if err != nil {
		ph.logger.Info("failed to convert id")
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	newItem := domain.Item{
		Name:  productRequest.Name,
		Stock: productRequest.Stock,
	}

	item, err := ph.service.AddItem(newItem)
	if err != nil {
		ph.logger.Info("error adding items")
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	// cetak pesan log dengan field email, item_id, item_name
	ph.logger.Info(
		"successfully adding new item",
		zap.String("email", claims.Email),
		zap.String("item_name", item.Name),
		zap.Uint("item_id", item.ID),
	)

	return c.JSON(http.StatusOK, map[string]any{
		"message": "successfully adding new item",
		"item":    item,
	})
}
