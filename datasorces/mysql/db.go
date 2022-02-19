package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_bank_username = "mysql_bank_username"
	mysql_bank_password = "mysql_bank_password"
	mysql_bank_host     = "mysql_bank_host"
	mysql_bank_schema   = "mysql_bank_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(mysql_bank_username)
	password = os.Getenv(mysql_bank_password)
	host     = os.Getenv(mysql_bank_host)
	schema   = os.Getenv(mysql_bank_schema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	log.Println(fmt.Sprintf("about to connect to %s", dataSourceName))

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database succesfully configured")
}
