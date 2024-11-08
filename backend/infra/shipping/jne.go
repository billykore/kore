package shipping

import (
	"github.com/billykore/kore/backend/domain/order"
)

const JNECompanyName = "JNE Shipping"

// JNE shipping service.
type JNE struct {
	name string
}

// NewJNE return instance of JNE.
func NewJNE() *JNE {
	return &JNE{JNECompanyName}
}

// Ship order.
func (jne *JNE) Ship(data order.ShippingData) (*order.ShippingResponse, error) {
	return &order.ShippingResponse{
		Id:           699,
		Fee:          10000,
		Status:       "created",
		ShipperName:  jne.name,
		Address:      data.Address,
		CustomerName: data.CustomerName,
	}, nil
}
