package controller

import (
	"gobio/helper"
	"gobio/model"
	"gobio/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(service *service.UserService) UserController {
	return UserController{
		UserService: *service,
	}
}

func (controller *UserController) Router(e *echo.Echo) {
	e.POST("/v1/user", controller.Create)
}

func (controller *UserController) Create(c echo.Context) error {
	var request = model.RegisterUserRequest{}

	err := c.Bind(&request)
	if err != nil {
		code := http.StatusUnprocessableEntity
		return c.JSON(code, helper.APIResponse("Account register failed", code, "FAILED", err.Error()))
	}

	response, err := controller.UserService.Register(request)
	if err != nil {
		code := http.StatusBadRequest
		return c.JSON(code, helper.APIResponse("Account register failed", code, "FAILED", err.Error()))
	}

	code := http.StatusOK
	return c.JSON(code, helper.APIResponse("Your account has been created", code, "OK", response))
}
