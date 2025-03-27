package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"vote-broadcast-server/proto/auth"
	"vote-broadcast-server/proto/poll"
	"vote-broadcast-server/services/gateway/pkg/config"
	"vote-broadcast-server/services/gateway/pkg/models"
)

type HandlersManager struct {
	services map[string]models.ServiceConfig
}

type Handlers interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
	GetPolls(c *gin.Context)
	GetPoll(c *gin.Context)
	CreatePoll(c *gin.Context)
}

func NewHandlersManager(c config.ConfigProvider) *HandlersManager {
	return &HandlersManager{
		services: c.GetConfig().Services,
	}
}

func (h *HandlersManager) getGRPCService(serviceName string) (interface{}, *grpc.ClientConn, error) {
	service, ok := h.services[serviceName]
	if !ok {
		return nil, nil, errors.New("service not found: " + serviceName)
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
	default:
		return nil, nil, errors.New("service not found: " + serviceName)
	}

	return clientInstance, conn, nil
}
