package handlers

import (
	"context"
	"vote-broadcast-server/proto/websocket"
	"vote-broadcast-server/services/websocket/pkg/models"
	"vote-broadcast-server/services/websocket/pkg/utils"
)

type HandlersManager struct {
	dataChannels models.DataChannels
}

func NewHandlersManager(dataChannels models.DataChannels) *HandlersManager {
	return &HandlersManager{
		dataChannels: dataChannels,
	}
}

type Handlers interface {
	GetPolls(ctx context.Context, req *websocket.PollsRequest) (*websocket.PollsResponse, error)
	GetVotes(ctx context.Context, req *websocket.VotesRequest) (*websocket.VotesResponse, error)
}

func (h *HandlersManager) GetPolls(ctx context.Context, req *websocket.PollsRequest) (*websocket.PollsResponse, error) {
	data := utils.ProtoPollsDataToModel(req)
	utils.NotifyChannels(h.dataChannels, models.GetPolls, data)
	response := &websocket.PollsResponse{
		Success: true,
	}

	return response, nil
}

func (h *HandlersManager) GetVotes(ctx context.Context, req *websocket.VotesRequest) (*websocket.VotesResponse, error) {
	data := utils.ProtoVotesDataToModel(req)
	utils.NotifyChannels(h.dataChannels, models.GetVotes, *data)

	response := &websocket.VotesResponse{
		Success: true,
	}

	return response, nil
}
