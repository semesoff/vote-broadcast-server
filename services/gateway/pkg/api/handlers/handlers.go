package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"vote-broadcast-server/services/gateway/pkg/api/proto/gateway"
	"vote-broadcast-server/services/gateway/pkg/config"
	"vote-broadcast-server/services/gateway/pkg/models"
	"vote-broadcast-server/services/gateway/pkg/utils"
)

type HandlersManager struct {
	services map[string]models.ServiceConfig
}

type Handlers interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
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
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Println("Failed to close grpc connection (getGRPCService)")
		}
	}(conn)

	if err != nil {
		return nil, nil, err
	}

	var clientInstance interface{}
	switch serviceName {
	case "auth":
		clientInstance = gateway.NewAuthServiceClient(conn)
	default:
		return nil, nil, errors.New("service not found: " + serviceName)
	}

	return clientInstance, conn, nil
}

func (h *HandlersManager) RegisterUser(c *gin.Context) {
	var request gateway.RegisterRequest

	if err := json.NewDecoder(c.Request.Body).Decode(&request); err != nil {
		utils.RespondWithError(c.Writer, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// gRPC Client
	serviceClient, conn, err := h.getGRPCService("auth")
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	if err != nil {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	authClient, ok := serviceClient.(gateway.AuthServiceClient)
	if !ok {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, "Failed to cast client to AuthServiceClient")
		return
	}

	// gRPC Request
	response, err := authClient.RegisterUser(context.Background(), &request)
	if err != nil {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	if err := utils.RespondWithJSON(c.Writer, http.StatusOK, response); err != nil {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *HandlersManager) LoginUser(c *gin.Context) {
	var request gateway.LoginRequest

	if err := json.NewDecoder(c.Request.Body).Decode(&request); err != nil {
		utils.RespondWithError(c.Writer, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// gRPC Client
	serviceClient, conn, err := h.getGRPCService("auth")
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	if err != nil {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	authClient, ok := serviceClient.(gateway.AuthServiceClient)
	if !ok {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, "Failed to cast client to AuthServiceClient")
		return
	}

	// gRPC Request
	response, err := authClient.LoginUser(context.Background(), &request)
	if err != nil {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	if err := utils.RespondWithJSON(c.Writer, http.StatusOK, response); err != nil {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
}
