package model

import (
	"strconv"
	"strings"

	"github.com/billykore/kore/backend/pkg/types"
	"gorm.io/gorm"
)

type OrderStatus string

const (
	OrderStatusCreated            OrderStatus = "created"
	OrderStatusWaitingForPayment  OrderStatus = "waiting_for_payment"
	OrderStatusPaymentSucceed     OrderStatus = "payment_succeed"
	OrderStatusWaitingForShipment OrderStatus = "waiting_for_shipment"
	OrderStatusCancelled          OrderStatus = "cancelled"
)

var OrderStatusCanCancel = []OrderStatus{OrderStatusCreated, OrderStatusWaitingForPayment}

func (status OrderStatus) String() string {
	return string(status)
}

type Order struct {
	gorm.Model
	Username      string
	PaymentMethod string
	CartIds       string
	Carts         []Cart
	Status        OrderStatus
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
