package responder

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

type OK struct {
	response []byte
}

func (r *OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(http.StatusOK)
	if len(r.response) > 0 {
		rw.Write(r.response)
	}
}

func (r *OK) WithResponse(response []byte) *OK {
	r.response = response
	return r
}

func DefaultOK() *OK {
	return &OK{}
}

type NoContent struct {
}

func (r *NoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(http.StatusNoContent)
}

func DefaultNoContent() *NoContent {
	return &NoContent{}
}

type BadRequest struct {
}

func (r *BadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(http.StatusBadRequest)
}

func DefaultBadRequest() *BadRequest {
	return &BadRequest{}
}

type Unauthorised struct {
}

func (r *Unauthorised) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(http.StatusUnauthorized)
}

func DefaultUnauthorised() *Unauthorised {
	return &Unauthorised{}
}

type ServerError struct {
}

func (r *ServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(http.StatusInternalServerError)
}

func DefaultServerError() *ServerError {
	return &ServerError{}
}
