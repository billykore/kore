package transaction

type PaymentMethod interface {
	Pay(amount uint64) (*PaymentResponse, error)
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
	Amount    uint64
	PaymentId int
}
