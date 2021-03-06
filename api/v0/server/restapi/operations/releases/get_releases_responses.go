// Code generated by go-swagger; DO NOT EDIT.

package releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/optikon/api/api/v0/models"
)

// GetReleasesOKCode is the HTTP code returned for type GetReleasesOK
const GetReleasesOKCode int = 200

/*GetReleasesOK OK

swagger:response getReleasesOK
*/
type GetReleasesOK struct {

	/*
	  In: Body
	*/
	Payload models.GetReleasesOKBody `json:"body,omitempty"`
}

// NewGetReleasesOK creates GetReleasesOK with default headers values
func NewGetReleasesOK() *GetReleasesOK {
	return &GetReleasesOK{}
}

// WithPayload adds the payload to the get releases o k response
func (o *GetReleasesOK) WithPayload(payload models.GetReleasesOKBody) *GetReleasesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get releases o k response
func (o *GetReleasesOK) SetPayload(payload models.GetReleasesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReleasesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.GetReleasesOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetReleasesBadRequestCode is the HTTP code returned for type GetReleasesBadRequest
const GetReleasesBadRequestCode int = 400

/*GetReleasesBadRequest Bad Request

swagger:response getReleasesBadRequest
*/
type GetReleasesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetReleasesBadRequest creates GetReleasesBadRequest with default headers values
func NewGetReleasesBadRequest() *GetReleasesBadRequest {
	return &GetReleasesBadRequest{}
}

// WithPayload adds the payload to the get releases bad request response
func (o *GetReleasesBadRequest) WithPayload(payload *models.APIResponse) *GetReleasesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get releases bad request response
func (o *GetReleasesBadRequest) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReleasesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReleasesUnauthorizedCode is the HTTP code returned for type GetReleasesUnauthorized
const GetReleasesUnauthorizedCode int = 401

/*GetReleasesUnauthorized Unauthorized

swagger:response getReleasesUnauthorized
*/
type GetReleasesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetReleasesUnauthorized creates GetReleasesUnauthorized with default headers values
func NewGetReleasesUnauthorized() *GetReleasesUnauthorized {
	return &GetReleasesUnauthorized{}
}

// WithPayload adds the payload to the get releases unauthorized response
func (o *GetReleasesUnauthorized) WithPayload(payload *models.APIResponse) *GetReleasesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get releases unauthorized response
func (o *GetReleasesUnauthorized) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReleasesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReleasesInternalServerErrorCode is the HTTP code returned for type GetReleasesInternalServerError
const GetReleasesInternalServerErrorCode int = 500

/*GetReleasesInternalServerError Internal Server Error

swagger:response getReleasesInternalServerError
*/
type GetReleasesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetReleasesInternalServerError creates GetReleasesInternalServerError with default headers values
func NewGetReleasesInternalServerError() *GetReleasesInternalServerError {
	return &GetReleasesInternalServerError{}
}

// WithPayload adds the payload to the get releases internal server error response
func (o *GetReleasesInternalServerError) WithPayload(payload *models.APIResponse) *GetReleasesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get releases internal server error response
func (o *GetReleasesInternalServerError) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReleasesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
