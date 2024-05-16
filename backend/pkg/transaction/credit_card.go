package transaction

import "github.com/billykore/kore/backend/pkg/types"

type CreditCard struct {
	Name       string
	CardNumber string
}

func NewCreditCard(name, number string) *CreditCard {
	return &CreditCard{
		Name:       name,
		CardNumber: number,
	}
}

func (cc *CreditCard) Pay(amount types.Money) (*PaymentResponse, error) {
	return &PaymentResponse{
		Method:    creditCard,
		Amount:    amount,
		PaymentId: 999,
	}, nil
}
