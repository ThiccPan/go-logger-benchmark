package main

import (
	"errors"
	"net/http"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/thiccpan/go-logger-benchmark/logger"
)

type Post struct {
	Id      uint
	Name    string
	Content string
}

type PostRepo struct {
	sync.Mutex
	store   map[uint]*Post
	logger  logger.Ilogger
	counter uint
}

func NewPostRepo(logger logger.Ilogger) *PostRepo {
	repo := PostRepo{
		store:   make(map[uint]*Post),
		logger:  logger,
		counter: 0,
	}
	return &repo
}

func (pr *PostRepo) addPost(post *Post) error {
	pr.Lock()
	defer pr.Unlock()

	pr.logger.LogInfo(post.Name)
	pr.store[pr.counter] = post
	pr.counter++
	pr.logger.LogInfo("post has been added")
	return nil
}

func (pr *PostRepo) getPost(id uint) (Post, error) {
	pr.Lock()
	defer pr.Unlock()

	post, found := pr.store[id]
	pr.logger.LogInfo("fetching post successfully")
	if !found {
		return Post{}, errors.New("post not found")
	}
	return *post, nil
}

func (pr *PostRepo) updatePost(id uint, newPost *Post) (Post, error) {
	pr.Lock()
	defer pr.Unlock()

	post, found := pr.store[id]
	pr.logger.LogInfo("fetching post successfully")
	if !found {
		return Post{}, errors.New("post not found")
	}

	if newPost.Name == "" {
		newPost.Name = post.Name
	}
	if newPost.Content == "" {
		newPost.Content = post.Content
	}
	pr.store[id] = newPost
	pr.logger.LogInfo("update post successfully")
	return *pr.store[id], nil
}

func (pr *PostRepo) deletePost(id uint) (Post, error) {
	pr.Lock()
	defer pr.Unlock()

	var post Post
	found := false
	for i, v := range pr.store {
		if v.Id == id {
			post = *pr.store[i]
			delete(pr.store, i)
			found = true
			break
		}
	}
	if !found {
		return Post{}, errors.New("post not found")
	}
	pr.logger.LogInfo("delete post successfully")
	return post, nil
}

type PostHandler struct {
	Repo   *PostRepo
	logger logger.Ilogger
}

func NewPostHandler(repo *PostRepo, logger logger.Ilogger) *PostHandler {
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

func (ph *PostHandler) getPostsHandler(c echo.Context) error {
	ph.logger.LogInfo("fetching all posts")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"post": ph.Repo.store,
	})
}
func (ph *PostHandler) getPostHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ph.logger.LogInfo("failed to convert id")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid id, use integer id",
		})
	}

	post, err := ph.Repo.getPost(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"post": post,
	})
}

func (ph *PostHandler) updatePostHandler(c echo.Context) error {
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

	newPost := &Post{
		Id:      ph.Repo.counter,
		Name:    request.Name,
		Content: request.Content,
	}
	post, err := ph.Repo.updatePost(uint(id), newPost)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"post": post,
	})
}

func (ph *PostHandler) deletePostHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ph.logger.LogInfo("failed to convert id")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid id, use integer id",
		})
	}

	post, err := ph.Repo.deletePost(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"post": post,
	})
}

func (ph *PostHandler) addPostHandler(c echo.Context) error {
	productRequest := &addPostRequest{}
	err := c.Bind(productRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	// add item to productRepo
	newPost := &Post{
		Id:      ph.Repo.counter,
		Name:    productRequest.Name,
		Content: productRequest.Content,
	}
	ph.Repo.addPost(newPost)

	ph.logger.LogInfo("successfully adding new post")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully adding new post",
		"post":    newPost,
	})
}

func main() {
	e := echo.New()

	// configure logger

	zapLogger := logger.InitZap()

	PostRepo := NewPostRepo(zapLogger)
	PostHandler := NewPostHandler(PostRepo, zapLogger)

	e.POST("/posts", PostHandler.addPostHandler)
	e.GET("/posts", PostHandler.getPostsHandler)
	e.GET("/posts/:id", PostHandler.getPostHandler)
	e.PUT("/posts/:id", PostHandler.updatePostHandler)
	e.DELETE("/posts/:id", PostHandler.deletePostHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
