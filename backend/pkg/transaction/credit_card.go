package transaction

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

func (cc *CreditCard) Pay(amount uint64) (*PaymentResponse, error) {
	return &PaymentResponse{
		Method:    creditCard,
		Amount:    amount,
		PaymentId: 999,
	}, nil
}
