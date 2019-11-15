// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// StorageReadFileHandlerFunc turns a function with the right signature into a storage read file handler
type StorageReadFileHandlerFunc func(StorageReadFileParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn StorageReadFileHandlerFunc) Handle(params StorageReadFileParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// StorageReadFileHandler interface for that can handle valid storage read file params
type StorageReadFileHandler interface {
	Handle(StorageReadFileParams, interface{}) middleware.Responder
}

// NewStorageReadFile creates a new http.Handler for the storage read file operation
func NewStorageReadFile(ctx *middleware.Context, handler StorageReadFileHandler) *StorageReadFile {
	return &StorageReadFile{Context: ctx, Handler: handler}
}

/*StorageReadFile swagger:route GET /readFile Storage storageReadFile

Download a file

Streaming a file contents to the client

*/
type StorageReadFile struct {
	Context *middleware.Context
	Handler StorageReadFileHandler
}

func (o *StorageReadFile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewStorageReadFileParams()

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
