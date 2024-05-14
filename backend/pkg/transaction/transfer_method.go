package transaction

import "context"

type TransferMethod[T, U any] interface {
	Inquiry(ctx context.Context) (T, error)
	Payment(ctx context.Context) (U, error)
}
