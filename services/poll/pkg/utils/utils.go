package utils

import (
	pollProto "vote-broadcast-server/proto/poll"
	"vote-broadcast-server/proto/websocket"
	"vote-broadcast-server/services/poll/pkg/models"
)

// Poll.proto

func ConvertToProtoPolls(polls []models.Poll) []*pollProto.Poll {
	protoPolls := make([]*pollProto.Poll, len(polls))
	for i, poll := range polls {
		protoPolls[i] = convertToProtoPoll(poll)
	}
	return protoPolls
}

func convertToProtoPoll(poll models.Poll) *pollProto.Poll {
	return &pollProto.Poll{
		Id:    int64(poll.ID),
		Title: poll.Title,
		Type:  pollProto.PollType(poll.Type),
	}
}

func ConvertToProtoPollData(poll models.Poll) *pollProto.PollData {
	pollType := pollProto.PollType(poll.Type)
	return &pollProto.PollData{
		Id:      int64(poll.ID),
		Title:   poll.Title,
		Type:    &pollType,
		Options: convertToProtoOptions(poll.Options),
	}
}

func convertToProtoOptions(options []models.Option) []*pollProto.Option {
	protoOptions := make([]*pollProto.Option, len(options))
	for i, option := range options {
		protoOptions[i] = &pollProto.Option{
			Id:   int64(option.ID),
			Text: option.Text,
		}
	}
	return protoOptions
}

func StringToPollType(pollType string) models.PollType {
	switch pollType {
	case "single":
		return models.Single
	case "multiple":
		return models.Multiple
	default:
		return -1
	}
}

// Websocket.proto

func ConvertToProtoWebSocketPolls(polls []models.Poll) []*websocket.Poll {
	response := make([]*websocket.Poll, len(polls))
	for i, poll := range polls {
		response[i] = &websocket.Poll{
			Id:    int64(poll.ID),
			Title: poll.Title,
		}
	}

	return response
}
