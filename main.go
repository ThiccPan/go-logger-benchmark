package main

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Post struct {
	Id      uint
	Name    string
	Content string
}

type PostRepo struct {
	sync.Mutex
	store   map[uint]*Post
	counter uint
}

func NewPostRepo() *PostRepo {
	repo := PostRepo{
		counter: 0,
	}
	repo.store = make(map[uint]*Post)
	return &repo
}

func (pr *PostRepo) addPost(post *Post) error {
	pr.Lock()
	defer pr.Unlock()

	log.Info().Msg(post.Name)
	pr.store[pr.counter] = post
	pr.counter++
	log.Info().Msg("post has been added")
	return nil
}

func (pr *PostRepo) getPost(id uint) (Post, error) {
	pr.Lock()
	defer pr.Unlock()

	post, found := pr.store[id]
	log.Info().Msg("fetching post successfully")
	if !found {
		return Post{}, errors.New("post not found")
	}
	return *post, nil
}

func (pr *PostRepo) updatePost(id uint, newPost *Post) (Post, error) {
	pr.Lock()
	defer pr.Unlock()

	post, found := pr.store[id]
	log.Info().Msg("fetching post successfully")
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
	log.Info().Msg("update post successfully")
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
	log.Info().Msg("delete post successfully")
	return post, nil
}

type PostHandler struct {
	Repo *PostRepo
	logger Ilogger
}

func NewPostHandler(repo *PostRepo, logger Ilogger) *PostHandler {
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
	log.Info().Msg("fetching all posts")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"post": ph.Repo.store,
	})
}
func (ph *PostHandler) getPostHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Warn().Msg("failed to convert id")
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
		log.Warn().Msg("failed to convert id")
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
		log.Warn().Msg("failed to convert id")
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

	ph.logger.logInfo("successfully adding new post")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully adding new post",
		"post":    newPost,
	})
}

func main() {
	e := echo.New()

	// configure logger
	logFile, err := os.OpenFile("./log-history.txt", os.O_RDWR, 0644)
	if err != nil {
		panic(1)
	}

	zerolog.TimeFieldFormat = zerolog.TimestampFieldName
	log.Logger = log.Output(logFile)
	log.Logger = log.With().Caller().Logger()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))

	PostRepo := NewPostRepo()
	zapLogger := InitZap()
	PostHandler := NewPostHandler(PostRepo, zapLogger)

	e.POST("/posts", PostHandler.addPostHandler)
	e.GET("/posts", PostHandler.getPostsHandler)
	e.GET("/posts/:id", PostHandler.getPostHandler)
	e.PUT("/posts/:id", PostHandler.updatePostHandler)
	e.DELETE("/posts/:id", PostHandler.deletePostHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
