package transaction

import "github.com/billykore/kore/backend/pkg/types"

// CreditCard payment method.
type CreditCard struct {
	Name       string
	CardNumber string
}

// NewCreditCard return instance of CreditCard.
func NewCreditCard(name, number string) *CreditCard {
	return &CreditCard{
		Name:       name,
		CardNumber: number,
	}
}

// Pay amount using CreditCard.
func (cc *CreditCard) Pay(amount types.Money) (*PaymentResponse, error) {
	return &PaymentResponse{
		Method:    creditCard,
		Amount:    amount,
		PaymentId: 999,
	}, nil
}
