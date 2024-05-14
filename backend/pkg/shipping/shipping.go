package shipping

type Shipping interface {
	Create(data *Data) (*Response, error)
}

const (
	jne = "JNE"
)

func New(shipperName string) Shipping {
	switch shipperName {
	case jne:
		return NewJNE()
	default:
		// JNE is the default shipper.
		return NewJNE()
	}
}

type Data struct {
	Address      string
	CustomerName string
	ShippingType string
}

type Response struct {
	Id           int
	Fee          uint64
	Status       string
	ShipperName  string
	Address      string
	CustomerName string
}
