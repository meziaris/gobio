package main

import (
	"gobio/controller"
	"gobio/entity"
	"gobio/repository"
	"gobio/service"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=gobio port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&entity.User{})

	// Setup Repository
	userRepository := repository.NewUserRepository(db)

	// Setup Service
	userService := service.NewUserService(&userRepository)

	// Setup Controller
	userController := controller.NewUserController(&userService)

	// Setup Echo
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${protocol} - [${method}], ${host} ${uri}, status=${status}, latency=${latency_human} ${error}\n",
	}))

	// Setup Router
	userController.Router(e)
	e.Logger.Fatal(e.Start(":8080"))
}
