package controller

import (
	"errors"
	"gobio/helper"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")

		tokenHeader := strings.Split(header, " ")
		if len(tokenHeader) < 2 {
			code := http.StatusUnauthorized
			return c.JSON(code, helper.APIResponse("Unauthorization", code, "FAILED", errors.New("you are not allowed to access this page")))
		}
		tokenString := tokenHeader[1]

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("invalid token")
			}

			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			code := http.StatusUnauthorized
			return c.JSON(code, helper.APIResponse("Unauthorization", code, "FAILED", errors.New("you are not allowed to access this page")))
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			code := http.StatusUnauthorized
			return c.JSON(code, helper.APIResponse("Unauthorization", code, "FAILED", errors.New("you are not allowed to access this page")))
		}

		userID := int(claims["user_id"].(float64))
		c.Set("currentUserID", userID)

		return next(c)
	}
}
