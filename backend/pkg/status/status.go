package status

import "github.com/billykore/kore/backend/pkg/codes"

// Status represent error throw by services.
type Status struct {
	// Code is the error code.
	Code codes.Code
	// Message is the error message.
	Message string
}

// Error return new Status.
func Error(c codes.Code, msg string) error {
	return &Status{
		Code:    c,
		Message: msg,
	}
}

func (s *Status) Error() string {
	return s.Message
}
