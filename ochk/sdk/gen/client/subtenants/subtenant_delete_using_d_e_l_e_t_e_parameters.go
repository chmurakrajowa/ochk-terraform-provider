// Code generated by go-swagger; DO NOT EDIT.

package subtenants

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

// NewSubtenantDeleteUsingDELETEParams creates a new SubtenantDeleteUsingDELETEParams object
// with the default values initialized.
func NewSubtenantDeleteUsingDELETEParams() *SubtenantDeleteUsingDELETEParams {
	var ()
	return &SubtenantDeleteUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSubtenantDeleteUsingDELETEParamsWithTimeout creates a new SubtenantDeleteUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSubtenantDeleteUsingDELETEParamsWithTimeout(timeout time.Duration) *SubtenantDeleteUsingDELETEParams {
	var ()
	return &SubtenantDeleteUsingDELETEParams{

		timeout: timeout,
	}
}

// NewSubtenantDeleteUsingDELETEParamsWithContext creates a new SubtenantDeleteUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewSubtenantDeleteUsingDELETEParamsWithContext(ctx context.Context) *SubtenantDeleteUsingDELETEParams {
	var ()
	return &SubtenantDeleteUsingDELETEParams{

		Context: ctx,
	}
}

// NewSubtenantDeleteUsingDELETEParamsWithHTTPClient creates a new SubtenantDeleteUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSubtenantDeleteUsingDELETEParamsWithHTTPClient(client *http.Client) *SubtenantDeleteUsingDELETEParams {
	var ()
	return &SubtenantDeleteUsingDELETEParams{
		HTTPClient: client,
	}
}

/*SubtenantDeleteUsingDELETEParams contains all the parameters to send to the API endpoint
for the subtenant delete using d e l e t e operation typically these are written to a http.Request
*/
type SubtenantDeleteUsingDELETEParams struct {

	/*SubtenantID
	  subtenantId

	*/
	SubtenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the subtenant delete using d e l e t e params
func (o *SubtenantDeleteUsingDELETEParams) WithTimeout(timeout time.Duration) *SubtenantDeleteUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subtenant delete using d e l e t e params
func (o *SubtenantDeleteUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subtenant delete using d e l e t e params
func (o *SubtenantDeleteUsingDELETEParams) WithContext(ctx context.Context) *SubtenantDeleteUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subtenant delete using d e l e t e params
func (o *SubtenantDeleteUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subtenant delete using d e l e t e params
func (o *SubtenantDeleteUsingDELETEParams) WithHTTPClient(client *http.Client) *SubtenantDeleteUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subtenant delete using d e l e t e params
func (o *SubtenantDeleteUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSubtenantID adds the subtenantID to the subtenant delete using d e l e t e params
func (o *SubtenantDeleteUsingDELETEParams) WithSubtenantID(subtenantID string) *SubtenantDeleteUsingDELETEParams {
	o.SetSubtenantID(subtenantID)
	return o
}

// SetSubtenantID adds the subtenantId to the subtenant delete using d e l e t e params
func (o *SubtenantDeleteUsingDELETEParams) SetSubtenantID(subtenantID string) {
	o.SubtenantID = subtenantID
}

// WriteToRequest writes these params to a swagger request
func (o *SubtenantDeleteUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param subtenantId
	if err := r.SetPathParam("subtenantId", o.SubtenantID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
