package grpc_server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"websocket-service/pkg/handlers"
	"websocket-service/pkg/middleware"
	"websocket-service/pkg/models"
	"websocket-service/proto/websocket"
)

type ServerManager struct {
	config   models.GRPCServer
	handlers handlers.Handlers
	websocket.WebSocketServiceServer
}

func NewServerManager(config models.GRPCServer, dataChannels models.DataChannels) *ServerManager {
	return &ServerManager{
		config:   config,
		handlers: handlers.NewHandlersManager(dataChannels),
	}
}

type Server interface {
	Start(waitGroup *sync.WaitGroup, ctx context.Context)
}

func (s *ServerManager) GetPolls(ctx context.Context, req *websocket.PollsRequest) (*websocket.PollsResponse, error) {
	return s.handlers.GetPolls(ctx, req)
}

func (s *ServerManager) GetVotes(ctx context.Context, req *websocket.VotesRequest) (*websocket.VotesResponse, error) {
	return s.handlers.GetVotes(ctx, req)
}

func (s *ServerManager) Start(waitGroup *sync.WaitGroup, ctx context.Context) {
	listener, err := net.Listen(s.config.Network, fmt.Sprintf(":%s", s.config.Port))
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.UnaryLoggingInterceptor),
	)

	websocket.RegisterWebSocketServiceServer(server, s)

	go func() {
		<-ctx.Done()
		server.GracefulStop()
		waitGroup.Done()
		log.Println("GRPC Server is stopped.")
	}()

	log.Println("websocket gGRPC server is running on port: ", s.config.Port)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
		return
	}
}
