package transaction

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

func (p *GoPay) Pay(amount uint64) (*PaymentResponse, error) {
	return &PaymentResponse{
		Method:    gopay,
		Amount:    amount,
		PaymentId: 999,
	}, nil
}
