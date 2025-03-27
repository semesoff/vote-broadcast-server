package poll

import (
	"vote-broadcast-server/services/poll/pkg/models"
	"vote-broadcast-server/services/poll/pkg/services"
)

type PollManager struct {
	*services.ServiceManager
}

type Poll interface {
	GetPolls() ([]models.Poll, error)
	GetPoll(poll models.Poll) (models.Poll, error)
	CreatePoll(poll models.Poll, userId int) error
}

func NewPollService(service *services.ServiceManager) *PollManager {
	return &PollManager{
		service,
	}
}

func (s *PollManager) GetPolls() ([]models.Poll, error) {
	polls, err := s.Db.GetPolls()
	if err != nil {
		return nil, err
	}
	return polls, nil
}

func (s *PollManager) GetPoll(poll models.Poll) (models.Poll, error) {
	fullPollData, err := s.Db.GetPoll(poll)
	if err != nil {
		return models.Poll{}, err
	}
	return fullPollData, err
}

func (s *PollManager) CreatePoll(poll models.Poll, userId int) error {
	err := s.Db.CreatePoll(poll, userId)
	return err
}
