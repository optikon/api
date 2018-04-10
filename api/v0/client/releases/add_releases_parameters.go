// Code generated by go-swagger; DO NOT EDIT.

package releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"wwwin-github.cisco.com/edge/optikon-api/api/v0/models"
)

// NewAddReleasesParams creates a new AddReleasesParams object
// with the default values initialized.
func NewAddReleasesParams() *AddReleasesParams {
	var ()
	return &AddReleasesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddReleasesParamsWithTimeout creates a new AddReleasesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddReleasesParamsWithTimeout(timeout time.Duration) *AddReleasesParams {
	var ()
	return &AddReleasesParams{

		timeout: timeout,
	}
}

// NewAddReleasesParamsWithContext creates a new AddReleasesParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddReleasesParamsWithContext(ctx context.Context) *AddReleasesParams {
	var ()
	return &AddReleasesParams{

		Context: ctx,
	}
}

// NewAddReleasesParamsWithHTTPClient creates a new AddReleasesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddReleasesParamsWithHTTPClient(client *http.Client) *AddReleasesParams {
	var ()
	return &AddReleasesParams{
		HTTPClient: client,
	}
}

/*AddReleasesParams contains all the parameters to send to the API endpoint
for the add releases operation typically these are written to a http.Request
*/
type AddReleasesParams struct {

	/*Body*/
	Body *models.ReleaseRelease

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add releases params
func (o *AddReleasesParams) WithTimeout(timeout time.Duration) *AddReleasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add releases params
func (o *AddReleasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add releases params
func (o *AddReleasesParams) WithContext(ctx context.Context) *AddReleasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add releases params
func (o *AddReleasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add releases params
func (o *AddReleasesParams) WithHTTPClient(client *http.Client) *AddReleasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add releases params
func (o *AddReleasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add releases params
func (o *AddReleasesParams) WithBody(body *models.ReleaseRelease) *AddReleasesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add releases params
func (o *AddReleasesParams) SetBody(body *models.ReleaseRelease) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddReleasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
