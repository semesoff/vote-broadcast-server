package others

import (
	"auth-service/pkg/models"
	"errors"
)

func CheckUserData(user models.User) error {
	switch {
	case user.Username == "" || user.Password == "":
		return errors.New("fields cannot be empty")
	case len(user.Password) < 8:
		return errors.New("password must be at least 8 characters long")
	case len(user.Username) < 4:
		return errors.New("username must be at least 4 characters long")
	}
	return nil
}
