package handlers

import (
	"context"
	"vote-broadcast-server/proto/auth"
	"vote-broadcast-server/services/auth/pkg/models"
	"vote-broadcast-server/services/auth/pkg/services"
	auth2 "vote-broadcast-server/services/auth/pkg/services/auth"
	"vote-broadcast-server/services/auth/pkg/utils/others"
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
