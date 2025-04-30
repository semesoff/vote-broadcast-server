package utils

import (
	"websocket-service/pkg/models"
	"websocket-service/proto/websocket"
)

func ProtoPollsDataToModel(req *websocket.PollsRequest) []*models.Poll {
	polls := make([]*models.Poll, 0)
	for _, pollProto := range req.Polls {
		poll := &models.Poll{}
		poll.ID = int(pollProto.Id)
		poll.Title = pollProto.Title
		polls = append(polls, poll)
	}
	return polls
}

func ProtoVotesDataToModel(req *websocket.VotesRequest) *models.PollVotes {
	pollVotes := models.PollVotes{
		ID:      int(req.PollId),
		Options: make(map[int]models.Option),
	}
	for _, optionProto := range req.Options {
		option := models.Option{}
		option.CountVotes = int(optionProto.CountVotes)
		users := make([]models.User, 0)
		for _, userProto := range optionProto.Users {
			user := models.User{}
			user.ID = int(userProto.Id)
			user.Name = userProto.Name
			users = append(users, user)
		}
		option.Users = users
		pollVotes.Options[int(optionProto.Id)] = option
	}
	return &pollVotes
}

func NotifyChannels(dataChannels models.DataChannels, method models.MethodType, data interface{}) {
	dataChannels <- models.Data{
		Method: method,
		Data:   data,
	}
}
