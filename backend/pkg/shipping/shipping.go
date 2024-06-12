package shipping

import "github.com/billykore/kore/backend/pkg/types"

// Shipping interface define the shipping services behavior.
type Shipping interface {
	Create(data *Data) (*Response, error)
}

const (
	jne = "JNE"
)

// New returns instance of Shipping service.
func New(shipperName string) Shipping {
	switch shipperName {
	case jne:
		return NewJNE()
	default:
		// JNE is the default shipper.
		return NewJNE()
	}
}

// Data for the order.
type Data struct {
	Address      string
	CustomerName string
	ShippingType string
}

type Response struct {
	Id           int
	Fee          types.Money
	Status       string
	ShipperName  string
	Address      string
	CustomerName string
}
