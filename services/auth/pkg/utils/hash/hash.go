package hash

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"vote-broadcast-server/services/auth/pkg/models"
)

func HashPassword(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error hashing password")
	}
	user.Password = string(hashedPassword)
	return nil
}
