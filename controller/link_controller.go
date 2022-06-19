package controller

import (
	"fmt"
	"gobio/helper"
	"gobio/model"
	"gobio/service"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(os.Getenv("JWT_SECRET_KEY")),
	}))
	r.POST("/link", controller.Add)
}

func (controller *LinkController) Add(c echo.Context) error {
	var request = model.AddLinkRequest{}

	header := c.Request().Header
	authv := header.Get("Authorization")

	// Get bearer token
	if !strings.HasPrefix(strings.ToLower(authv), "bearer") {
		fmt.Println("no token")
	}

	values := strings.Split(authv, " ")
	if len(values) < 2 {
		fmt.Println("terjadi kesalahan")
	}

	tokenString := values[1]
	var mySigningKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return mySigningKey, nil
	})

	var userID int
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID = int(claims["user_id"].(float64))
	}

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
