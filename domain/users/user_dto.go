package users

import (
	"net/mail"
	"strings"

	"github.com/smbirch/bookstore_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return errors.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if len(user.Password) < 6 {
		return errors.NewBadRequestError("invalid password: must be 6+ characters. one must be special character")
	}

	if !strings.ContainsAny(user.Password, ",./!@#$%^&*()-_+=\\|}]{[:;") {
		return errors.NewBadRequestError("invalid password: must be 6+ characters. one must be special character")

	}

	return nil
}
