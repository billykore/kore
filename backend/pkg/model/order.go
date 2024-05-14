package model

import (
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type OrderStatus string

const (
	OrderStatusCreated            OrderStatus = "created"
	OrderStatusCancelled          OrderStatus = "cancelled"
	OrderStatusWaitingForPayment  OrderStatus = "waiting_for_payment"
	OrderStatusPaymentSucceed     OrderStatus = "payment_succeed"
	OrderStatusWaitingForShipment OrderStatus = "waiting_for_shipment"
)

func (status OrderStatus) String() string {
	return string(status)
}

type Order struct {
	gorm.Model
	UserId        int
	PaymentMethod string
	CartIds       string
	Carts         []Cart
	Status        OrderStatus
	ShippingId    int
}

// TotalPrice calculate total price of the items in one order.
func (o *Order) TotalPrice() uint64 {
	var total uint64
	for _, c := range o.Carts {
		total += c.Product.Price
	}
	return total
}

func (o *Order) IntCartIds() []int {
	// this is stupid hack.
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
