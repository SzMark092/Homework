// Code generated by go-swagger; DO NOT EDIT.

package sql_web_handler

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/SzMark092/MS_Gyak/PSQL_Server/models"
)

// GetModuleOKCode is the HTTP code returned for type GetModuleOK
const GetModuleOKCode int = 200

/*GetModuleOK JSON object containing table information

swagger:response getModuleOK
*/
type GetModuleOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Module `json:"body,omitempty"`
}

// NewGetModuleOK creates GetModuleOK with default headers values
func NewGetModuleOK() *GetModuleOK {

	return &GetModuleOK{}
}

// WithPayload adds the payload to the get module o k response
func (o *GetModuleOK) WithPayload(payload []*models.Module) *GetModuleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get module o k response
func (o *GetModuleOK) SetPayload(payload []*models.Module) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetModuleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Module, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetModuleNotFoundCode is the HTTP code returned for type GetModuleNotFound
const GetModuleNotFoundCode int = 404

/*GetModuleNotFound There is no data to send

swagger:response getModuleNotFound
*/
type GetModuleNotFound struct {
}

// NewGetModuleNotFound creates GetModuleNotFound with default headers values
func NewGetModuleNotFound() *GetModuleNotFound {

	return &GetModuleNotFound{}
}

// WriteResponse to the client
func (o *GetModuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
