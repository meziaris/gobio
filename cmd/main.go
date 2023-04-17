package main

import (
	"gobio/internal/app/controller"
	"gobio/internal/app/repository"
	"gobio/internal/app/service"
	"gobio/internal/pkg/config"
	"gobio/internal/pkg/validation"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	configuration := config.New()
	db := config.NewDatabase(configuration)

	// Setup Repository
	userRepository := repository.NewUserRepository(db)
	linkRepository := repository.NewLinkRepository(db)

	// Setup Service
	userService := service.NewUserService(&userRepository)
	linkService := service.NewLinkService(&linkRepository, &userRepository)

	// Setup Controller
	userController := controller.NewUserController(&userService)
	linkController := controller.NewLinkController(&linkService)

	// Setup Echo
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${protocol} - [${method}], ${host} ${uri}, status=${status}, latency=${latency_human} ${error}\n",
	}))
	e.Validator = &validation.CustomValidator{Validator: validator.New()}

	// Setup Router
	userController.Router(e)
	linkController.Router(e)
	e.Logger.Fatal(e.Start(":8080"))
}
