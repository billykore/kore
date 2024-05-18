package entity

import (
	"errors"
	"net/http"
	"time"

	"github.com/billykore/kore/backend/pkg/status"
)

type Data map[string]any

func makeData(fieldName string, fieldValue any) Data {
	data := make(Data)
	data[fieldName] = fieldValue
	return data
}

type Response struct {
	Status     string `json:"status,omitempty"`
	Message    string `json:"message,omitempty"`
	Data       Data   `json:"data,omitempty"`
	ServerTime int64  `json:"serverTime,omitempty"`
}

func ResponseSuccess(fieldName string, fieldValue any) (int, *Response) {
	return http.StatusOK, &Response{
		Status:     "OK",
		Data:       makeData(fieldName, fieldValue),
		ServerTime: time.Now().Unix(),
	}
}

func ResponseSuccessNilData() (int, *Response) {
	return http.StatusOK, &Response{
		Status:     "OK",
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

func ResponseBadRequest(err error) (int, *Response) {
	return http.StatusInternalServerError, &Response{
		Status:     "BAD_REQUEST",
		Message:    err.Error(),
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
