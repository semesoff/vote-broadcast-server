package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
	"vote-broadcast-server/services/auth/pkg/models"
)

type JWTManager struct {
	secretKey []byte
}

type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewJWTManager() *JWTManager {
	// TODO: Change the method to get the secret key
	return &JWTManager{
		secretKey: []byte(os.Getenv("JWT_SECRET_KEY")),
	}
}

func (j *JWTManager) GenerateToken(user models.UserWithID) (string, error) {
	// TODO: Время истечения на будущее изменить
	claims := Claims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 2 * time.Duration(1))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
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
