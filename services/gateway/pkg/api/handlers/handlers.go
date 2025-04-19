package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"vote-broadcast-server/proto/auth"
	"vote-broadcast-server/proto/poll"
	"vote-broadcast-server/proto/vote"
	"vote-broadcast-server/services/gateway/pkg/config"
	"vote-broadcast-server/services/gateway/pkg/models"
	jwtModel "vote-broadcast-server/services/gateway/pkg/services/jwt"
)

type HandlersManager struct {
	grpcConnections map[string]*grpc.ClientConn
	services        map[string]models.ServiceConfig
	jwtProvider     jwtModel.JWTProvider
}

type Handlers interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
	GetPolls(c *gin.Context)
	GetPoll(c *gin.Context)
	CreatePoll(c *gin.Context)
	GetVotes(c *gin.Context)
	CreateVote(c *gin.Context)
}

func NewHandlersManager(c config.ConfigProvider) *HandlersManager {
	return &HandlersManager{
		services:        c.GetConfig().Services,
		jwtProvider:     jwtModel.NewJWTManager(c.GetConfig().JWTSecretKey),
		grpcConnections: make(map[string]*grpc.ClientConn),
	}
}

func (h *HandlersManager) getGRPCService(serviceName string) (interface{}, *grpc.ClientConn, error) {
	service, ok := h.services[serviceName]
	if !ok {
		return nil, nil, errors.New("service not found: " + serviceName)
	}

	// Check if the service is already connected
	if conn, exists := h.grpcConnections[serviceName]; exists {
		// Check if the connection is still alive
		fmt.Printf("state of %s: %s", serviceName, conn.GetState().String())
		return conn, nil, nil
	}

	conn, err := grpc.NewClient(service.URL, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, nil, err
	}

	var clientInstance interface{}
	switch serviceName {
	case "auth":
		clientInstance = auth.NewAuthServiceClient(conn)
	case "poll":
		clientInstance = poll.NewPollServiceClient(conn)
	case "vote":
		clientInstance = vote.NewVoteServiceClient(conn)
	default:
		return nil, nil, errors.New("service not found: " + serviceName)
	}

	return clientInstance, conn, nil
}

func validationTokenData(claims jwtModel.Claims) error {
	switch {
	case claims.UserID == "":
		return errors.New("user id is required")
	case claims.Username == "":
		return errors.New("user name is required")
	default:
		return nil
	}
}
