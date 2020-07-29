// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// LoginUserOKCode is the HTTP code returned for type LoginUserOK
const LoginUserOKCode int = 200

/*LoginUserOK successful operation

swagger:response loginUserOK
*/
type LoginUserOK struct {
}

// NewLoginUserOK creates LoginUserOK with default headers values
func NewLoginUserOK() *LoginUserOK {

	return &LoginUserOK{}
}

// WriteResponse to the client
func (o *LoginUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// LoginUserBadRequestCode is the HTTP code returned for type LoginUserBadRequest
const LoginUserBadRequestCode int = 400

/*LoginUserBadRequest Invalid username/password supplied

swagger:response loginUserBadRequest
*/
type LoginUserBadRequest struct {
}

// NewLoginUserBadRequest creates LoginUserBadRequest with default headers values
func NewLoginUserBadRequest() *LoginUserBadRequest {

	return &LoginUserBadRequest{}
}

// WriteResponse to the client
func (o *LoginUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// LoginUserUnauthorizedCode is the HTTP code returned for type LoginUserUnauthorized
const LoginUserUnauthorizedCode int = 401

/*LoginUserUnauthorized Unauthorised user credentials

swagger:response loginUserUnauthorized
*/
type LoginUserUnauthorized struct {
}

// NewLoginUserUnauthorized creates LoginUserUnauthorized with default headers values
func NewLoginUserUnauthorized() *LoginUserUnauthorized {

	return &LoginUserUnauthorized{}
}

// WriteResponse to the client
func (o *LoginUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// LoginUserInternalServerErrorCode is the HTTP code returned for type LoginUserInternalServerError
const LoginUserInternalServerErrorCode int = 500

/*LoginUserInternalServerError internal error

swagger:response loginUserInternalServerError
*/
type LoginUserInternalServerError struct {
}

// NewLoginUserInternalServerError creates LoginUserInternalServerError with default headers values
func NewLoginUserInternalServerError() *LoginUserInternalServerError {

	return &LoginUserInternalServerError{}
}

// WriteResponse to the client
func (o *LoginUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
