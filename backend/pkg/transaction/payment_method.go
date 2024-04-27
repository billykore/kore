package transaction

type PaymentMethod[T any] interface {
	Pay(amount uint64) (T, error)
}
