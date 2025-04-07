package handlers

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
	"strings"
	"vote-broadcast-server/proto/poll"
	"vote-broadcast-server/services/gateway/pkg/models"
	"vote-broadcast-server/services/gateway/pkg/utils"
)

func (h *HandlersManager) GetPolls(c *gin.Context) {
	var request poll.GetPollsRequest

	// gRPC Client
	serviceClient, conn, err := h.getGRPCService("poll")
	if err != nil {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	pollClient, ok := serviceClient.(poll.PollServiceClient)
	if !ok {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, "Failed to cast client to PollServiceClient")
		return
	}

	// gRPC Request
	response, err := pollClient.GetPolls(context.Background(), &request)
	if err != nil {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	if err := utils.RespondWithJSON(c.Writer, http.StatusOK, response); err != nil {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *HandlersManager) GetPoll(c *gin.Context) {
	var request poll.GetPollRequest

	pollStrId := c.Param("id")
	if pollStrId == "" {
		utils.RespondWithError(c.Writer, http.StatusBadRequest, "pollId is required")
		return
	}

	pollIntId, err := strconv.Atoi(pollStrId)
	if err != nil {
		utils.RespondWithError(c.Writer, http.StatusBadRequest, "pollId must be an integer")
		return
	}

	request.Id = int64(pollIntId)

	// gRPC Client
	serviceClient, conn, err := h.getGRPCService("poll")
	if err != nil {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	pollClient, ok := serviceClient.(poll.PollServiceClient)
	if !ok {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, "Failed to cast client to PollServiceClient")
		return
	}

	// gRPC Request
	response, err := pollClient.GetPoll(context.Background(), &request)
	if err != nil {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	if err := utils.RespondWithJSON(c.Writer, http.StatusOK, response); err != nil {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *HandlersManager) CreatePoll(c *gin.Context) {
	var pollRequest models.PollRequest

	if err := json.NewDecoder(c.Request.Body).Decode(&pollRequest); err != nil {
		utils.RespondWithError(c.Writer, http.StatusBadRequest, "Invalid request payload")
		return
	}

	token := strings.Split(c.GetHeader("Authorization"), " ")[1]
	if token == "" {
		utils.RespondWithError(c.Writer, http.StatusBadRequest, "Authorization token is required")
		return
	}

	dataFromToken, err := h.jwtProvider.VerifyToken(token)
	if err != nil {
		utils.RespondWithError(c.Writer, http.StatusUnauthorized, "Invalid token")
		return
	}

	if err := validationTokenData(*dataFromToken); err != nil {
		utils.RespondWithError(c.Writer, http.StatusUnauthorized, "Invalid token")
		return
	}

	userIntId, err := strconv.Atoi(dataFromToken.UserID)
	if err != nil {
		utils.RespondWithError(c.Writer, http.StatusBadRequest, "userId must be an integer")
		return
	}

	pollData := models.Poll{
		Title:   pollRequest.Title,
		Type:    pollRequest.Type,
		Options: pollRequest.Options,
		UserID:  userIntId,
	}

	request := &poll.CreatePollRequest{}
	request.Poll = pollDataToPollCreateData(pollData)

	// gRPC Client
	serviceClient, conn, err := h.getGRPCService("poll")
	if err != nil {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	pollClient, ok := serviceClient.(poll.PollServiceClient)
	if !ok {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, "Failed to cast client to PollServiceClient")
		return
	}

	// gRPC Request
	response, err := pollClient.CreatePoll(context.Background(), request)
	if err != nil {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	if err := utils.RespondWithJSON(c.Writer, http.StatusOK, response); err != nil {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
}

func pollDataToPollCreateData(pollData models.Poll) *poll.PollCreateData {
	var pollCreateData poll.PollCreateData
	pollCreateData.Title = pollData.Title
	pollCreateData.Type = poll.PollType(pollData.Type)
	pollCreateData.UserId = int64(pollData.UserID)
	pollCreateData.Options = []*poll.Option{}

	for _, option := range pollData.Options {
		pollCreateData.Options = append(pollCreateData.Options, &poll.Option{
			Id:   int64(option.ID),
			Text: option.Text,
		})
	}

	return &pollCreateData
}
