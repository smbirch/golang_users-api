//this file is the access layer to the database

package users

import (
	"fmt"

	"github.com/smbirch/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already exists", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists", user.Id))
	}

	usersDB[user.Id] = user
	return nil
}
