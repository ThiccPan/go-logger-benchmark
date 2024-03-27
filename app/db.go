package app

import (
	"github.com/thiccpan/go-logger-benchmark/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("log-experiment.db"), &gorm.Config{})
	if err != nil {

		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&domain.Post{})
	db.AutoMigrate(&domain.User{})
	return db
}
