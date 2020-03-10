// Code generated by go-swagger; DO NOT EDIT.

package create_table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCreateTableParams creates a new CreateTableParams object
// no default values defined in spec.
func NewCreateTableParams() CreateTableParams {

	return CreateTableParams{}
}

// CreateTableParams contains all the bound params for the create table operation
// typically these are obtained from a http.Request
//
// swagger:parameters CreateTable
type CreateTableParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Type of the table that have to be created.
	  Required: true
	  In: query
	*/
	TableType int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateTableParams() beforehand.
func (o *CreateTableParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qTableType, qhkTableType, _ := qs.GetOK("TableType")
	if err := o.bindTableType(qTableType, qhkTableType, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindTableType binds and validates parameter TableType from query.
func (o *CreateTableParams) bindTableType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("TableType", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("TableType", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("TableType", "query", "int64", raw)
	}
	o.TableType = value

	return nil
}
