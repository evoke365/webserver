// Code generated by go-swagger; DO NOT EDIT.

package note

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddNoteHandlerFunc turns a function with the right signature into a add note handler
type AddNoteHandlerFunc func(AddNoteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddNoteHandlerFunc) Handle(params AddNoteParams) middleware.Responder {
	return fn(params)
}

// AddNoteHandler interface for that can handle valid add note params
type AddNoteHandler interface {
	Handle(AddNoteParams) middleware.Responder
}

// NewAddNote creates a new http.Handler for the add note operation
func NewAddNote(ctx *middleware.Context, handler AddNoteHandler) *AddNote {
	return &AddNote{Context: ctx, Handler: handler}
}

/*AddNote swagger:route POST /note note addNote

user creates a new note

user creates a new note

*/
type AddNote struct {
	Context *middleware.Context
	Handler AddNoteHandler
}

func (o *AddNote) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddNoteParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
