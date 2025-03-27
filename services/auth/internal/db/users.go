package db

import (
	"database/sql"
	"errors"
	"vote-broadcast-server/services/auth/pkg/models"
)

func (d *DatabaseManager) AddUser(user models.User) (models.UserWithID, error) {
	userWithID := models.UserWithID{
		Username: user.Username,
	}
	err := (*d).db.QueryRow("INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id",
		user.Username, user.Password).Scan(&userWithID.ID)
	return userWithID, err
}

func (d *DatabaseManager) GetUser(user models.User) (models.UserWithID, bool, error) {
	var exists bool
	userWithID := models.UserWithID{
		Username: user.Username,
	}
	err := (*d).db.QueryRow("SELECT true, id, username FROM users WHERE username=$1",
		user.Username).Scan(&exists, &userWithID.ID, &userWithID.Username)

	if errors.Is(err, sql.ErrNoRows) {
		return userWithID, false, nil
	}

	return userWithID, exists, err
}

func (d *DatabaseManager) GetUserWithPassword(user models.User) (models.UserWithPassword, bool, error) {
	var exists bool
	userWithID := models.UserWithPassword{
		Username: user.Username,
	}
	err := (*d).db.QueryRow("SELECT true, id, username, password FROM users WHERE username=$1",
		user.Username).Scan(&exists, &userWithID.ID, &userWithID.Username, &userWithID.Password)

	if errors.Is(err, sql.ErrNoRows) {
		return userWithID, false, nil
	}

	return userWithID, exists, err
}
