package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"vote-service/pkg/handlers"
	"vote-service/pkg/middleware"
	"vote-service/pkg/models"
	"vote-service/proto/vote"
)

type ServerManager struct {
	config   models.GRPCServer
	handlers *handlers.HandlersManager
	vote.VoteServiceServer
}

func (s *ServerManager) GetVotes(ctx context.Context, req *vote.GetVotesRequest) (*vote.GetVotesResponse, error) {
	return s.handlers.GetVotes(ctx, req)
}

func (s *ServerManager) CreateVote(ctx context.Context, req *vote.CreateVoteRequest) (*vote.CreateVoteResponse, error) {
	return s.handlers.CreateVote(ctx, req)
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
		log.Fatalf("Failed to listen: %v", err)
		return
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.UnaryLoggingInterceptor),
	)

	vote.RegisterVoteServiceServer(server, s)

	log.Println("vote gGRPC server is running on port: ", s.config.Port)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
		return
	}
}
