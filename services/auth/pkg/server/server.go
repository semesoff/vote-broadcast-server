package server

import (
	"auth-service/pkg/config"
	"auth-service/pkg/services"
	"auth-service/pkg/services/server"
)

type ServerManager struct {
	cfg    config.ConfigProvider
	server *server.ServerManager
}

func NewServerManager() *ServerManager {
	return &ServerManager{}
}

func (s *ServerManager) Start() {
	// Initialize auth-config
	var authConfig config.ConfigProvider = config.NewConfigManager()
	s.cfg = authConfig

	// Initialize ServiceManager
	service := services.NewServiceManager(authConfig.GetConfig().Database)

	// Initialize GRPCServer
	s.server = server.NewServerManager(s.cfg.GetConfig().GRPCServer, service, s.cfg.GetConfig().JWTSecretKey)

	// Start GRPCServer
	s.server.Start(service, s.cfg.GetConfig().JWTSecretKey)
}
