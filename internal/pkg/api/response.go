package api

import (
	"errors"
	"net/http"

	customerror "github.com/billykore/todolist/internal/errors"
)

const (
	StatusInternalServerError = "INTERNAL_SERVER_ERROR"
	StatusSuccess             = "SUCCESS"
	StatusBadRequest          = "BAD_REQUEST"
	StatusNotFound            = "NOT_FOUND"
)

type Response struct {
	Status string `json:"status"`
	Error  error  `json:"error,omitempty"`
	Data   any    `json:"data,omitempty"`
}

func ResponseSuccess(data any) (int, *Response) {
	return http.StatusOK, &Response{
		Status: StatusSuccess,
		Data:   data,
	}
}

func ResponseError(err error) (int, *Response) {
	var e *customerror.Error
	if errors.As(err, &e) {
		switch e.Type {
		case customerror.TypeNotFound:
			return ResponseNotFound(e)
		case customerror.TypeInternalServerError:
			return ResponseInternalServerError(e)
		default:
			ResponseInternalServerError(e)
		}
	}
	return ResponseInternalServerError(e)
}

func ResponseBadRequest(err error) (int, *Response) {
	return http.StatusBadRequest, &Response{
		Status: StatusBadRequest,
		Error:  err,
	}
}

func ResponseNotFound(err error) (int, *Response) {
	return http.StatusNotFound, &Response{
		Status: StatusNotFound,
		Error:  err,
	}
}

func ResponseInternalServerError(err error) (int, *Response) {
	return http.StatusInternalServerError, &Response{
		Status: StatusInternalServerError,
		Error:  err,
	}
}
