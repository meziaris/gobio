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
	JWTService  service.JWTService
}

func NewUserController(userService *service.UserService, jwtService *service.JWTService) UserController {
	return UserController{
		UserService: *userService,
		JWTService:  *jwtService,
	}
}

func (controller *UserController) Router(e *echo.Echo) {
	r := e.Group("/v1")
	r.POST("/user", controller.Create)
	r.POST("/login", controller.Login)
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

func (controller *UserController) Login(c echo.Context) error {
	var request = model.LoginUserRequest{}

	err := c.Bind(&request)
	if err != nil {
		code := http.StatusUnprocessableEntity
		return c.JSON(code, helper.APIResponse("login failed", code, "FAILED", err.Error()))
	}

	user, err := controller.UserService.Login(request, "token")
	if err != nil {
		code := http.StatusBadRequest
		return c.JSON(code, helper.APIResponse("login failed", code, "FAILED", err.Error()))
	}

	token, err := controller.JWTService.GenerateToken(user.ID)
	if err != nil {
		code := http.StatusBadRequest
		return c.JSON(code, helper.APIResponse("login failed", code, "FAILED", err.Error()))
	}

	response := helper.LoginResponse(user, token)

	code := http.StatusOK
	return c.JSON(code, helper.APIResponse("login success", code, "OK", response))
}
