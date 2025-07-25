package dto

import "net/http"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func Success(message string, data interface{}) Response {
	return Response{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	}
}

func Created(message string, data interface{}) Response {
	return Response{
		Status:  http.StatusCreated,
		Message: message,
		Data:    data,
	}
}

func BadRequest(err interface{}) Response {
	return Response{
		Status:  http.StatusBadRequest,
		Message: "Bad Request",
		Error:   err,
	}
}

func Unauthorized(err interface{}) Response {
	return Response{
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized",
		Error:   err,
	}
}

func InternalError(err interface{}) Response {
	return Response{
		Status:  http.StatusInternalServerError,
		Message: "Internal Server Error",
		Error:   err,
	}
}
