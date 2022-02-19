package merchants

import (
	"github.com/rivalnofirm/test_go_bank/utils/errors"
	"strings"
)

const (
	AmountDefault = 0
)

type Merchant struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Amount    int64  `json:"amount"`
	CreatedAt string `json:"created_at"`
}

type Merchants []Merchant

func (merchant *Merchant) ValidateMerchant() *errors.RestErr {
	merchant.FirstName = strings.TrimSpace(merchant.FirstName)
	merchant.LastName = strings.TrimSpace(merchant.LastName)

	merchant.FirstName = strings.TrimSpace(strings.ToLower(merchant.FirstName))
	if merchant.FirstName == "" {
		return errors.NewBadRequestError("invalid first name")
	}

	merchant.LastName = strings.TrimSpace(strings.ToLower(merchant.LastName))
	if merchant.LastName == "" {
		return errors.NewBadRequestError("invalid last name")
	}

	merchant.Email = strings.TrimSpace(strings.ToLower(merchant.Email))
	if merchant.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	merchant.Password = strings.TrimSpace(merchant.Password)
	if merchant.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
