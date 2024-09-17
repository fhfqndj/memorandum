package service

import (
    "github.com/dgrijalva/jwt-go"
	"memorandum-backend/internal/entities"
    "time"
)

type JWTService struct {
    secretKey string
}

func NewJWTService(secretKey string) entities.TokenService {
    return &JWTService{secretKey: secretKey}
}

func (s *JWTService) GenerateToken(user *entities.User) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })
    return token.SignedString([]byte(s.secretKey))
}
