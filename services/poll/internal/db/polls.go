package db

import (
	"vote-broadcast-server/services/poll/pkg/models"
)

func (d *DatabaseManager) GetPolls() ([]models.Poll, error) {
	rows, err := d.db.Query("SELECT id, title FROM polls")
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	polls := make([]models.Poll, 0)
	for rows.Next() {
		var poll models.Poll
		if err := rows.Scan(&poll.ID, &poll.Title); err != nil {
			return nil, err
		}
		polls = append(polls, poll)
	}
	return polls, nil
}

func (d *DatabaseManager) GetPoll(poll models.Poll) (models.Poll, error) {
	err := d.db.QueryRow("SELECT id, title FROM polls WHERE id = $1", poll.ID).Scan(&poll.ID, &poll.Title)
	if err != nil {
		return models.Poll{}, err
	}

	rows, err := d.db.Query("SELECT id, option_text FROM poll_options WHERE poll_id = $1", poll.ID)
	defer rows.Close()
	if err != nil {
		return models.Poll{}, err
	}

	for rows.Next() {
		var option models.Option
		if err := rows.Scan(&option.ID, &option.Text); err != nil {
			return models.Poll{}, err
		}
		poll.Options = append(poll.Options, option)
	}

	return poll, nil
}

func (d *DatabaseManager) CreatePoll(poll models.Poll, userId int) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}

	err = tx.QueryRow("INSERT INTO polls (title, creator_id, poll_type) VALUES ($1, $2, $3)  RETURNING id", poll.Title, userId, poll.Type.String()).Scan(&poll.ID)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	for _, option := range poll.Options {
		_, err := tx.Exec("INSERT INTO poll_options (poll_id, option_text) VALUES ($1, $2)", poll.ID, option.Text)
		if err != nil {
			if err := tx.Rollback(); err != nil {
				return err
			}
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
