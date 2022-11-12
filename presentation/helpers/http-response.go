package helpers

import (
	"errors"
	"net/http"
)

type HttpResponse struct {
	StatusCode int
	Detail     string
	Body       interface{}
}

func NewHttpResponse() (h HttpResponse) {
	return
}

func (h *HttpResponse) BadRequest(detail string) {
	h.StatusCode = http.StatusBadRequest
	h.Detail = errors.New(detail).Error()
}

func (h *HttpResponse) UnauthorizedError(detail string) {
	// h.StatusCode = 401
	h.StatusCode = http.StatusUnauthorized
	h.Detail = errors.New(detail).Error()
}

func (h *HttpResponse) ServerError(detail string) {
	// h.StatusCode = 500
	h.StatusCode = http.StatusInternalServerError
	h.Detail = errors.New(detail).Error()
}
