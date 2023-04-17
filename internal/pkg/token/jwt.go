package token

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	key                  = []byte(os.Getenv("JWT_SECRET_KEY"))
	accessTokenExpiredAt = time.Now().UTC().Add(time.Hour * 1)
	// refreshTokenExpiredAt = time.Now().UTC().Add(time.Hour * 720)
)

func GenerateAccessToken(userID int) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(accessTokenExpiredAt),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   fmt.Sprintf("%d", userID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
