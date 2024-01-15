package model

import "net/http"

type Response[T any] struct {
	Message    string `json:"message"`
	Successful bool   `json:"successful"`
	Data       T      `json:"data,omitempty"`
	Status     int    `json:"status"`
}

func NewOkResponse[T any](message string, data T) Response[T] {
	return Response[T]{
		Message:    message,
		Data:       data,
		Status:     http.StatusOK,
		Successful: true,
	}
}

func NewCreatedResponse[T any](message string, data T) Response[T] {
	return Response[T]{
		Message:    message,
		Data:       data,
		Status:     http.StatusCreated,
		Successful: true,
	}
}
