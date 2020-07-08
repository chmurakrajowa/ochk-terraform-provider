// Code generated by go-swagger; DO NOT EDIT.

package security_group_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListUsingGET6Params creates a new ListUsingGET6Params object
// with the default values initialized.
func NewListUsingGET6Params() *ListUsingGET6Params {

	return &ListUsingGET6Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewListUsingGET6ParamsWithTimeout creates a new ListUsingGET6Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewListUsingGET6ParamsWithTimeout(timeout time.Duration) *ListUsingGET6Params {

	return &ListUsingGET6Params{

		timeout: timeout,
	}
}

// NewListUsingGET6ParamsWithContext creates a new ListUsingGET6Params object
// with the default values initialized, and the ability to set a context for a request
func NewListUsingGET6ParamsWithContext(ctx context.Context) *ListUsingGET6Params {

	return &ListUsingGET6Params{

		Context: ctx,
	}
}

// NewListUsingGET6ParamsWithHTTPClient creates a new ListUsingGET6Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListUsingGET6ParamsWithHTTPClient(client *http.Client) *ListUsingGET6Params {

	return &ListUsingGET6Params{
		HTTPClient: client,
	}
}

/*ListUsingGET6Params contains all the parameters to send to the API endpoint
for the list using g e t 6 operation typically these are written to a http.Request
*/
type ListUsingGET6Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list using g e t 6 params
func (o *ListUsingGET6Params) WithTimeout(timeout time.Duration) *ListUsingGET6Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list using g e t 6 params
func (o *ListUsingGET6Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list using g e t 6 params
func (o *ListUsingGET6Params) WithContext(ctx context.Context) *ListUsingGET6Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list using g e t 6 params
func (o *ListUsingGET6Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list using g e t 6 params
func (o *ListUsingGET6Params) WithHTTPClient(client *http.Client) *ListUsingGET6Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list using g e t 6 params
func (o *ListUsingGET6Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListUsingGET6Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
