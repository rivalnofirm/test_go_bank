package app

import (
	"github.com/rivalnofirm/test_go_bank/controllers/merchants"
	"github.com/rivalnofirm/test_go_bank/controllers/users"
)

func mapUrls() {
	// endpoint user
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users/", users.CreateUser)
	router.POST("/users/login/", users.Login)

	// endpoint merchant
	router.GET("/merchants/:merchant_id", merchants.GetMerchant)
}
