package payment

import (
	"github.com/billykore/kore/backend/internal/domain/order"
	"github.com/billykore/kore/backend/pkg/types"
)

const GoPayCompanyName = "GoPay by GoTo"

// GoPay payment method.
type GoPay struct {
	name string
}

// NewGoPay return instance of GoPay.
func NewGoPay() *GoPay {
	return &GoPay{
		name: GoPayCompanyName,
	}
}

// Pay pays amount using GoPay.
func (p *GoPay) Pay(srcName, srcAccount string, amount types.Money) (*order.PaymentResponse, error) {
	return &order.PaymentResponse{
		PaymentId:     999,
		Method:        "gopay",
		Amount:        amount,
		AccountName:   srcName,
		AccountNumber: srcAccount,
	}, nil
}
