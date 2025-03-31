package handlers

import (
	"context"
	voteProto "vote-broadcast-server/proto/vote"
	"vote-broadcast-server/services/vote/pkg/services"
	"vote-broadcast-server/services/vote/pkg/services/vote"
	"vote-broadcast-server/services/vote/pkg/utils"
)

type HandlersManager struct {
	vote vote.Vote
}

func NewHandlersManager(service *services.ServiceManager) *HandlersManager {
	return &HandlersManager{
		vote: vote.NewVoteManager(service),
	}
}

type Handlers interface {
	GetVotes(ctx context.Context, req *voteProto.GetVotesRequest) (*voteProto.GetVotesResponse, error)
	CreateVote(ctx context.Context, req *voteProto.CreateVoteRequest) (*voteProto.CreateVoteResponse, error)
}

func (h *HandlersManager) GetVotes(ctx context.Context, req *voteProto.GetVotesRequest) (*voteProto.GetVotesResponse, error) {
	if err := utils.CheckGetVotesData(req); err != nil {
		return nil, err
	}

	votes, err := h.vote.GetVotes(int(req.PollId))
	if err != nil {
		return nil, err
	}

	response := utils.ToGetVotesProtoData(votes)

	return response, nil
}

func (h *HandlersManager) CreateVote(ctx context.Context, req *voteProto.CreateVoteRequest) (*voteProto.CreateVoteResponse, error) {
	if err := utils.CheckCreateVoteData(req); err != nil {
		return nil, err
	}

	userVote := utils.ProtoToCreateVoteData(req)

	err := h.vote.CreateVote(*userVote)
	if err != nil {
		return nil, err
	}

	response := &voteProto.CreateVoteResponse{
		Success: true,
	}

	return response, nil
}
