package service

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

type jwtToken struct{}

func NewJWTToken() JWTService {
	return &jwtToken{}
}

func (service *jwtToken) GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID

	var JWT_SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(JWT_SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil

}
