package merchants

import (
	"github.com/rivalnofirm/test_go_bank/datasorces/mysql"
	"github.com/rivalnofirm/test_go_bank/logger"
	"github.com/rivalnofirm/test_go_bank/utils/errors"
)

const (
	queryGetMerchantById = "SELECT id, first_name, last_name, email, password, amount, created_at FROM merchants WHERE id=?;"
	queryInsertUser      = "INSERT INTO users (first_name, last_name, email, password, amount, created_at) VALUES(?, ?, ?, ?, ?, ?);"
	queryLogin           = "SELECT id, first_name, last_name, email, amount, created_at FROM users WHERE email=? AND password=?;"
)

func (merchant *Merchant) Get() *errors.RestErr {
	stmt, err := mysql.Client.Prepare(queryGetMerchantById)
	if err != nil {
		logger.Error("error when trying to prepare get merchant statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(merchant.Id)
	if getErr := result.Scan(
		&merchant.Id,
		&merchant.FirstName,
		&merchant.LastName,
		&merchant.Email,
		&merchant.Password,
		&merchant.Amount,
		&merchant.CreatedAt,
	); getErr != nil {
		logger.Error("when trying to get merchant by id", getErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}
