package database

import (
	"github.com/jinzhu/gorm"
	"github.com/rivalnofirm/test_go_bank/helpers"
)

// Create global variable
var DB *gorm.DB

// Create InitDatabase function
func InitDatabase() {
	database, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=bank password=kepo12ad sslmode=disable")
	helpers.HandleErr(err)
	// Set up connection pool
	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)
	DB = database
}
