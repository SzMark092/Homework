// Code generated by go-swagger; DO NOT EDIT.

package create_table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// CreateTableMethodNotAllowedCode is the HTTP code returned for type CreateTableMethodNotAllowed
const CreateTableMethodNotAllowedCode int = 405

/*CreateTableMethodNotAllowed Invalid input

swagger:response createTableMethodNotAllowed
*/
type CreateTableMethodNotAllowed struct {
}

// NewCreateTableMethodNotAllowed creates CreateTableMethodNotAllowed with default headers values
func NewCreateTableMethodNotAllowed() *CreateTableMethodNotAllowed {

	return &CreateTableMethodNotAllowed{}
}

// WriteResponse to the client
func (o *CreateTableMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}
