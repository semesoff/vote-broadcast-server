package handlers

import (
	"context"
	pollProto "vote-broadcast-server/proto/poll"
	"vote-broadcast-server/services/poll/pkg/models"
	"vote-broadcast-server/services/poll/pkg/services"
	"vote-broadcast-server/services/poll/pkg/services/notification_service"
	"vote-broadcast-server/services/poll/pkg/services/poll"
	"vote-broadcast-server/services/poll/pkg/utils"
)

type HandlersManager struct {
	poll                poll.Poll
	notificationService notification_service.NotificationService
}

func NewHandlersManager(service *services.ServiceManager, notificationService notification_service.NotificationService) *HandlersManager {
	return &HandlersManager{
		poll:                poll.NewPollService(service),
		notificationService: notificationService,
	}
}

func (h *HandlersManager) GetPollManager() *poll.Poll {
	return &h.poll
}

type Handlers interface {
	GetPolls(ctx context.Context, req *pollProto.GetPollsRequest) (*pollProto.GetPollsResponse, error)
	CreatePoll(ctx context.Context, req *pollProto.CreatePollRequest) (*pollProto.CreatePollResponse, error)
	GetPoll(ctx context.Context, req *pollProto.GetPollRequest) (*pollProto.GetPollResponse, error)
}

func (h *HandlersManager) GetPolls(ctx context.Context, req *pollProto.GetPollsRequest) (*pollProto.GetPollsResponse, error) {
	if req == nil {
		return nil, models.ErrEmptyRequest{}
	}

	polls, err := h.poll.GetPolls()
	if err != nil {
		return nil, err
	}

	response := &pollProto.GetPollsResponse{
		Polls: utils.ConvertToProtoPolls(polls),
	}

	return response, nil
}

func (h *HandlersManager) CreatePoll(ctx context.Context, req *pollProto.CreatePollRequest) (*pollProto.CreatePollResponse, error) {
	if err := checkCreatePollData(req); err != nil {
		return nil, err
	}

	pollData, userId := getCreatePollData(req)
	if err := h.poll.CreatePoll(*pollData, userId); err != nil {
		return nil, err
	}

	// notify channel about updated polls
	h.notificationService.GetPolls()

	response := &pollProto.CreatePollResponse{
		Success: true,
	}

	return response, nil
}

func (h *HandlersManager) GetPoll(ctx context.Context, req *pollProto.GetPollRequest) (*pollProto.GetPollResponse, error) {
	if req.Id == 0 {
		return nil, models.ErrEmptyRequest{}
	}

	requestedPoll := models.Poll{
		ID: int(req.Id),
	}

	pollData, err := h.poll.GetPoll(requestedPoll)
	if err != nil {
		return nil, err
	}

	response := &pollProto.GetPollResponse{
		Poll: utils.ConvertToProtoPollData(pollData),
	}

	return response, nil
}

func checkCreatePollData(req *pollProto.CreatePollRequest) error {
	switch {
	case req == nil:
		return models.ErrEmptyRequest{}
	case req.Poll == nil:
		return models.ErrEmptyData{}
	case req.Poll.Title == "":
		return models.ErrEmptyData{}
	case req.Poll.Options == nil || len(req.Poll.Options) == 0:
		return models.ErrEmptyData{}
	case !(models.MinPollType <= req.Poll.Type && req.Poll.Type <= models.MaxPollType):
		return models.ErrInvalidData{}
	case len(req.Poll.Title) > models.MaxPollTitle || len(req.Poll.Title) < models.MinPollTitle:
		return models.ErrInvalidData{}
	case len(req.Poll.Options) > models.MaxPollOptions:
		return models.ErrInvalidData{}
	default:
		for _, option := range req.Poll.Options {
			if lenText := len(option.Text); lenText > models.MaxOptionText || lenText < models.MinOptionText {
				return models.ErrInvalidData{}
			}
		}
		return nil
	}
}

func getCreatePollData(req *pollProto.CreatePollRequest) (*models.Poll, int) {
	userId := int(req.Poll.UserId)

	pollData := models.Poll{
		ID:      0,
		Title:   req.Poll.Title,
		Type:    models.PollType(req.Poll.Type),
		Options: []models.Option{},
	}

	for _, option := range req.Poll.Options {
		pollData.Options = append(pollData.Options, models.Option{
			ID:   0,
			Text: option.Text,
		})
	}

	pollData.MaxOptions = len(pollData.Options)

	return &pollData, userId
}
