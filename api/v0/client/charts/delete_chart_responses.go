// Code generated by go-swagger; DO NOT EDIT.

package charts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"wwwin-github.cisco.com/edge/optikon-api/api/v0/models"
)

// DeleteChartReader is a Reader for the DeleteChart structure.
type DeleteChartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteChartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteChartOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteChartBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteChartUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteChartNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteChartInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteChartOK creates a DeleteChartOK with default headers values
func NewDeleteChartOK() *DeleteChartOK {
	return &DeleteChartOK{}
}

/*DeleteChartOK handles this case with default header values.

OK
*/
type DeleteChartOK struct {
}

func (o *DeleteChartOK) Error() string {
	return fmt.Sprintf("[DELETE /charts/{chartId}][%d] deleteChartOK ", 200)
}

func (o *DeleteChartOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteChartBadRequest creates a DeleteChartBadRequest with default headers values
func NewDeleteChartBadRequest() *DeleteChartBadRequest {
	return &DeleteChartBadRequest{}
}

/*DeleteChartBadRequest handles this case with default header values.

Bad Request
*/
type DeleteChartBadRequest struct {
	Payload *models.APIResponse
}

func (o *DeleteChartBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /charts/{chartId}][%d] deleteChartBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteChartBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteChartUnauthorized creates a DeleteChartUnauthorized with default headers values
func NewDeleteChartUnauthorized() *DeleteChartUnauthorized {
	return &DeleteChartUnauthorized{}
}

/*DeleteChartUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteChartUnauthorized struct {
	Payload *models.APIResponse
}

func (o *DeleteChartUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /charts/{chartId}][%d] deleteChartUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteChartUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteChartNotFound creates a DeleteChartNotFound with default headers values
func NewDeleteChartNotFound() *DeleteChartNotFound {
	return &DeleteChartNotFound{}
}

/*DeleteChartNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteChartNotFound struct {
	Payload *models.APIResponse
}

func (o *DeleteChartNotFound) Error() string {
	return fmt.Sprintf("[DELETE /charts/{chartId}][%d] deleteChartNotFound  %+v", 404, o.Payload)
}

func (o *DeleteChartNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteChartInternalServerError creates a DeleteChartInternalServerError with default headers values
func NewDeleteChartInternalServerError() *DeleteChartInternalServerError {
	return &DeleteChartInternalServerError{}
}

/*DeleteChartInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteChartInternalServerError struct {
	Payload *models.APIResponse
}

func (o *DeleteChartInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /charts/{chartId}][%d] deleteChartInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteChartInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
