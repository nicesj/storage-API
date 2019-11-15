// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewStorageReadFileParams creates a new StorageReadFileParams object
// no default values defined in spec.
func NewStorageReadFileParams() StorageReadFileParams {

	return StorageReadFileParams{}
}

// StorageReadFileParams contains all the bound params for the storage read file operation
// typically these are obtained from a http.Request
//
// swagger:parameters storageReadFile
type StorageReadFileParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A file what would be transferred
	  In: query
	*/
	F *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewStorageReadFileParams() beforehand.
func (o *StorageReadFileParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qF, qhkF, _ := qs.GetOK("f")
	if err := o.bindF(qF, qhkF, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindF binds and validates parameter F from query.
func (o *StorageReadFileParams) bindF(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.F = &raw

	return nil
}
