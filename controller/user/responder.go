package user

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

type DefaultResponderOK struct {
}

func (dr *DefaultResponderOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(http.StatusOK)
}

type DefaultResponderNoContent struct {
}

func (dr *DefaultResponderNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(http.StatusNoContent)
}

type DefaultResponderBadRequest struct {
}

func (dr *DefaultResponderBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(http.StatusBadRequest)
}

type DefaultResponderError struct {
}

func (dr *DefaultResponderError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(http.StatusInternalServerError)
}
