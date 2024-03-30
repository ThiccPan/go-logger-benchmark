package main

import (
	"github.com/labstack/echo/v4"
	"github.com/thiccpan/go-logger-benchmark/app"
	"github.com/thiccpan/go-logger-benchmark/handler"
	"github.com/thiccpan/go-logger-benchmark/logger"
	"github.com/thiccpan/go-logger-benchmark/repository"
	"github.com/thiccpan/go-logger-benchmark/service"
)

func main() {
	e := echo.New()

	// configure logger
	logger := logger.InitZap()
	// logger := logger.InitLogrusLogger()
	// logger := logger.InitZerolog()

	// Initialized db conn
	db := app.InitDB()

	// ItemRepo := repository.NewItemRepo(logger)
	ItemRepo := repository.NewSQLiteItemRepo(logger, db)
	AuthRepo := repository.NewSQLiteAuthRepo(logger, db)
	AuthService := service.NewAuthService(logger, AuthRepo)

	ItemHandler := handler.NewItemHandler(ItemRepo, logger)
	AuthHandler := handler.NewAuthHandler(logger, *AuthService)

	e.POST("/register", AuthHandler.RegisterHandler)
	e.GET("/login", AuthHandler.LoginHandler)

	e.POST("/items", ItemHandler.AddItemHandler)
	e.GET("/items", ItemHandler.GetItemsHandler)
	e.GET("/items/:id", ItemHandler.GetItemHandler)
	e.PUT("/items/:id", ItemHandler.UpdateItemHandler)
	e.DELETE("/items/:id", ItemHandler.DeleteItemHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
