package order

import (
	"strconv"
	"strings"

	"github.com/billykore/kore/backend/pkg/types"
	"gorm.io/gorm"
)

type Status string

const (
	StatusCreated            Status = "created"
	StatusWaitingForPayment  Status = "waiting_for_payment"
	StatusPaymentSucceed     Status = "payment_succeed"
	StatusWaitingForShipment Status = "waiting_for_shipment"
	StatusCancelled          Status = "cancelled"
)

var StatusCanCancel = []Status{StatusCreated, StatusWaitingForPayment}

func (status Status) String() string {
	return string(status)
}

type Cart struct {
	gorm.Model
	Username  string
	ProductId int
	Quantity  int
	Product   Product `gorm:"foreignKey:ProductId"`
}

type Order struct {
	gorm.Model
	Username      string
	PaymentMethod string
	CartIds       string
	Carts         []Cart
	Status        Status
	ShippingId    int
}

func (o *Order) SetCartIds(ids []uint) {
	// this is stupid hack.
	if ids != nil {
		s := "{"
		for i, id := range ids {
			s += strconv.Itoa(int(id))
			if i != len(ids)-1 {
				s += ","
			}
		}
		s += "}"
		o.CartIds = s
	}
}

// TotalPrice calculate total price of the items in one order.
func (o *Order) TotalPrice() types.Money {
	var total types.Money
	for _, c := range o.Carts {
		total += c.Product.Price
	}
	return total
}

func (o *Order) IntCartIds() []int {
	// again, this is stupid hack.
	var ids []int
	s := strings.ReplaceAll(o.CartIds, "{", "")
	s = strings.ReplaceAll(s, "}", "")
	for _, v := range strings.Split(s, ",") {
		id, err := strconv.Atoi(v)
		if err != nil {
			ids = append(ids, 0)
		}
		ids = append(ids, id)
	}
	return ids
}

type Product struct {
	gorm.Model
	Name             string
	Description      string
	SKU              string
	Price            types.Money
	CategoryId       int
	InventoryId      int
	DiscountId       int
	ProductCategory  ProductCategory  `gorm:"foreignKey:CategoryId"`
	ProductInventory ProductInventory `gorm:"foreignKey:InventoryId"`
	Discount         Discount         `gorm:"foreignKey:DiscountId"`
}

type ProductCategory struct {
	gorm.Model
	Name        string
	Description string
}

type ProductInventory struct {
	gorm.Model
	Quantity    int
	Description string
}

type Discount struct {
	gorm.Model
	Name               string
	Description        string
	DiscountPercentage float64
	Active             bool
}
