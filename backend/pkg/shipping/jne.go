package shipping

const JNECompanyName = "JNE Shipping"

type JNE struct {
	name string
}

func NewJNE() *JNE {
	return &JNE{JNECompanyName}
}

func (jne *JNE) Create(data *Data) (*Response, error) {
	return &Response{
		Id:           699,
		Fee:          10000,
		Status:       "created",
		ShipperName:  jne.name,
		Address:      data.Address,
		CustomerName: data.CustomerName,
	}, nil
}
