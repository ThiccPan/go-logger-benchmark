package main

import (
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/thiccpan/go-logger-benchmark/app"
	"github.com/thiccpan/go-logger-benchmark/handler"
	"github.com/thiccpan/go-logger-benchmark/helper"
	"github.com/thiccpan/go-logger-benchmark/logger"
	"github.com/thiccpan/go-logger-benchmark/repository"
	"github.com/thiccpan/go-logger-benchmark/service"
)

func main() {
	e := echo.New()

	// configure logger
	logger := logger.InitZap()
	// logger := logger.InitLogrusLogger()

	// Initialized db conn
	db := app.InitDB()

	// init jwt helper
	jwtGen := helper.NewJWTGen("secret123")

	// ItemRepo := repository.NewItemRepo(logger)
	ItemRepo := repository.NewSQLiteItemRepo(logger, db)
	ItemService := service.NewItemService(ItemRepo)
	ItemHandler := handler.NewItemHandler(ItemRepo, ItemService, logger)

	AuthRepo := repository.NewSQLiteAuthRepo(logger, db)
	AuthService := service.NewAuthService(logger, AuthRepo)
	AuthHandler := handler.NewAuthHandler(logger, AuthService, *jwtGen)

	e.POST("/login", AuthHandler.LoginHandler)
	e.POST("/register", AuthHandler.RegisterHandler)

	r := e.Group("/items")

	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helper.JwtCustomClaims)
		},
		SigningKey: []byte("secret123"),
	}
	r.Use(echojwt.WithConfig(config))

	r.POST("", ItemHandler.AddItemHandler)
	r.GET("", ItemHandler.GetItemsHandler)
	r.GET("/:id", ItemHandler.GetItemHandler)
	r.PUT("/:id", ItemHandler.UpdateItemHandler)
	r.DELETE("/:id", ItemHandler.DeleteItemHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
