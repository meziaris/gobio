package service

type JWTService interface {
	GenerateToken(userID int) (string, error)
}
