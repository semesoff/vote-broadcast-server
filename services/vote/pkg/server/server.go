package server

import (
	"vote-service/pkg/config"
	"vote-service/pkg/handlers"
	"vote-service/pkg/services"
	"vote-service/pkg/services/notification_service"
	"vote-service/pkg/services/server"
	"vote-service/pkg/services/websocket_notifier"
)

type ServerManager struct {
	cfg                 config.ConfigProvider
	server              *server.ServerManager
	websocketNotifier   websocket_notifier.WebSocketNotifier
	notificationService notification_service.NotificationService
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

	// Initialize and start WebSocketNotifierService
	s.websocketNotifier = websocket_notifier.NewWebSocketNotifierService(s.cfg.GetConfig().WebSocketGRPCServer)
	go s.websocketNotifier.Start()

	// Initialize NotificationService
	s.notificationService = notification_service.NewNotificationServiceManager(s.websocketNotifier.GetDataChannel(), service)

	// Initialize HandlersManager
	handlersManager := handlers.NewHandlersManager(service, s.notificationService)

	// Initialize GRPC Server
	s.server = server.NewServerManager(s.cfg.GetConfig().GRPCServer, handlersManager)

	// Start GRPCServer
	s.server.Start()
}
