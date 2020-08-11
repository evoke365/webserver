// Code generated by go-swagger; DO NOT EDIT.

package note

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UpdateNoteOKCode is the HTTP code returned for type UpdateNoteOK
const UpdateNoteOKCode int = 200

/*UpdateNoteOK Successful Operation

swagger:response updateNoteOK
*/
type UpdateNoteOK struct {
}

// NewUpdateNoteOK creates UpdateNoteOK with default headers values
func NewUpdateNoteOK() *UpdateNoteOK {

	return &UpdateNoteOK{}
}

// WriteResponse to the client
func (o *UpdateNoteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UpdateNoteBadRequestCode is the HTTP code returned for type UpdateNoteBadRequest
const UpdateNoteBadRequestCode int = 400

/*UpdateNoteBadRequest Invalid request data supplied

swagger:response updateNoteBadRequest
*/
type UpdateNoteBadRequest struct {
}

// NewUpdateNoteBadRequest creates UpdateNoteBadRequest with default headers values
func NewUpdateNoteBadRequest() *UpdateNoteBadRequest {

	return &UpdateNoteBadRequest{}
}

// WriteResponse to the client
func (o *UpdateNoteBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// UpdateNoteUnauthorizedCode is the HTTP code returned for type UpdateNoteUnauthorized
const UpdateNoteUnauthorizedCode int = 401

/*UpdateNoteUnauthorized Unauthorised user credentials

swagger:response updateNoteUnauthorized
*/
type UpdateNoteUnauthorized struct {
}

// NewUpdateNoteUnauthorized creates UpdateNoteUnauthorized with default headers values
func NewUpdateNoteUnauthorized() *UpdateNoteUnauthorized {

	return &UpdateNoteUnauthorized{}
}

// WriteResponse to the client
func (o *UpdateNoteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// UpdateNoteInternalServerErrorCode is the HTTP code returned for type UpdateNoteInternalServerError
const UpdateNoteInternalServerErrorCode int = 500

/*UpdateNoteInternalServerError internal error

swagger:response updateNoteInternalServerError
*/
type UpdateNoteInternalServerError struct {
}

// NewUpdateNoteInternalServerError creates UpdateNoteInternalServerError with default headers values
func NewUpdateNoteInternalServerError() *UpdateNoteInternalServerError {

	return &UpdateNoteInternalServerError{}
}

// WriteResponse to the client
func (o *UpdateNoteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}