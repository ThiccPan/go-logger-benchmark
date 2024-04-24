package handler

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/thiccpan/go-logger-benchmark/domain"
	"github.com/thiccpan/go-logger-benchmark/helper"
	"github.com/thiccpan/go-logger-benchmark/logger"
	"github.com/thiccpan/go-logger-benchmark/repository"
)

type ItemHandler struct {
	Repo   repository.IItemRepo
	logger logger.Ilogger
}

func NewItemHandler(repo repository.IItemRepo, logger logger.Ilogger) *ItemHandler {
	return &ItemHandler{Repo: repo, logger: logger}
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

	items, err := ph.Repo.GetItems()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}
	ph.logger.LogInfo("fetching all items successfull", map[string]any{
		"email": claims.Email,
	})

	return c.JSON(http.StatusOK, map[string]any{
		"item": items,
	})
}
func (ph *ItemHandler) GetItemHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*helper.JwtCustomClaims)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ph.logger.LogInfo("failed to convert id")
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "invalid id, use integer id",
		})
	}

	item, err := ph.Repo.GetItem(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	ph.logger.LogInfo("getting item successfull", map[string]any{
		"email": claims.Email,
		"item_id": id,
	})

	return c.JSON(http.StatusOK, map[string]any{
		"item": item,
	})
}

func (ph *ItemHandler) UpdateItemHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*helper.JwtCustomClaims)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ph.logger.LogInfo("failed to convert id")
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "invalid id, use integer id",
		})
	}

	request := &updateItemRequest{}
	err = c.Bind(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	newItem := &domain.Item{
		Name:  request.Name,
		Stock: request.Stock,
	}
	item, err := ph.Repo.UpdateItem(uint(id), newItem)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	ph.logger.LogInfo("updating item successfull", map[string]any{
		"email": claims.Email,
		"item_id": id,
	})

	return c.JSON(http.StatusOK, map[string]any{
		"item": item,
	})
}

func (ph *ItemHandler) DeleteItemHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*helper.JwtCustomClaims)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ph.logger.LogInfo("failed to convert id")
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "invalid id, use integer id",
		})
	}

	item, err := ph.Repo.DeleteItem(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	ph.logger.LogInfo("updating item successfull", map[string]any{
		"email": claims.Email,
		"item_id": id,
	})

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
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	// add item to productRepo
	newItem := &domain.Item{
		Name:  productRequest.Name,
		Stock: productRequest.Stock,
	}

	if err = ph.Repo.AddItem(newItem); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	ph.logger.LogInfo("successfully adding new item", map[string]any{
		"email": claims.Email,
	})

	return c.JSON(http.StatusOK, map[string]any{
		"message": "successfully adding new item",
		"item":    newItem,
	})
}
