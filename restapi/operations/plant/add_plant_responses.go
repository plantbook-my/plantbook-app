// Code generated by go-swagger; DO NOT EDIT.

package plant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AddPlantMethodNotAllowedCode is the HTTP code returned for type AddPlantMethodNotAllowed
const AddPlantMethodNotAllowedCode int = 405

/*AddPlantMethodNotAllowed Invalid input

swagger:response addPlantMethodNotAllowed
*/
type AddPlantMethodNotAllowed struct {
}

// NewAddPlantMethodNotAllowed creates AddPlantMethodNotAllowed with default headers values
func NewAddPlantMethodNotAllowed() *AddPlantMethodNotAllowed {

	return &AddPlantMethodNotAllowed{}
}

// WriteResponse to the client
func (o *AddPlantMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}
