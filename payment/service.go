package payment

import (
	"strconv"

	"github.com/gemm123/crowdfunding/user"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type service struct{}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	midtrans.ServerKey = ""
	midtrans.Environment = midtrans.Sandbox

	var snapGateway = snap.Client{}
	snapGateway.New(midtrans.ServerKey, midtrans.Environment)

	snapReq := &snap.Request{
		CustomerDetail: &midtrans.CustomerDetails{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	paymentURL, err := snapGateway.CreateTransactionUrl(snapReq)
	if err != nil {
		return "", err
	}

	return paymentURL, nil
}
