package services

import (
	"github.com/smbirch/bookstore_users-api/domain/users"
	"github.com/smbirch/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil

	// return &user, &errors.RestErr{
	// 	Status: http.StatusInternalServerError,
	// }
}
