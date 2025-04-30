package handlers

import (
	"auth-service/pkg/models"
	"auth-service/pkg/services"
	auth2 "auth-service/pkg/services/auth"
	"auth-service/pkg/utils/others"
	"auth-service/proto/auth"
	"context"
)

type HandlersManager struct {
	auth auth2.Auth
}

func NewHandlersManager(service *services.ServiceManager, JWTSecretKey []byte) *HandlersManager {
	return &HandlersManager{
		auth: auth2.NewAuthService(service, JWTSecretKey),
	}
}

type Handlers interface {
	RegisterUser(ctx *context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error)
	LoginUser(ctx *context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error)
}

func (h *HandlersManager) RegisterUser(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	user := models.User{
		Username: req.Username,
		Password: req.Password,
	}

	if err := others.CheckUserData(user); err != nil {
		return nil, err
	}

	token, err := h.auth.RegisterUser(&user)
	if err != nil {
		return nil, err
	}

	response := &auth.RegisterResponse{
		Success: true,
		Message: token.Token,
	}

	return response, nil
}

func (h *HandlersManager) LoginUser(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	user := models.User{
		Username: req.Username,
		Password: req.Password,
	}

	token, err := h.auth.LoginUser(&user)
	if err != nil {
		return nil, err
	}

	response := &auth.LoginResponse{
		Success: true,
		Token:   token.Token,
	}

	return response, nil
}
