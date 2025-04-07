package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"vote-broadcast-server/proto/auth"
	"vote-broadcast-server/services/auth/pkg/handlers"
	"vote-broadcast-server/services/auth/pkg/middleware"
	"vote-broadcast-server/services/auth/pkg/models"
	"vote-broadcast-server/services/auth/pkg/services"
)

type ServerManager struct {
	config   models.GRPCServer
	handlers *handlers.HandlersManager
	auth.UnimplementedAuthServiceServer
}

func NewServerManager(config models.GRPCServer, service *services.ServiceManager, jwtSecretKey models.JWTSecretKey) *ServerManager {
	return &ServerManager{
		handlers: handlers.NewHandlersManager(service, jwtSecretKey),
		config:   config,
	}
}

func (s *ServerManager) RegisterUser(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	return s.handlers.RegisterUser(ctx, req)
}

func (s *ServerManager) LoginUser(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	return s.handlers.LoginUser(ctx, req)
}

func (s *ServerManager) Start(service *services.ServiceManager, jwtSecretKey models.JWTSecretKey) {
	listener, err := net.Listen(s.config.Network, fmt.Sprintf(":%s", s.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.UnaryLoggingInterceptor),
	)
	auth.RegisterAuthServiceServer(server, NewServerManager(s.config, service, jwtSecretKey))

	log.Println("auth gRPC server is running on port: ", s.config.Port)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}
}
