package notification_service

import (
	"poll-service/pkg/models"
	"poll-service/pkg/services"
	"poll-service/pkg/services/poll"
)

type NotificationServiceManager struct {
	dataChannel models.DataChannel
	pollManager poll.Poll
}

func NewNotificationServiceManager(dataChannel models.DataChannel, service *services.ServiceManager) *NotificationServiceManager {
	return &NotificationServiceManager{
		dataChannel: dataChannel,
		pollManager: poll.NewPollService(service),
	}
}

type NotificationService interface {
	GetPolls()
}

func (s *NotificationServiceManager) GetPolls() {
	polls, err := s.pollManager.GetPolls()
	if err != nil {
		return
	}

	s.dataChannel <- models.Data{
		Method: models.GetPolls,
		Data:   polls,
	}
}
