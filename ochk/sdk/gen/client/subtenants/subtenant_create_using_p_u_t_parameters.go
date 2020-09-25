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

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// NewSubtenantCreateUsingPUTParams creates a new SubtenantCreateUsingPUTParams object
// with the default values initialized.
func NewSubtenantCreateUsingPUTParams() *SubtenantCreateUsingPUTParams {
	var ()
	return &SubtenantCreateUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSubtenantCreateUsingPUTParamsWithTimeout creates a new SubtenantCreateUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSubtenantCreateUsingPUTParamsWithTimeout(timeout time.Duration) *SubtenantCreateUsingPUTParams {
	var ()
	return &SubtenantCreateUsingPUTParams{

		timeout: timeout,
	}
}

// NewSubtenantCreateUsingPUTParamsWithContext creates a new SubtenantCreateUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewSubtenantCreateUsingPUTParamsWithContext(ctx context.Context) *SubtenantCreateUsingPUTParams {
	var ()
	return &SubtenantCreateUsingPUTParams{

		Context: ctx,
	}
}

// NewSubtenantCreateUsingPUTParamsWithHTTPClient creates a new SubtenantCreateUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSubtenantCreateUsingPUTParamsWithHTTPClient(client *http.Client) *SubtenantCreateUsingPUTParams {
	var ()
	return &SubtenantCreateUsingPUTParams{
		HTTPClient: client,
	}
}

/*SubtenantCreateUsingPUTParams contains all the parameters to send to the API endpoint
for the subtenant create using p u t operation typically these are written to a http.Request
*/
type SubtenantCreateUsingPUTParams struct {

	/*SubtenantInstance
	  subtenantInstance

	*/
	SubtenantInstance *models.SubtenantInstance

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the subtenant create using p u t params
func (o *SubtenantCreateUsingPUTParams) WithTimeout(timeout time.Duration) *SubtenantCreateUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subtenant create using p u t params
func (o *SubtenantCreateUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subtenant create using p u t params
func (o *SubtenantCreateUsingPUTParams) WithContext(ctx context.Context) *SubtenantCreateUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subtenant create using p u t params
func (o *SubtenantCreateUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subtenant create using p u t params
func (o *SubtenantCreateUsingPUTParams) WithHTTPClient(client *http.Client) *SubtenantCreateUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subtenant create using p u t params
func (o *SubtenantCreateUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSubtenantInstance adds the subtenantInstance to the subtenant create using p u t params
func (o *SubtenantCreateUsingPUTParams) WithSubtenantInstance(subtenantInstance *models.SubtenantInstance) *SubtenantCreateUsingPUTParams {
	o.SetSubtenantInstance(subtenantInstance)
	return o
}

// SetSubtenantInstance adds the subtenantInstance to the subtenant create using p u t params
func (o *SubtenantCreateUsingPUTParams) SetSubtenantInstance(subtenantInstance *models.SubtenantInstance) {
	o.SubtenantInstance = subtenantInstance
}

// WriteToRequest writes these params to a swagger request
func (o *SubtenantCreateUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SubtenantInstance != nil {
		if err := r.SetBodyParam(o.SubtenantInstance); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}