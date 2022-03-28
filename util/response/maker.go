package response

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// NewStatusOK is responsible for creating a 'ok' response.
func NewStatusOK(data any) (int, respBody) {
	resp := respBody{
		Status: http.StatusOK,
		Data:   data,
	}

	return resp.Status, resp
}

// NewStatusCreated is responsible for creating a 'created' response.
func NewStatusCreated(data any) (int, respBody) {
	resp := respBody{
		Status: http.StatusCreated,
		Data:   data,
	}

	return resp.Status, resp
}

// NewStatusBadRequest is responsible for creating a 'bad request' response.
func NewStatusBadRequest(msg string) (int, respBody) {
	resp := respBody{
		Status: http.StatusBadRequest,
		Data:   msg,
	}

	return resp.Status, resp
}

// NewStatusBadRequest is responsible for creating a 'not found' response.
func NewStatusNotFound(msg string) (int, respBody) {
	resp := respBody{
		Status: http.StatusNotFound,
		Data:   msg,
	}

	return resp.Status, resp
}

// StatusUnprocessableEntity is responsible for creating a 'unprocessable entity' response.
func NewStatusUnprocessableEntity(err error) (int, respBody) {
	data := make(map[string]string, 0)
	for _, err := range err.(validator.ValidationErrors) {
		msg := fmt.Sprintf("validation error on '%s' rule", err.Tag())
		data[err.Field()] = msg
	}

	resp := respBody{
		Status: http.StatusUnprocessableEntity,
		Data:   data,
	}

	return resp.Status, resp
}

// NewStatusInternalServerError is responsible for creating a 'internal server error' response.
func NewStatusInternalServerError(msg string) (int, respBody) {
	resp := respBody{
		Status: http.StatusInternalServerError,
		Data:   msg,
	}

	return resp.Status, resp
}
