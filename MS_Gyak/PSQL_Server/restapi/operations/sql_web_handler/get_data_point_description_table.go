// Code generated by go-swagger; DO NOT EDIT.

package sql_web_handler

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetDataPointDescriptionTableHandlerFunc turns a function with the right signature into a get data point description table handler
type GetDataPointDescriptionTableHandlerFunc func(GetDataPointDescriptionTableParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDataPointDescriptionTableHandlerFunc) Handle(params GetDataPointDescriptionTableParams) middleware.Responder {
	return fn(params)
}

// GetDataPointDescriptionTableHandler interface for that can handle valid get data point description table params
type GetDataPointDescriptionTableHandler interface {
	Handle(GetDataPointDescriptionTableParams) middleware.Responder
}

// NewGetDataPointDescriptionTable creates a new http.Handler for the get data point description table operation
func NewGetDataPointDescriptionTable(ctx *middleware.Context, handler GetDataPointDescriptionTableHandler) *GetDataPointDescriptionTable {
	return &GetDataPointDescriptionTable{Context: ctx, Handler: handler}
}

/*GetDataPointDescriptionTable swagger:route GET /GetDataPointDescriptionTable SQLWebHandler getDataPointDescriptionTable

Get the specified table from the SQL server.

*/
type GetDataPointDescriptionTable struct {
	Context *middleware.Context
	Handler GetDataPointDescriptionTableHandler
}

func (o *GetDataPointDescriptionTable) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetDataPointDescriptionTableParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
