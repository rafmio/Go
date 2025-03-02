// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetHelloOKCode is the HTTP code returned for type GetHelloOK
const GetHelloOKCode int = 200

/*
GetHelloOK Successful response

swagger:response getHelloOK
*/
type GetHelloOK struct {

	/*
	  In: Body
	*/
	Payload *GetHelloOKBody `json:"body,omitempty"`
}

// NewGetHelloOK creates GetHelloOK with default headers values
func NewGetHelloOK() *GetHelloOK {

	return &GetHelloOK{}
}

// WithPayload adds the payload to the get hello o k response
func (o *GetHelloOK) WithPayload(payload *GetHelloOKBody) *GetHelloOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get hello o k response
func (o *GetHelloOK) SetPayload(payload *GetHelloOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHelloOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
