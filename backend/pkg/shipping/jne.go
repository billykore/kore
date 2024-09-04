package shipping

const JNECompanyName = "JNE Shipping"

// JNE shipping service.
type JNE struct {
	name string
}

// NewJNE return instance of JNE.
func NewJNE() *JNE {
	return &JNE{JNECompanyName}
}

// Create new JNE shipping.
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
