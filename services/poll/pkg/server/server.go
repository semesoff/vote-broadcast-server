package server

import (
	"vote-broadcast-server/services/poll/pkg/config"
	"vote-broadcast-server/services/poll/pkg/services"
	"vote-broadcast-server/services/poll/pkg/services/server"
)

type ServerManager struct {
	cfg    config.ConfigProvider
	server *server.ServerManager
}

func NewServerManager() *ServerManager {
	return &ServerManager{}
}

func (s *ServerManager) Start() {
	// Initialize poll-config
	var pollConfig config.ConfigProvider = config.NewConfigManager()
	s.cfg = pollConfig

	// Initialize ServiceManager
	service := services.NewServiceManager(s.cfg.GetConfig().Database)

	// Initialize GRPCServer
	s.server = server.NewServerManager(s.cfg.GetConfig().GRPCServer, service)

	// Start GRPCServer
	s.server.Start()
}
