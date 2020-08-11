// Code generated by go-swagger; DO NOT EDIT.

package note

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetNotesHandlerFunc turns a function with the right signature into a get notes handler
type GetNotesHandlerFunc func(GetNotesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetNotesHandlerFunc) Handle(params GetNotesParams) middleware.Responder {
	return fn(params)
}

// GetNotesHandler interface for that can handle valid get notes params
type GetNotesHandler interface {
	Handle(GetNotesParams) middleware.Responder
}

// NewGetNotes creates a new http.Handler for the get notes operation
func NewGetNotes(ctx *middleware.Context, handler GetNotesHandler) *GetNotes {
	return &GetNotes{Context: ctx, Handler: handler}
}

/*GetNotes swagger:route GET /notes/{userId} note getNotes

Get Notes of a user

Retrieve notes by User ID

*/
type GetNotes struct {
	Context *middleware.Context
	Handler GetNotesHandler
}

func (o *GetNotes) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetNotesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}