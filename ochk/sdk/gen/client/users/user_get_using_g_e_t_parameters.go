// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewUserGetUsingGETParams creates a new UserGetUsingGETParams object
// with the default values initialized.
func NewUserGetUsingGETParams() *UserGetUsingGETParams {
	var ()
	return &UserGetUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserGetUsingGETParamsWithTimeout creates a new UserGetUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserGetUsingGETParamsWithTimeout(timeout time.Duration) *UserGetUsingGETParams {
	var ()
	return &UserGetUsingGETParams{

		timeout: timeout,
	}
}

// NewUserGetUsingGETParamsWithContext creates a new UserGetUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserGetUsingGETParamsWithContext(ctx context.Context) *UserGetUsingGETParams {
	var ()
	return &UserGetUsingGETParams{

		Context: ctx,
	}
}

// NewUserGetUsingGETParamsWithHTTPClient creates a new UserGetUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserGetUsingGETParamsWithHTTPClient(client *http.Client) *UserGetUsingGETParams {
	var ()
	return &UserGetUsingGETParams{
		HTTPClient: client,
	}
}

/*UserGetUsingGETParams contains all the parameters to send to the API endpoint
for the user get using g e t operation typically these are written to a http.Request
*/
type UserGetUsingGETParams struct {

	/*UserID
	  userId

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user get using g e t params
func (o *UserGetUsingGETParams) WithTimeout(timeout time.Duration) *UserGetUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user get using g e t params
func (o *UserGetUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user get using g e t params
func (o *UserGetUsingGETParams) WithContext(ctx context.Context) *UserGetUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user get using g e t params
func (o *UserGetUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user get using g e t params
func (o *UserGetUsingGETParams) WithHTTPClient(client *http.Client) *UserGetUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user get using g e t params
func (o *UserGetUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the user get using g e t params
func (o *UserGetUsingGETParams) WithUserID(userID string) *UserGetUsingGETParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the user get using g e t params
func (o *UserGetUsingGETParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UserGetUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
