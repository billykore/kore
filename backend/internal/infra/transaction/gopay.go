package transaction

import "github.com/billykore/kore/backend/pkg/types"

// GoPay payment method.
type GoPay struct {
	Name        string
	PhoneNumber string
}

// NewGoPay return instance of GoPay.
func NewGoPay(name, phoneNumber string) *GoPay {
	return &GoPay{
		Name:        name,
		PhoneNumber: phoneNumber,
	}
}

// Pay amount using GoPay.
func (p *GoPay) Pay(amount types.Money) (*PaymentResponse, error) {
	return &PaymentResponse{
		Method:    gopay,
		Amount:    amount,
		PaymentId: 999,
	}, nil
}
