package transaction

import "github.com/billykore/kore/backend/pkg/types"

// PaymentMethod is interface to represents payment service.
type PaymentMethod interface {
	Pay(amount types.Money) (*PaymentResponse, error)
}

const (
	gopay      = "GoPay"
	creditCard = "Credit Card"
)

// NewPayment return instance or PaymentMethod.
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
