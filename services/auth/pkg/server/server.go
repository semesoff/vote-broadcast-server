package server

import (
	"vote-broadcast-server/services/auth/pkg/config"
	"vote-broadcast-server/services/auth/pkg/services"
	"vote-broadcast-server/services/auth/pkg/services/server"
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
	s.server = server.NewServerManager(s.cfg.GetConfig().GRPCServer, service)

	// Start GRPCServer
	s.server.Start(service)
}
