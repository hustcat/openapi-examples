// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/hustcat/openapi-examples/petstore-v2/models"
)

// AddPetOKCode is the HTTP code returned for type AddPetOK
const AddPetOKCode int = 200

/*AddPetOK pet response

swagger:response addPetOK
*/
type AddPetOK struct {

	/*
	  In: Body
	*/
	Payload *models.Pet `json:"body,omitempty"`
}

// NewAddPetOK creates AddPetOK with default headers values
func NewAddPetOK() *AddPetOK {

	return &AddPetOK{}
}

// WithPayload adds the payload to the add pet o k response
func (o *AddPetOK) WithPayload(payload *models.Pet) *AddPetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add pet o k response
func (o *AddPetOK) SetPayload(payload *models.Pet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddPetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AddPetDefault unexpected error

swagger:response addPetDefault
*/
type AddPetDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewAddPetDefault creates AddPetDefault with default headers values
func NewAddPetDefault(code int) *AddPetDefault {
	if code <= 0 {
		code = 500
	}

	return &AddPetDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add pet default response
func (o *AddPetDefault) WithStatusCode(code int) *AddPetDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add pet default response
func (o *AddPetDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add pet default response
func (o *AddPetDefault) WithPayload(payload *models.ErrorModel) *AddPetDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add pet default response
func (o *AddPetDefault) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddPetDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
