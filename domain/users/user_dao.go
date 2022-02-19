package users

import (
	"github.com/rivalnofirm/test_go_bank/datasorces/mysql"
	"github.com/rivalnofirm/test_go_bank/logger"
	"github.com/rivalnofirm/test_go_bank/utils/errors"
	"github.com/rivalnofirm/test_go_bank/utils/mysql_utils"
	"strings"
)

const (
	queryGetUser    = "SELECT id, first_name, last_name, email, password, amount, created_at FROM users WHERE id=?;"
	queryInsertUser = "INSERT INTO users (first_name, last_name, email, password, amount, created_at) VALUES(?, ?, ?, ?, ?, ?);"
	queryLogin      = "SELECT id, first_name, last_name, email, amount, created_at FROM users WHERE email=? AND password=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := mysql.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Amount, &user.CreatedAt); getErr != nil {
		logger.Error("when trying to get user by id", getErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := mysql.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, user.Amount, user.CreatedAt)
	if saveErr != nil {
		logger.Error("error when trying to save user", saveErr)
		return errors.NewInternalServerError("database error")
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return errors.NewInternalServerError("database error")
	}
	user.Id = userId
	return nil
}

func (user *User) FindByEmailAndPassword() *errors.RestErr {
	stmt, err := mysql.Client.Prepare(queryLogin)
	if err != nil {
		logger.Error("error when trying to prepare get user by email and password statement", err)
		return errors.NewInternalServerError("error when tying to find user")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Amount, &user.CreatedAt); getErr != nil {
		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
			return errors.NewNotFoundError("invalid user credentials")
		}
		logger.Error("error when trying to get user by email and password", getErr)
		return errors.NewInternalServerError("error when tying to find user")
	}
	return nil
}
