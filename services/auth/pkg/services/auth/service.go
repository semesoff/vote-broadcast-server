package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"vote-broadcast-server/services/auth/pkg/models"
	"vote-broadcast-server/services/auth/pkg/services"
	"vote-broadcast-server/services/auth/pkg/services/jwt"
	"vote-broadcast-server/services/auth/pkg/utils/hash"
)

type AuthManager struct {
	*services.ServiceManager
	jwtSecretKey models.JWTSecretKey
}

type Auth interface {
	RegisterUser(user *models.User) (models.Token, error)
	LoginUser(user *models.User) (models.Token, error)
}

func NewAuthService(service *services.ServiceManager, jwtSecretKey models.JWTSecretKey) *AuthManager {
	return &AuthManager{
		service,
		jwtSecretKey,
	}
}

func (s *AuthManager) RegisterUser(user *models.User) (models.Token, error) {
	userWithID, ok, err := s.ServiceManager.Db.GetUser(*user)
	if err != nil {
		return models.Token{}, err
	}

	if ok {
		return models.Token{}, errors.New("user already exists")
	}

	if err := hash.HashPassword(user); err != nil {
		return models.Token{}, err
	}

	userWithID, err = s.ServiceManager.Db.AddUser(*user)
	if err != nil {
		return models.Token{}, errors.New("error adding user")
	}

	jwtManager := jwtModel.NewJWTManager(s.jwtSecretKey)
	token, err := jwtManager.GenerateToken(userWithID)

	if err != nil {
		return models.Token{}, errors.New("error generating token")
	}

	return models.Token{Token: token}, nil
}

func (s *AuthManager) LoginUser(user *models.User) (models.Token, error) {
	userWithPassword, ok, err := s.ServiceManager.Db.GetUserWithPassword(*user)
	if err != nil {
		return models.Token{}, err
	}

	if !ok {
		return models.Token{}, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword(userWithPassword.Password, []byte(user.Password)); err != nil {
		return models.Token{}, errors.New("invalid password")
	}

	jwtManager := jwtModel.NewJWTManager(s.jwtSecretKey)
	token, err := jwtManager.GenerateToken(
		models.UserWithID{
			ID:       userWithPassword.ID,
			Username: userWithPassword.Username,
		},
	)

	if err != nil {
		return models.Token{}, errors.New("error generating token")
	}

	return models.Token{Token: token}, nil
}
