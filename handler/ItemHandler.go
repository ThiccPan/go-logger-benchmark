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
)

type ItemHandler struct {
	Repo    repository.IItemRepo
	service service.IItemService
}

func NewItemHandler(repo repository.IItemRepo, service service.IItemService) *ItemHandler {
	return &ItemHandler{Repo: repo, service: service}
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
	_ = user.Claims.(*helper.JwtCustomClaims)

	items, err := ph.service.GetItems()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}
	// ph.logger.LogInfo("fetching all items successfull", map[string]any{
	// 	"email": claims.Email,
	// })

	return c.JSON(http.StatusOK, map[string]any{
		"item": items,
	})
}
func (ph *ItemHandler) GetItemHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	_ = user.Claims.(*helper.JwtCustomClaims)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// ph.logger.LogInfo("failed to convert id")
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "invalid id, use integer id",
		})
	}

	item, err := ph.service.GetItem(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	// ph.logger.LogInfo("getting item successfull", map[string]any{
	// 	"email":   claims.Email,
	// 	"item_id": id,
	// })

	return c.JSON(http.StatusOK, map[string]any{
		"item": item,
	})
}

func (ph *ItemHandler) UpdateItemHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	_ = user.Claims.(*helper.JwtCustomClaims)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// ph.logger.LogInfo("failed to convert id")
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

	newItem := domain.Item{
		Name:  request.Name,
		Stock: request.Stock,
	}
	item, err := ph.service.UpdateItem(uint(id), newItem)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	// ph.logger.LogInfo("updating item successfull", map[string]any{
	// 	"email":     claims.Email,
	// 	"item_id":   id,
	// 	"item_name": newItem.Name,
	// })

	return c.JSON(http.StatusOK, map[string]any{
		"item": item,
	})
}

func (ph *ItemHandler) DeleteItemHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	_ = user.Claims.(*helper.JwtCustomClaims)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// ph.logger.LogInfo("failed to convert id")
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "invalid id, use integer id",
		})
	}

	item, err := ph.service.DeleteItem(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	// ph.logger.LogInfo("updating item successfull", map[string]any{
	// 	"email":   claims.Email,
	// 	"item_id": id,
	// })

	return c.JSON(http.StatusOK, map[string]any{
		"item": item,
	})
}

func (ph *ItemHandler) AddItemHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	_ = user.Claims.(*helper.JwtCustomClaims)

	productRequest := &addItemRequest{}
	err := c.Bind(productRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	// add item to productRepo
	newItem := domain.Item{
		Name:  productRequest.Name,
		Stock: productRequest.Stock,
	}

	item, err := ph.service.AddItem(newItem)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	// ph.logger.LogInfo("successfully adding new item", map[string]any{
	// 	"email":     claims.Email,
	// 	"item_id":   item.ID,
	// 	"item_name": item.Name,
	// })

	return c.JSON(http.StatusOK, map[string]any{
		"message": "successfully adding new item",
		"item":    item,
	})
}
