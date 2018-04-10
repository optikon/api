// Code generated by go-swagger; DO NOT EDIT.

package releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"wwwin-github.cisco.com/edge/optikon-api/api/v0/models"
)

// DeleteReleaseReader is a Reader for the DeleteRelease structure.
type DeleteReleaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteReleaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteReleaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteReleaseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteReleaseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteReleaseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteReleaseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteReleaseOK creates a DeleteReleaseOK with default headers values
func NewDeleteReleaseOK() *DeleteReleaseOK {
	return &DeleteReleaseOK{}
}

/*DeleteReleaseOK handles this case with default header values.

OK
*/
type DeleteReleaseOK struct {
}

func (o *DeleteReleaseOK) Error() string {
	return fmt.Sprintf("[DELETE /releases/{releaseId}][%d] deleteReleaseOK ", 200)
}

func (o *DeleteReleaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteReleaseBadRequest creates a DeleteReleaseBadRequest with default headers values
func NewDeleteReleaseBadRequest() *DeleteReleaseBadRequest {
	return &DeleteReleaseBadRequest{}
}

/*DeleteReleaseBadRequest handles this case with default header values.

Bad Request
*/
type DeleteReleaseBadRequest struct {
	Payload *models.APIResponse
}

func (o *DeleteReleaseBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /releases/{releaseId}][%d] deleteReleaseBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteReleaseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteReleaseUnauthorized creates a DeleteReleaseUnauthorized with default headers values
func NewDeleteReleaseUnauthorized() *DeleteReleaseUnauthorized {
	return &DeleteReleaseUnauthorized{}
}

/*DeleteReleaseUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteReleaseUnauthorized struct {
	Payload *models.APIResponse
}

func (o *DeleteReleaseUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /releases/{releaseId}][%d] deleteReleaseUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteReleaseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteReleaseNotFound creates a DeleteReleaseNotFound with default headers values
func NewDeleteReleaseNotFound() *DeleteReleaseNotFound {
	return &DeleteReleaseNotFound{}
}

/*DeleteReleaseNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteReleaseNotFound struct {
	Payload *models.APIResponse
}

func (o *DeleteReleaseNotFound) Error() string {
	return fmt.Sprintf("[DELETE /releases/{releaseId}][%d] deleteReleaseNotFound  %+v", 404, o.Payload)
}

func (o *DeleteReleaseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteReleaseInternalServerError creates a DeleteReleaseInternalServerError with default headers values
func NewDeleteReleaseInternalServerError() *DeleteReleaseInternalServerError {
	return &DeleteReleaseInternalServerError{}
}

/*DeleteReleaseInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteReleaseInternalServerError struct {
	Payload *models.APIResponse
}

func (o *DeleteReleaseInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /releases/{releaseId}][%d] deleteReleaseInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteReleaseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
