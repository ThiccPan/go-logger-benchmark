package main

import (
	"github.com/labstack/echo/v4"
	"github.com/thiccpan/go-logger-benchmark/handler"
	"github.com/thiccpan/go-logger-benchmark/logger"
	"github.com/thiccpan/go-logger-benchmark/repository"
)

func main() {
	e := echo.New()

	// configure logger

	logger := logger.InitZap()
	// logger := logger.InitZerolog()

	PostRepo := repository.NewPostRepo(logger)
	PostHandler := handler.NewPostHandler(PostRepo, logger)

	e.POST("/posts", PostHandler.AddPostHandler)
	e.GET("/posts", PostHandler.GetPostsHandler)
	e.GET("/posts/:id", PostHandler.GetPostHandler)
	e.PUT("/posts/:id", PostHandler.UpdatePostHandler)
	e.DELETE("/posts/:id", PostHandler.DeletePostHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
