package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/thiccpan/go-logger-benchmark/logger"
	"github.com/thiccpan/go-logger-benchmark/repository"
)

type PostHandler struct {
	Repo   *repository.PostRepo
	logger logger.Ilogger
}

func NewPostHandler(repo *repository.PostRepo, logger logger.Ilogger) *PostHandler {
	return &PostHandler{Repo: repo, logger: logger}
}

type addPostRequest struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type updatePostRequest struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (ph *PostHandler) GetPostsHandler(c echo.Context) error {
	ph.logger.LogInfo("fetching all posts")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"post": ph.Repo,
	})
}
func (ph *PostHandler) GetPostHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ph.logger.LogInfo("failed to convert id")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid id, use integer id",
		})
	}

	post, err := ph.Repo.GetPost(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"post": post,
	})
}

func (ph *PostHandler) UpdatePostHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ph.logger.LogInfo("failed to convert id")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid id, use integer id",
		})
	}

	request := &updatePostRequest{}
	err = c.Bind(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	newPost := &repository.Post{
		Name:    request.Name,
		Content: request.Content,
	}
	post, err := ph.Repo.UpdatePost(uint(id), newPost)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"post": post,
	})
}

func (ph *PostHandler) DeletePostHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ph.logger.LogInfo("failed to convert id")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid id, use integer id",
		})
	}

	post, err := ph.Repo.DeletePost(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"post": post,
	})
}

func (ph *PostHandler) AddPostHandler(c echo.Context) error {
	productRequest := &addPostRequest{}
	err := c.Bind(productRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	// add item to productRepo
	newPost := &repository.Post{
		Name:    productRequest.Name,
		Content: productRequest.Content,
	}
	ph.Repo.AddPost(newPost)

	ph.logger.LogInfo("successfully adding new post")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully adding new post",
		"post":    newPost,
	})
}
