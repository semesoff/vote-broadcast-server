package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"vote-broadcast-server/services/websocket/pkg/config"
	"vote-broadcast-server/services/websocket/pkg/models"
	"vote-broadcast-server/services/websocket/pkg/services/grpc_server"
	"vote-broadcast-server/services/websocket/pkg/services/websocket_server"
)

type ServerManager struct {
	config          config.ConfigProvider
	grpcServer      grpc_server.Server
	websocketServer websocket_server.Server
	waitGroup       *sync.WaitGroup
}

func NewServerManager() *ServerManager {
	return &ServerManager{
		config:    config.NewConfigManager(),
		waitGroup: &sync.WaitGroup{},
	}
}

func (s *ServerManager) Start() {
	// Initialize data channels for communication between gRPC and WebSocket servers
	dataChannels := make(models.DataChannels)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	s.waitGroup.Add(2)

	// Initialize and start gRPC server
	s.grpcServer = grpc_server.NewServerManager(s.config.GetConfig().GRPCServer, dataChannels)
	go s.grpcServer.Start(s.waitGroup, ctx)

	// Initialize and start websocket server
	s.websocketServer = websocket_server.NewServerManager(s.config.GetConfig().WebSocketServer, dataChannels)
	go s.websocketServer.Start(s.waitGroup, ctx)

	// Wait for interrupt signal to gracefully shut down the servers
	go func() {
		<-interrupt
		cancel()
		log.Println("Received interrupt signal, shutting down servers...")
	}()

	log.Println("Server started")
	s.waitGroup.Wait()
	close(dataChannels)
	log.Println("Server shut down")
}
