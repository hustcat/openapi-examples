// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddPetHandlerFunc turns a function with the right signature into a add pet handler
type AddPetHandlerFunc func(AddPetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddPetHandlerFunc) Handle(params AddPetParams) middleware.Responder {
	return fn(params)
}

// AddPetHandler interface for that can handle valid add pet params
type AddPetHandler interface {
	Handle(AddPetParams) middleware.Responder
}

// NewAddPet creates a new http.Handler for the add pet operation
func NewAddPet(ctx *middleware.Context, handler AddPetHandler) *AddPet {
	return &AddPet{Context: ctx, Handler: handler}
}

/* AddPet swagger:route POST /pets addPet

Creates a new pet in the store.  Duplicates are allowed

*/
type AddPet struct {
	Context *middleware.Context
	Handler AddPetHandler
}

func (o *AddPet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAddPetParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
