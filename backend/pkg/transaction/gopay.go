package transaction

import "github.com/billykore/kore/backend/pkg/types"

type GoPay struct {
	Name        string
	PhoneNumber string
}

func NewGoPay(name, phoneNumber string) *GoPay {
	return &GoPay{
		Name:        name,
		PhoneNumber: phoneNumber,
	}
}

func (p *GoPay) Pay(amount types.Money) (*PaymentResponse, error) {
	return &PaymentResponse{
		Method:    gopay,
		Amount:    amount,
		PaymentId: 999,
	}, nil
}
