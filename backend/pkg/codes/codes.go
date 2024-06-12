package codes

// Code is enum to represents success or error codes.
type Code int

const (
	_ Code = iota // OK
	BadRequest
	Unauthenticated
	NotFound
	Internal
)
