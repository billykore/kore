package entity

import (
	"errors"
	"net/http"
	"time"

	"github.com/billykore/kore/backend/pkg/status"
)

type Response struct {
	Status     string `json:"status,omitempty"`
	Data       any    `json:"data,omitempty"`
	Message    string `json:"message,omitempty"`
	ServerTime int64  `json:"serverTime,omitempty"`
}

func ResponseSuccess(data any) (int, *Response) {
	return http.StatusOK, &Response{
		Status:     "OK",
		Data:       data,
		ServerTime: time.Now().Unix(),
	}
}

func ResponseError(err error) (int, *Response) {
	var s *status.Status
	errors.As(err, &s)
	return responseCode[s.Code], &Response{
		Status:     responseStatus[s.Code],
		Message:    s.Message,
		ServerTime: time.Now().Unix(),
	}
}

var responseCode = []int{
	http.StatusOK,
	http.StatusBadRequest,
	http.StatusUnauthorized,
	http.StatusNotFound,
	http.StatusInternalServerError,
}

var responseStatus = []string{
	"OK",
	"BAD_REQUEST",
	"UNAUTHORIZED",
	"NOT_FOUND",
	"INTERNAL_SERVER_ERROR",
}
