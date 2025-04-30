package main

import (
	"gateway-service/pkg/api/handlers"
	"gateway-service/pkg/api/server"
	"gateway-service/pkg/config"
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
