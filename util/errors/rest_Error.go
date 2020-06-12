package errors

import "net/http"

type Rest_Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *Rest_Error {
	return &Rest_Error{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *Rest_Error {
	return &Rest_Error{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}
func NewInternalServerError(message string) *Rest_Error {
	return &Rest_Error{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "INTERNAL_SERVER_ERROR",
	}
}
