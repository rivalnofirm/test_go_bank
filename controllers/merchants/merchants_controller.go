package merchants

import (
	"github.com/gin-gonic/gin"
	"github.com/rivalnofirm/test_go_bank/services"
	"github.com/rivalnofirm/test_go_bank/utils/errors"
	"net/http"
	"strconv"
)

func getMerchantId(MerchantIdParam string) (int64, *errors.RestErr) {
	merchantId, userErr := strconv.ParseInt(MerchantIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("merchant id should be a number")
	}
	return merchantId, nil
}

func GetMerchant(c *gin.Context) {
	merchantId, idErr := getMerchantId(c.Param("merchant_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	merchant, getErr := services.MerchantService.GetMerchant(merchantId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, merchant)
}
