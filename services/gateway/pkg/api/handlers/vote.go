package handlers

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
	"vote-broadcast-server/proto/vote"
	"vote-broadcast-server/services/gateway/pkg/models"
	"vote-broadcast-server/services/gateway/pkg/utils"
)

func (h *HandlersManager) GetVotes(c *gin.Context) {
	var request vote.GetVotesRequest

	pollStrId := c.Param("id")
	if pollStrId == "" {
		utils.RespondWithError(c.Writer, http.StatusBadRequest, "pollId is required")
		return
	}

	pollIntId, err := strconv.Atoi(pollStrId)
	if err != nil {
		utils.RespondWithError(c.Writer, http.StatusBadRequest, "pollId must be int")
		return
	}

	request.PollId = int64(pollIntId)

	// gRPC Client
	serviceClient, conn, err := h.getGRPCService("vote")
	if err != nil {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	voteClient, ok := serviceClient.(vote.VoteServiceClient)
	if !ok {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, "Failed to cast client to VoteServiceClient")
		return
	}

	// gRPC Request
	response, err := voteClient.GetVotes(context.Background(), &request)
	if err != nil {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	if err := utils.RespondWithJSON(c.Writer, http.StatusOK, response); err != nil {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *HandlersManager) CreateVote(c *gin.Context) {
	var voteData models.Vote

	if err := json.NewDecoder(c.Request.Body).Decode(&voteData); err != nil {
		utils.RespondWithError(c.Writer, http.StatusBadRequest, err.Error())
		return
	}

	request := utils.ToProtoCreateVoteData(voteData)

	// gRPC Client
	serviceClient, conn, err := h.getGRPCService("vote")
	if err != nil {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	voteClient, ok := serviceClient.(vote.VoteServiceClient)
	if !ok {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, "Failed to cast client to VoteServiceClient")
		return
	}

	// gRPC Request
	response, err := voteClient.CreateVote(context.Background(), request)
	if err != nil {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	if err := utils.RespondWithJSON(c.Writer, http.StatusOK, response); err != nil {
		utils.RespondWithError(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
}
