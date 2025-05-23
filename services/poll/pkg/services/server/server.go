package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"poll-service/pkg/handlers"
	"poll-service/pkg/middleware"
	"poll-service/pkg/models"
	"poll-service/proto/poll"
)

type ServerManager struct {
	config   models.GRPCServer
	handlers *handlers.HandlersManager
	poll.PollServiceServer
}

func (s *ServerManager) GetPolls(ctx context.Context, req *poll.GetPollsRequest) (*poll.GetPollsResponse, error) {
	return s.handlers.GetPolls(ctx, req)
}

func (s *ServerManager) CreatePoll(ctx context.Context, req *poll.CreatePollRequest) (*poll.CreatePollResponse, error) {
	return s.handlers.CreatePoll(ctx, req)
}

func (s *ServerManager) GetPoll(ctx context.Context, req *poll.GetPollRequest) (*poll.GetPollResponse, error) {
	return s.handlers.GetPoll(ctx, req)
}

func NewServerManager(config models.GRPCServer, handlers *handlers.HandlersManager) *ServerManager {
	return &ServerManager{
		config:   config,
		handlers: handlers,
	}
}

func (s *ServerManager) Start() {
	listener, err := net.Listen(s.config.Network, fmt.Sprintf(":%s", s.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.UnaryLoggingInterceptor),
	)
	poll.RegisterPollServiceServer(server, s)

	log.Println("poll gRPC server is running on port: ", s.config.Port)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}
}
