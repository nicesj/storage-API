// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ReadFolderHandlerFunc turns a function with the right signature into a read folder handler
type ReadFolderHandlerFunc func(ReadFolderParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadFolderHandlerFunc) Handle(params ReadFolderParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReadFolderHandler interface for that can handle valid read folder params
type ReadFolderHandler interface {
	Handle(ReadFolderParams, interface{}) middleware.Responder
}

// NewReadFolder creates a new http.Handler for the read folder operation
func NewReadFolder(ctx *middleware.Context, handler ReadFolderHandler) *ReadFolder {
	return &ReadFolder{Context: ctx, Handler: handler}
}

/*ReadFolder swagger:route GET /readFolder Storage readFolder

Get a list of files and folders

Get a list of files and folders in a given path (folder)

*/
type ReadFolder struct {
	Context *middleware.Context
	Handler ReadFolderHandler
}

func (o *ReadFolder) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadFolderParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
