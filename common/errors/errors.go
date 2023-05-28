package errors

import "net/http"

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *Error {
	return &Error{
		Message: message,
		Status:  http.StatusBadRequest,
		Error: "bad_request",
	}
}

func NewNotFoundError(message string) *Error {
	return &Error{
		Message: message,
		Status:  http.StatusNotFound,
		Error: "not_found",
	}
}

func NewInternalServerError(message string) *Error {
	return &Error{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error: "internal_server_error",
	}
}

func NewStatusConflictError(message string) *Error {
	return &Error{
		Message: message,
		Status: http.StatusConflict,
		Error: "conflict_error",
	}
}

func NewStatusForbidden(message string) *Error {
	return &Error{
		Message: message,
		Status: http.StatusForbidden,
		Error: "forbidden",
	}
}
