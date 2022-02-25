package main

import (
	"github.com/rivalnofirm/test_go_bank/api"
	"github.com/rivalnofirm/test_go_bank/database"
)

func main() {
	database.InitDatabase()
	api.StartApi()
}
