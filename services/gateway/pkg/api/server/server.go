package server

import (
	"fmt"
	"gateway-service/pkg/api/handlers"
	"gateway-service/pkg/api/middleware"
	"gateway-service/pkg/api/routes"
	"gateway-service/pkg/config"
	"github.com/gin-gonic/gin"
	"log"
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
		mux:      gin.New(),
	}
}

func (s *ServerManager) Start() {
	cfg := s.c.GetConfig()

	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	s.mux.Use(middleware.LoggingMiddleware)

	// Initialize Routes
	routes.InitRoutes(s.mux, cfg.Services, s.handlers)

	// Start http server
	log.Println("Server is running on port: ", cfg.Port)
	if err := s.mux.Run(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		log.Fatalln(fmt.Errorf("failed to start server: %v", err))
	}
}
