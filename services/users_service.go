package services

import (
	"github.com/smbirch/bookstore_users-api/domain/users"
	"github.com/smbirch/bookstore_users-api/utils/errors"
)

func GetUser(userID int64) (*users.User, *errors.RestErr) {

}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
