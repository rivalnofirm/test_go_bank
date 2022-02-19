package users

import (
	"github.com/rivalnofirm/test_go_bank/utils/errors"
	"strings"
)

const (
	AmountDefault = 0
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Amount    int64  `json:"amount"`
	CreatedAt string `json:"created_at"`
}

type Users []User

func (user *User) ValidateUser() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	if user.FirstName == "" {
		return errors.NewBadRequestError("invalid first name")
	}

	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	if user.LastName == "" {
		return errors.NewBadRequestError("invalid last name")
	}

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
