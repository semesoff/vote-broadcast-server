package utils

import (
	"vote-broadcast-server/proto/websocket"
	"vote-broadcast-server/services/websocket/pkg/models"
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

func ProtoVotesDataToModel(req *websocket.VotesRequest) map[int]*models.Option {
	options := make(map[int]*models.Option)
	for _, optionProto := range req.Options {
		option := &models.Option{}
		option.ID = int(optionProto.Id)
		option.CountVotes = int(optionProto.CountVotes)
		users := make([]models.User, 0)
		for _, userProto := range optionProto.Users {
			user := models.User{}
			user.ID = int(userProto.Id)
			user.Name = userProto.Name
			users = append(users, user)
		}
		option.Users = users
	}
	return options
}

func NotifyChannels(dataChannels models.DataChannels, method models.MethodType, data interface{}) {
	dataChannels <- models.Data{
		Method: method,
		Data:   data,
	}
}
