package utils

import (
	"errors"
	"vote-service/pkg/models"
	voteProto "vote-service/proto/vote"
	"vote-service/proto/websocket"
)

func CheckGetVotesData(req *voteProto.GetVotesRequest) error {
	if req.PollId <= 0 {
		return errors.New("PollId is must be greater than zero")
	}
	return nil
}

func ToGetVotesProtoData(votes models.Votes) *voteProto.GetVotesResponse {
	getVotesResponse := &voteProto.GetVotesResponse{
		Options: make([]*voteProto.Option, 0),
	}
	for optionId, optionData := range votes {
		option := &voteProto.Option{
			Id:         int64(optionId),
			CountVotes: int64(optionData.CountVotes),
			Users:      make([]*voteProto.User, 0),
		}
		for _, user := range optionData.Users {
			votedUser := &voteProto.User{
				Id:   int64(user.ID),
				Name: user.Name,
			}
			option.Users = append(option.Users, votedUser)
		}
		getVotesResponse.Options = append(getVotesResponse.Options, option)
	}
	return getVotesResponse
}

func CheckCreateVoteData(req *voteProto.CreateVoteRequest) error {
	switch {
	case req.PollId <= 0:
		return errors.New("PollId is must be greater than zero")
	case req.UserId <= 0:
		return errors.New("UserId is must be greater than zero")
	default:
		for _, optionId := range req.OptionsId {
			if optionId <= 0 {
				return errors.New("OptionId is must be greater than zero")
			}
		}
		return nil
	}
}

func ProtoToCreateVoteData(req *voteProto.CreateVoteRequest) *models.UserVote {
	userVote := &models.UserVote{
		PollId: int(req.PollId),
		UserId: int(req.UserId),
	}

	for _, optionId := range req.OptionsId {
		userVote.OptionsId = append(userVote.OptionsId, int(optionId))
	}

	return userVote
}

func ConvertStringToPollType(pollType string) models.PollType {
	switch pollType {
	case "single":
		return models.Single
	case "multiple":
		return models.Multiple
	default:
		return models.PollType(-1)
	}
}

func ConvertToProtoWebsocketVotes(pollVotes models.PollVotes) []*websocket.Option {
	response := make([]*websocket.Option, len(pollVotes.Votes))
	k := 0
	for optionId, optionData := range pollVotes.Votes {
		protoOption := &websocket.Option{}
		protoOption.Id = int64(optionId)
		protoOption.CountVotes = int64(optionData.CountVotes)
		protoOption.Users = make([]*websocket.User, len(optionData.Users))
		for i, user := range optionData.Users {
			protoOption.Users[i] = &websocket.User{
				Id:   int64(user.ID),
				Name: user.Name,
			}
		}
		response[k] = protoOption
		k++
	}
	return response
}
