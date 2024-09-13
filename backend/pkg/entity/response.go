package entity

import (
	"errors"
	"net/http"
	"time"

	"github.com/billykore/kore/backend/pkg/status"
)

const serverTimeFormat = "2006-01-02T15:04:05"

type Response struct {
	Status     string `json:"status,omitempty"`
	Message    string `json:"message,omitempty"`
	Data       any    `json:"data,omitempty"`
	ServerTime string `json:"serverTime,omitempty"`
}

// ResponseSuccess returns status code 200 and success response with data.
func ResponseSuccess(data any) (int, *Response) {
	return http.StatusOK, &Response{
		Status:     "OK",
		Data:       data,
		ServerTime: time.Now().Format(serverTimeFormat),
	}
}

// ResponseSuccessNilData returns status code 200 and success response without data.
func ResponseSuccessNilData() (int, *Response) {
	return http.StatusOK, &Response{
		Status:     "OK",
		ServerTime: time.Now().Format(serverTimeFormat),
	}
}

// ResponseError returns error status code and error.
func ResponseError(err error) (int, *Response) {
	var s *status.Status
	errors.As(err, &s)
	return responseCode[s.Code], &Response{
		Status:     responseStatus[s.Code],
		Message:    s.Message,
		ServerTime: time.Now().Format(serverTimeFormat),
	}
}

// ResponseBadRequest returns status code 400 and error response.
func ResponseBadRequest(err error) (int, *Response) {
	return http.StatusBadRequest, &Response{
		Status:     "BAD_REQUEST",
		Message:    err.Error(),
		ServerTime: time.Now().Format(serverTimeFormat),
	}
}

// ResponseUnauthorized returns status code 401 and error response.
func ResponseUnauthorized(err error) (int, *Response) {
	return http.StatusUnauthorized, &Response{
		Status:     "UNAUTHORIZED",
		Message:    err.Error(),
		ServerTime: time.Now().Format(serverTimeFormat),
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
