package server

import (
	"vote-broadcast-server/services/vote/pkg/config"
	"vote-broadcast-server/services/vote/pkg/services"
	"vote-broadcast-server/services/vote/pkg/services/server"
)

type ServerManager struct {
	cfg    config.ConfigProvider
	server *server.ServerManager
}

func NewServerManager() *ServerManager {
	return &ServerManager{}
}

func (s *ServerManager) Start() {
	// Initialize vote-config
	var voteConfig config.ConfigProvider = config.NewConfigManager()
	s.cfg = voteConfig

	// Initialize ServiceManager
	service := services.NewServiceManager(s.cfg.GetConfig().Database)

	// Initialize GRPC Server
	s.server = server.NewServerManager(s.cfg.GetConfig().GRPCServer, service)

	// Start GRPCServer
	s.server.Start()
}
