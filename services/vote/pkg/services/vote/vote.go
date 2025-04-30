package vote

import (
	"vote-service/pkg/models"
	"vote-service/pkg/services"
)

type VoteManager struct {
	*services.ServiceManager
}

type Vote interface {
	GetVotes(pollId int) (models.Votes, error)
	CreateVote(userVote models.UserVote) error
}

func NewVoteManager(service *services.ServiceManager) *VoteManager {
	return &VoteManager{
		service,
	}
}

func (v *VoteManager) GetVotes(pollId int) (models.Votes, error) {
	votes, err := v.Db.GetVotes(pollId)
	if err != nil {
		return nil, err
	}
	return votes, nil
}

func (v *VoteManager) CreateVote(userVote models.UserVote) error {
	err := v.Db.CreateVote(userVote)
	return err
}
