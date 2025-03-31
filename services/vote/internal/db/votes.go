package db

import (
	"errors"
	"fmt"
	"vote-broadcast-server/services/vote/pkg/models"
	"vote-broadcast-server/services/vote/pkg/utils"
)

// GetVotes Get poll votes by pollId
func (d *DatabaseManager) GetVotes(pollId int) (models.Votes, error) {
	rows, err := d.db.Query(
		"SELECT v.option_id, v.user_id, u.username FROM votes v JOIN users u ON v.user_id = u.id WHERE v.poll_id = $1;",
		pollId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	votes := make(models.Votes)
	for rows.Next() {
		var vote models.Vote
		vote.Users = []models.User{{}}
		vote.CountVotes = 1

		if err := rows.Scan(&vote.OptionId, &vote.Users[0].ID, &vote.Users[0].Name); err != nil {
			return nil, err
		}

		if _, ok := votes[vote.OptionId]; !ok {
			votes[vote.OptionId] = vote
		} else {
			hashedVote := votes[vote.OptionId]
			hashedVote.CountVotes++
			hashedVote.Users = append(hashedVote.Users, vote.Users[0])
			votes[vote.OptionId] = hashedVote
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return votes, nil
}

// CreateVote create poll vote by user_id, poll_id, option_id
func (d *DatabaseManager) CreateVote(userVote models.UserVote) error {
	tx, err := d.db.Begin()
	if err != nil {
		return errors.New("failed to begin transaction")
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	pollType, maxOptions, err := d.getTypePoll(userVote.PollId)
	if err != nil {
		return err
	}

	if err := checkCreateData(pollType, maxOptions, userVote); err != nil {
		return err
	}

	var exists bool
	err = tx.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM votes WHERE poll_id = $1 AND user_id = $2)",
		userVote.PollId, userVote.UserId,
	).Scan(&exists)

	if err != nil {
		return err
	}

	if exists {
		return errors.New("user already voted for this poll")
	}

	for _, optionId := range userVote.OptionsId {
		var valid bool
		err = tx.QueryRow(
			"SELECT EXISTS(SELECT 1 FROM poll_options WHERE id = $1 AND poll_id = $2)",
			optionId, userVote.PollId,
		).Scan(&valid)
		if err != nil {
			return err
		}
		if !valid {
			return errors.New("option does not belong to the poll")
		}

		_, err = tx.Exec(
			"INSERT INTO votes (user_id, poll_id, option_id) VALUES ($1, $2, $3)",
			userVote.UserId, userVote.PollId, optionId,
		)
		if err != nil {
			return err
		}
	}

	return err
}

func (d *DatabaseManager) getTypePoll(pollId int) (models.PollType, int, error) {
	var pollType string
	var maxOptions int
	err := d.db.QueryRow(
		"SELECT poll_type, max_options FROM polls WHERE id = $1",
		pollId,
	).Scan(&pollType, &maxOptions)

	if err != nil {
		return models.PollType(-1), -1, err
	}

	return utils.ConvertStringToPollType(pollType), maxOptions, nil
}

func checkCreateData(pollType models.PollType, maxOptions int, userVote models.UserVote) error {
	if countOptions := len(userVote.OptionsId); pollType == models.Single {
		if countOptions != 1 {
			return errors.New("single poll can only have one option")
		}
	} else if pollType == models.Multiple {
		if countOptions < 1 {
			return errors.New("multiple poll must have at least one option")
		} else if countOptions > maxOptions {
			return fmt.Errorf("multiple polls can have maximum %d options", maxOptions)
		}
	}
	return nil
}
