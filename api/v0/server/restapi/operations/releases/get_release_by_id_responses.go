// Code generated by go-swagger; DO NOT EDIT.

package releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"wwwin-github.cisco.com/edge/optikon-api/api/v0/models"
)

// GetReleaseByIDOKCode is the HTTP code returned for type GetReleaseByIDOK
const GetReleaseByIDOKCode int = 200

/*GetReleaseByIDOK OK

swagger:response getReleaseByIdOK
*/
type GetReleaseByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.ReleaseRelease `json:"body,omitempty"`
}

// NewGetReleaseByIDOK creates GetReleaseByIDOK with default headers values
func NewGetReleaseByIDOK() *GetReleaseByIDOK {
	return &GetReleaseByIDOK{}
}

// WithPayload adds the payload to the get release by Id o k response
func (o *GetReleaseByIDOK) WithPayload(payload *models.ReleaseRelease) *GetReleaseByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get release by Id o k response
func (o *GetReleaseByIDOK) SetPayload(payload *models.ReleaseRelease) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReleaseByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReleaseByIDBadRequestCode is the HTTP code returned for type GetReleaseByIDBadRequest
const GetReleaseByIDBadRequestCode int = 400

/*GetReleaseByIDBadRequest Bad Request

swagger:response getReleaseByIdBadRequest
*/
type GetReleaseByIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetReleaseByIDBadRequest creates GetReleaseByIDBadRequest with default headers values
func NewGetReleaseByIDBadRequest() *GetReleaseByIDBadRequest {
	return &GetReleaseByIDBadRequest{}
}

// WithPayload adds the payload to the get release by Id bad request response
func (o *GetReleaseByIDBadRequest) WithPayload(payload *models.APIResponse) *GetReleaseByIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get release by Id bad request response
func (o *GetReleaseByIDBadRequest) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReleaseByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReleaseByIDUnauthorizedCode is the HTTP code returned for type GetReleaseByIDUnauthorized
const GetReleaseByIDUnauthorizedCode int = 401

/*GetReleaseByIDUnauthorized Unauthorized

swagger:response getReleaseByIdUnauthorized
*/
type GetReleaseByIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetReleaseByIDUnauthorized creates GetReleaseByIDUnauthorized with default headers values
func NewGetReleaseByIDUnauthorized() *GetReleaseByIDUnauthorized {
	return &GetReleaseByIDUnauthorized{}
}

// WithPayload adds the payload to the get release by Id unauthorized response
func (o *GetReleaseByIDUnauthorized) WithPayload(payload *models.APIResponse) *GetReleaseByIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get release by Id unauthorized response
func (o *GetReleaseByIDUnauthorized) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReleaseByIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReleaseByIDNotFoundCode is the HTTP code returned for type GetReleaseByIDNotFound
const GetReleaseByIDNotFoundCode int = 404

/*GetReleaseByIDNotFound The specified resource was not found

swagger:response getReleaseByIdNotFound
*/
type GetReleaseByIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetReleaseByIDNotFound creates GetReleaseByIDNotFound with default headers values
func NewGetReleaseByIDNotFound() *GetReleaseByIDNotFound {
	return &GetReleaseByIDNotFound{}
}

// WithPayload adds the payload to the get release by Id not found response
func (o *GetReleaseByIDNotFound) WithPayload(payload *models.APIResponse) *GetReleaseByIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get release by Id not found response
func (o *GetReleaseByIDNotFound) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReleaseByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReleaseByIDInternalServerErrorCode is the HTTP code returned for type GetReleaseByIDInternalServerError
const GetReleaseByIDInternalServerErrorCode int = 500

/*GetReleaseByIDInternalServerError Internal Server Error

swagger:response getReleaseByIdInternalServerError
*/
type GetReleaseByIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetReleaseByIDInternalServerError creates GetReleaseByIDInternalServerError with default headers values
func NewGetReleaseByIDInternalServerError() *GetReleaseByIDInternalServerError {
	return &GetReleaseByIDInternalServerError{}
}

// WithPayload adds the payload to the get release by Id internal server error response
func (o *GetReleaseByIDInternalServerError) WithPayload(payload *models.APIResponse) *GetReleaseByIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get release by Id internal server error response
func (o *GetReleaseByIDInternalServerError) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReleaseByIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
