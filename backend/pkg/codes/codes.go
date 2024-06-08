package codes

type Code int

const (
	_ Code = iota // OK
	BadRequest
	Unauthenticated
	NotFound
	Internal
)
