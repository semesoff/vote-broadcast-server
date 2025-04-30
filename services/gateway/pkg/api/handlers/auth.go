package handlers

import (
	"context"
	"encoding/json"
	"gateway-service/pkg/utils"
	"gateway-service/proto/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
)

func (h *HandlersManager) RegisterUser(c *gin.Context) {
	var request auth.RegisterRequest

	if err := json.NewDecoder(c.Request.Body).Decode(&request); err != nil {
		utils.RespondWithError(c.Writer, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// gRPC Client
	serviceClient, conn, err := h.getGRPCService("auth")
	if err != nil {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	authClient, ok := serviceClient.(auth.AuthServiceClient)
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
	var request auth.LoginRequest

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

	authClient, ok := serviceClient.(auth.AuthServiceClient)
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
