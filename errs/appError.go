package errs

import "net/http"

type AppError struct {
	Code    int
	Status  string
	Message string
}

func NotFoundError(message string, status string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Status:  status,
		Message: message,
	}
}

func UnExpectedError(message string, status string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Status:  status,
		Message: message,
	}
}
