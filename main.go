package main

import (
	"errors"
	"fmt"
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

	var post Post
	found := false
	for _, v := range pr.store {
		if v.Id == id {
			post = *v
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
	for i, v := range pr.store {
		if v.Id == id {
			pr.store[i].Name = newPost.Name
			found = true
			post = *pr.store[i]
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
		"post": ph.Repo.store,
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
	err := c.Bind(productRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	// add item to productRepo
	newPost := &Post{
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
	zerolog.TimeFieldFormat = zerolog.TimestampFieldName

	PostRepo := NewPostRepo()
	fmt.Println(PostRepo.store)
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
