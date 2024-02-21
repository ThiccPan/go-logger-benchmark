package main

import (
	"errors"
	"net/http"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Post struct {
	Id      uint
	Name    string
	content string
}

type PostRepo struct {
	sync.Mutex
	list    []Post
	counter uint
}

func NewPostRepo() *PostRepo {
	return &PostRepo{
		list:    []Post{},
		counter: 0,
	}
}

func (pr *PostRepo) addPost(post Post) error {
	pr.Lock()
	defer pr.Unlock()

	pr.list = append(pr.list, post)
	pr.counter++
	zerolog.TimeFieldFormat = zerolog.TimestampFieldName
	log.Info().Msg("post has been added")
	return nil
}

func (pr *PostRepo) getPost(id uint) (Post, error) {
	pr.Lock()
	defer pr.Unlock()

	var post Post
	found := false
	for _, v := range pr.list {
		if v.Id == id {
			post = v
			found = true
			break
		}
	}
	if !found {
		return Post{}, errors.New("post not found")
	}
	return post, nil
}

func (pr *PostRepo) updatePost(id uint, newPost *updatePostRequest) (Post, error) {
	pr.Lock()
	defer pr.Unlock()

	var post Post
	found := false
	for i, v := range pr.list {
		if v.Id == id {
			pr.list[i].Name = newPost.Name
			found = true
			post = pr.list[i]
			break
		}
	}
	if !found {
		return Post{}, errors.New("post not found")
	}
	log.Info().Msg("update post successfully")
	return post, nil
}

func (pr *PostRepo) deletePost(id uint) (Post, error) {
	pr.Lock()
	defer pr.Unlock()

	var post Post
	found := false
	for i, v := range pr.list {
		if v.Id == id {
			post = pr.list[i]
			pr.list = append(pr.list[0:i], pr.list[i+1:]...)
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
}

func NewPostHandler(repo *PostRepo) *PostHandler {
	return &PostHandler{Repo: repo}
}

type addPostRequest struct {
	Name string `json:"name"`
}

type updatePostRequest struct {
	Name string `json:"name"`
}

func (ph *PostHandler) getPostsHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"post": ph.Repo.list,
	})
}
func (ph *PostHandler) getPostHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
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

	post, err := ph.Repo.updatePost(uint(id), request)
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
	c.Bind(productRequest)

	// add item to productRepo
	newPost := Post{
		Id:   ph.Repo.counter,
		Name: productRequest.Name,
	}
	ph.Repo.addPost(newPost)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully adding new post",
		"post":    newPost,
	})
}

func main() {
	PostRepo := new(PostRepo)
	PostHandler := NewPostHandler(PostRepo)
	e := echo.New()

	e.GET("/", func(c echo.Context) error { return c.JSON(http.StatusOK, map[string]interface{}{"online": "online"}) })

	e.POST("/posts", PostHandler.addPostHandler)
	e.GET("/posts", PostHandler.getPostsHandler)
	e.GET("/posts/:id", PostHandler.getPostHandler)
	e.PUT("/posts/:id", PostHandler.updatePostHandler)
	e.DELETE("/posts/:id", PostHandler.deletePostHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
