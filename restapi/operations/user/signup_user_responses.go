// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// SignupUserOKCode is the HTTP code returned for type SignupUserOK
const SignupUserOKCode int = 200

/*SignupUserOK successful operation

swagger:response signupUserOK
*/
type SignupUserOK struct {
}

// NewSignupUserOK creates SignupUserOK with default headers values
func NewSignupUserOK() *SignupUserOK {

	return &SignupUserOK{}
}

// WriteResponse to the client
func (o *SignupUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}