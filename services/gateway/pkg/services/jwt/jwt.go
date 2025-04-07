package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secretKey []byte
}

type JWTProvider interface {
	VerifyToken(tokenString string) (*Claims, error)
}

type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewJWTManager(secretKey []byte) *JWTManager {
	return &JWTManager{
		secretKey: secretKey,
	}
}

func (j *JWTManager) VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return j.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
