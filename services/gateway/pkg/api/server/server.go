package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"vote-broadcast-server/services/gateway/pkg/api/handlers"
	"vote-broadcast-server/services/gateway/pkg/api/routes"
	"vote-broadcast-server/services/gateway/pkg/config"
)

type ServerManager struct {
	c        config.ConfigProvider
	handlers handlers.Handlers
	mux      *gin.Engine
}

func NewServerManager(c config.ConfigProvider, h handlers.Handlers) *ServerManager {
	return &ServerManager{
		c:        c,
		handlers: h,
		mux:      gin.Default(),
	}
}

func (s *ServerManager) Start() {
	cfg := s.c.GetConfig()

	// Initialize Routes
	routes.InitRoutes(s.mux, cfg.Services, s.handlers)

	// Start http server
	log.Println("Server is running on port: ", cfg.Port)
	if err := s.mux.Run(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		log.Fatalln(fmt.Errorf("failed to start server: %v", err))
	}
}
