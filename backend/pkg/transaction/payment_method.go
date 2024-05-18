package transaction

import "github.com/billykore/kore/backend/pkg/types"

type PaymentMethod interface {
	Pay(amount types.Money) (*PaymentResponse, error)
}

const (
	gopay      = "GoPay"
	creditCard = "Credit Card"
)

func NewPayment(method, name, account string) PaymentMethod {
	switch method {
	case gopay:
		return NewGoPay(name, account)
	case creditCard:
		return NewCreditCard(name, account)
	default:
		// Credit Card is the default payment method.
		return NewCreditCard(name, account)
	}
}

type PaymentResponse struct {
	Method    string
	Amount    types.Money
	PaymentId int
}
