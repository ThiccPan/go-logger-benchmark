package repository

import (
	"errors"
	"sync"

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

func (pr *PostRepo) AddPost(post *Post) error {
	pr.Lock()
	defer pr.Unlock()

	pr.logger.LogInfo(post.Name)
	post.Id = pr.counter
	pr.store[pr.counter] = post
	pr.counter++
	pr.logger.LogInfo("post has been added")
	return nil
}

func (pr *PostRepo) GetPosts() ([]Post, error) {
	pr.Lock()
	defer pr.Unlock()
	var posts []Post
	for _, post := range pr.store {
		posts = append(posts, *post)
	}
	return posts, nil
}

func (pr *PostRepo) GetPost(id uint) (Post, error) {
	pr.Lock()
	defer pr.Unlock()

	post, found := pr.store[id]
	pr.logger.LogInfo("fetching post successfully")
	if !found {
		return Post{}, errors.New("post not found")
	}
	return *post, nil
}

func (pr *PostRepo) UpdatePost(id uint, newPost *Post) (Post, error) {
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

func (pr *PostRepo) DeletePost(id uint) (Post, error) {
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
