package services

import (
	"github.com/rivalnofirm/test_go_bank/domain/merchants"
	"github.com/rivalnofirm/test_go_bank/utils/errors"
)

var (
	MerchantService merchantsServiceInterface = &merchantsService{}
)

type merchantsService struct {
}

type merchantsServiceInterface interface {
	GetMerchant(int64) (*merchants.Merchant, *errors.RestErr)
}

func (s *merchantsService) GetMerchant(merchantId int64) (*merchants.Merchant, *errors.RestErr) {
	result := &merchants.Merchant{Id: merchantId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
