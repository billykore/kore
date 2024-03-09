package status

import "github.com/billykore/kore/backend/pkg/codes"

type Status struct {
	Code    codes.Code
	Message string
}

func Error(c codes.Code, msg string) error {
	return &Status{
		Code:    c,
		Message: msg,
	}
}

func (s *Status) Error() string {
	return s.Message
}
