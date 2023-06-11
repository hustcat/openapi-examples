// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/hustcat/openapi-examples/petstore-v2/models"
)

// FindPetsOKCode is the HTTP code returned for type FindPetsOK
const FindPetsOKCode int = 200

/*FindPetsOK pet response

swagger:response findPetsOK
*/
type FindPetsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Pet `json:"body,omitempty"`
}

// NewFindPetsOK creates FindPetsOK with default headers values
func NewFindPetsOK() *FindPetsOK {

	return &FindPetsOK{}
}

// WithPayload adds the payload to the find pets o k response
func (o *FindPetsOK) WithPayload(payload []*models.Pet) *FindPetsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find pets o k response
func (o *FindPetsOK) SetPayload(payload []*models.Pet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindPetsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Pet, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*FindPetsDefault unexpected error

swagger:response findPetsDefault
*/
type FindPetsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewFindPetsDefault creates FindPetsDefault with default headers values
func NewFindPetsDefault(code int) *FindPetsDefault {
	if code <= 0 {
		code = 500
	}

	return &FindPetsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find pets default response
func (o *FindPetsDefault) WithStatusCode(code int) *FindPetsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the find pets default response
func (o *FindPetsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the find pets default response
func (o *FindPetsDefault) WithPayload(payload *models.ErrorModel) *FindPetsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find pets default response
func (o *FindPetsDefault) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindPetsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
