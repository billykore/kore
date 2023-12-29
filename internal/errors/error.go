package errors

type ErrType int

const (
	TypeInternalServerError ErrType = iota
	TypeBadRequest
	TypeNotFound
)

type Error struct {
	Type    ErrType `json:"-"`
	Message string  `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}
