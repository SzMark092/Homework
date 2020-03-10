// Code generated by go-swagger; DO NOT EDIT.

package create_table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateTableHandlerFunc turns a function with the right signature into a create table handler
type CreateTableHandlerFunc func(CreateTableParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateTableHandlerFunc) Handle(params CreateTableParams) middleware.Responder {
	return fn(params)
}

// CreateTableHandler interface for that can handle valid create table params
type CreateTableHandler interface {
	Handle(CreateTableParams) middleware.Responder
}

// NewCreateTable creates a new http.Handler for the create table operation
func NewCreateTable(ctx *middleware.Context, handler CreateTableHandler) *CreateTable {
	return &CreateTable{Context: ctx, Handler: handler}
}

/*CreateTable swagger:route POST /CreateTable CreateTable createTable

Create a table with the given code.

*/
type CreateTable struct {
	Context *middleware.Context
	Handler CreateTableHandler
}

func (o *CreateTable) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateTableParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
