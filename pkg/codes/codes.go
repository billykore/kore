package codes

type Code int

const (
	OK Code = iota
	BadRequest
	Unauthenticated
	NotFound
	Internal
)
