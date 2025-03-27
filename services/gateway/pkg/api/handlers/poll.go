package handlers

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
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
	var pollData models.Poll

	if err := json.NewDecoder(c.Request.Body).Decode(&pollData); err != nil {
		utils.RespondWithError(c.Writer, http.StatusBadRequest, "Invalid request payload")
		return
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
