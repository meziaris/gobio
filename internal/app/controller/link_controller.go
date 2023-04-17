package controller

import (
	"gobio/internal/app/model"
	rsp "gobio/internal/app/response"
	"gobio/internal/app/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type LinkController struct {
	LinkService service.LinkService
}

func NewLinkController(linkService *service.LinkService) *LinkController {
	return &LinkController{
		LinkService: *linkService,
	}
}

func (controller *LinkController) Router(e *echo.Echo) {
	r := e.Group("/v1")
	r.POST("/link", controller.Add, JWTMiddleware)
	r.DELETE("/link/:id", controller.Delete, JWTMiddleware)
	r.PATCH("/link/:id", controller.Update, JWTMiddleware)

	r.GET("/:username", controller.UserLink)
}

func (controller *LinkController) Add(c echo.Context) error {
	var request = model.AddLinkRequest{}
	userID := c.Get("currentUserID").(int)

	err := c.Bind(&request)
	if err != nil {
		code := http.StatusUnprocessableEntity
		return c.JSON(code, rsp.APIResponse("Add link failed", code, "FAILED", err.Error()))
	}

	if err := c.Validate(request); err != nil {
		code := http.StatusUnprocessableEntity
		return c.JSON(code, rsp.APIResponse("Add link failed", code, "FAILED", err.Error()))
	}

	response, err := controller.LinkService.AddLink(request, userID)
	if err != nil {
		code := http.StatusUnprocessableEntity
		return c.JSON(code, rsp.APIResponse("Add link failed", code, "FAILED", err.Error()))
	}

	code := http.StatusOK
	return c.JSON(code, rsp.APIResponse("Add link success", code, "OK", response))
}

func (controller *LinkController) Update(c echo.Context) error {
	userID := c.Get("currentUserID").(int)
	var request = model.UpdateLinkRequest{}
	linkID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		code := http.StatusUnprocessableEntity
		return c.JSON(code, rsp.APIResponse("Update link failed", code, "FAILED", err.Error()))
	}

	err = c.Bind(&request)
	if err != nil {
		code := http.StatusUnprocessableEntity
		return c.JSON(code, rsp.APIResponse("Update link failed", code, "FAILED", err.Error()))
	}

	if err := c.Validate(request); err != nil {
		code := http.StatusUnprocessableEntity
		return c.JSON(code, rsp.APIResponse("Update link failed", code, "FAILED", err.Error()))
	}

	response, err := controller.LinkService.UpdateLink(request, linkID, userID)
	if err != nil {
		code := http.StatusUnprocessableEntity
		return c.JSON(code, rsp.APIResponse("Update link failed", code, "FAILED", err.Error()))
	}

	code := http.StatusOK
	return c.JSON(code, rsp.APIResponse("Your link has been updated", code, "OK", response))
}

func (controller *LinkController) Delete(c echo.Context) error {
	userID := c.Get("currentUserID").(int)
	linkID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		code := http.StatusUnprocessableEntity
		return c.JSON(code, rsp.APIResponse("Link not found", code, "FAILED", err.Error()))
	}

	err = controller.LinkService.DeleteLink(linkID, userID)
	if err != nil {
		code := http.StatusUnprocessableEntity
		return c.JSON(code, rsp.APIResponse("Link not found", code, "FAILED", err.Error()))
	}

	code := http.StatusOK
	return c.JSON(code, rsp.APIResponse("Delete link success", code, "OK", "Link has been deleted"))
}

func (controller *LinkController) UserLink(c echo.Context) error {
	username := c.Param("username")

	response, err := controller.LinkService.List(username)
	if err != nil {
		code := http.StatusNotFound
		return c.JSON(code, rsp.APIResponse("User not found", code, "FAILED", err.Error()))
	}

	code := http.StatusOK
	return c.JSON(code, rsp.APIResponse("Get link success", code, "OK", response))
}
