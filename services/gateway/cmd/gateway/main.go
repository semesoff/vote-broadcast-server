package main

import (
	"vote-broadcast-server/services/gateway/pkg/api/handlers"
	"vote-broadcast-server/services/gateway/pkg/api/server"
	"vote-broadcast-server/services/gateway/pkg/config"
)

func main() {
	// Initialize ConfigManager
	var cfg config.ConfigProvider = config.NewConfigManager()

	// Initialize HandlerManager
	var handlerManager handlers.Handlers = handlers.NewHandlersManager(cfg)

	// Initialize ServerManager
	var serverManager = server.NewServerManager(cfg, handlerManager)
	serverManager.Start()
}
