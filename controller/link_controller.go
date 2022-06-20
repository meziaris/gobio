package controller

import (
	"gobio/helper"
	"gobio/model"
	"gobio/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LinkController struct {
	LinkService service.LinkService
	JWTService  service.JWTService
}

func NewLinkController(linkService *service.LinkService) LinkController {
	return LinkController{
		LinkService: *linkService,
	}
}

func (controller *LinkController) Router(e *echo.Echo) {
	r := e.Group("/v1")
	r.POST("/link", controller.Add, JWTMiddleware)
	r.GET("/:username", controller.UserLink)
}

func (controller *LinkController) Add(c echo.Context) error {
	var request = model.AddLinkRequest{}
	userID := c.Get("currentUserID").(int)

	err := c.Bind(&request)
	if err != nil {
		code := http.StatusUnprocessableEntity
		return c.JSON(code, helper.APIResponse("Add link failed", code, "FAILED", err.Error()))
	}

	response, err := controller.LinkService.AddLink(request, userID)
	if err != nil {
		code := http.StatusUnprocessableEntity
		return c.JSON(code, helper.APIResponse("Add link failed", code, "FAILED", err.Error()))
	}

	code := http.StatusOK
	return c.JSON(code, helper.APIResponse("Add link success", code, "OK", response))
}

func (controller *LinkController) UserLink(c echo.Context) error {
	username := c.Param("username")

	response, err := controller.LinkService.List(username)
	if err != nil {
		code := http.StatusNotFound
		return c.JSON(code, helper.APIResponse("User not found", code, "FAILED", err.Error()))
	}

	code := http.StatusOK
	return c.JSON(code, helper.APIResponse("Get link success", code, "OK", response))
}
