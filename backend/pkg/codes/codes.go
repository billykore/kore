package codes

// Code is enum to represents success or error codes.
type Code int

const (
	// OK is success code.
	_ Code = iota

	// BadRequest is code to represents bad request error.
	BadRequest

	// Unauthenticated is code to represents unauthenticated error.
	Unauthenticated

	// NotFound is code to represents not found error.
	NotFound

	// Conflict is code to represent conflict error.
	Conflict

	// Internal is code to represents internal server error.
	Internal
)
