// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// FindUserOKCode is the HTTP code returned for type FindUserOK
const FindUserOKCode int = 200

/*FindUserOK successful operation

swagger:response findUserOK
*/
type FindUserOK struct {
}

// NewFindUserOK creates FindUserOK with default headers values
func NewFindUserOK() *FindUserOK {

	return &FindUserOK{}
}

// WriteResponse to the client
func (o *FindUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// FindUserNotFoundCode is the HTTP code returned for type FindUserNotFound
const FindUserNotFoundCode int = 404

/*FindUserNotFound user not found

swagger:response findUserNotFound
*/
type FindUserNotFound struct {
}

// NewFindUserNotFound creates FindUserNotFound with default headers values
func NewFindUserNotFound() *FindUserNotFound {

	return &FindUserNotFound{}
}

// WriteResponse to the client
func (o *FindUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
