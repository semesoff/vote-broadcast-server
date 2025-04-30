package notification_service

import (
	"vote-service/pkg/models"
	"vote-service/pkg/services"
	"vote-service/pkg/services/vote"
)

type NotificationServiceManager struct {
	dataChannel models.DataChannel
	voteManager vote.Vote
}

func NewNotificationServiceManager(dataChannel models.DataChannel, service *services.ServiceManager) *NotificationServiceManager {
	return &NotificationServiceManager{
		dataChannel: dataChannel,
		voteManager: vote.NewVoteManager(service),
	}
}

type NotificationService interface {
	GetVotes(pollId int)
}

func (s *NotificationServiceManager) GetVotes(pollId int) {
	votes, err := s.voteManager.GetVotes(pollId)
	if err != nil {
		return
	}

	var pollVotes models.PollVotes
	pollVotes.PollId = pollId
	pollVotes.Votes = votes

	s.dataChannel <- models.Data{
		Method: models.GetVotes,
		Data:   pollVotes,
	}
}
