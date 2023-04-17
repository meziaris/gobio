package controller

import (
	"gobio/internal/app/model"
	rsp "gobio/internal/app/response"
	"gobio/internal/app/service"
	"gobio/internal/pkg/token"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{UserService: *userService}
}

func (controller *UserController) Router(e *echo.Echo) {
	r := e.Group("/v1")
	r.POST("/user", controller.Create)
	r.POST("/login", controller.Login)
	r.POST("/avatar", controller.UploadAvatar, JWTMiddleware)
}

func (controller *UserController) Create(c echo.Context) error {
	var request = model.RegisterUserRequest{}

	err := c.Bind(&request)
	if err != nil {
		code := http.StatusUnprocessableEntity
		return c.JSON(code, rsp.APIResponse("Account register failed", code, "FAILED", err.Error()))
	}

	if err := c.Validate(request); err != nil {
		code := http.StatusUnprocessableEntity
		return c.JSON(code, rsp.APIResponse("Account register failed", code, "FAILED", err.Error()))
	}

	response, err := controller.UserService.Register(request)
	if err != nil {
		code := http.StatusBadRequest
		return c.JSON(code, rsp.APIResponse("Account register failed", code, "FAILED", err.Error()))
	}

	code := http.StatusOK
	return c.JSON(code, rsp.APIResponse("Your account has been created", code, "OK", response))
}

func (controller *UserController) Login(c echo.Context) error {
	var request = model.LoginUserRequest{}

	err := c.Bind(&request)
	if err != nil {
		code := http.StatusUnprocessableEntity
		return c.JSON(code, rsp.APIResponse("login failed", code, "FAILED", err.Error()))
	}

	if err := c.Validate(request); err != nil {
		code := http.StatusUnprocessableEntity
		return c.JSON(code, rsp.APIResponse("login failed", code, "FAILED", err.Error()))
	}

	user, err := controller.UserService.Login(request)
	if err != nil {
		code := http.StatusBadRequest
		return c.JSON(code, rsp.APIResponse("login failed", code, "FAILED", err.Error()))
	}

	token, err := token.GenerateAccessToken(user.ID)
	if err != nil {
		code := http.StatusBadRequest
		return c.JSON(code, rsp.APIResponse("login failed", code, "FAILED", err.Error()))
	}
	user.Token = token

	code := http.StatusOK
	return c.JSON(code, rsp.APIResponse("login success", code, "OK", user))
}

func (controller *UserController) UploadAvatar(c echo.Context) error {
	userID := c.Get("currentUserID").(int)
	request := model.UpdateAvatarRequest{}

	err := c.Bind(&request)
	if err != nil {
		code := http.StatusUnprocessableEntity
		return c.JSON(code, rsp.APIResponse("upload failed", code, "FAILED", err.Error()))
	}

	if err := c.Validate(request); err != nil {
		code := http.StatusUnprocessableEntity
		return c.JSON(code, rsp.APIResponse("upload failed", code, "FAILED", err.Error()))
	}

	response, err := controller.UserService.UploadAvatar(userID, request.AvatarUrl)
	if err != nil {
		code := http.StatusUnprocessableEntity
		return c.JSON(code, rsp.APIResponse("upload failed", code, "FAILED", err.Error()))
	}

	code := http.StatusOK
	return c.JSON(code, rsp.APIResponse("upload success", code, "OK", response))
}
