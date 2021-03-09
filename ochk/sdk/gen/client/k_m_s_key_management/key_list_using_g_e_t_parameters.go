// Code generated by go-swagger; DO NOT EDIT.

package k_m_s_key_management

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

// NewKeyListUsingGETParams creates a new KeyListUsingGETParams object
// with the default values initialized.
func NewKeyListUsingGETParams() *KeyListUsingGETParams {
	var ()
	return &KeyListUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewKeyListUsingGETParamsWithTimeout creates a new KeyListUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewKeyListUsingGETParamsWithTimeout(timeout time.Duration) *KeyListUsingGETParams {
	var ()
	return &KeyListUsingGETParams{

		timeout: timeout,
	}
}

// NewKeyListUsingGETParamsWithContext creates a new KeyListUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewKeyListUsingGETParamsWithContext(ctx context.Context) *KeyListUsingGETParams {
	var ()
	return &KeyListUsingGETParams{

		Context: ctx,
	}
}

// NewKeyListUsingGETParamsWithHTTPClient creates a new KeyListUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewKeyListUsingGETParamsWithHTTPClient(client *http.Client) *KeyListUsingGETParams {
	var ()
	return &KeyListUsingGETParams{
		HTTPClient: client,
	}
}

/*KeyListUsingGETParams contains all the parameters to send to the API endpoint
for the key list using g e t operation typically these are written to a http.Request
*/
type KeyListUsingGETParams struct {

	/*DisplayName
	  displayName

	*/
	DisplayName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the key list using g e t params
func (o *KeyListUsingGETParams) WithTimeout(timeout time.Duration) *KeyListUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the key list using g e t params
func (o *KeyListUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the key list using g e t params
func (o *KeyListUsingGETParams) WithContext(ctx context.Context) *KeyListUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the key list using g e t params
func (o *KeyListUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the key list using g e t params
func (o *KeyListUsingGETParams) WithHTTPClient(client *http.Client) *KeyListUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the key list using g e t params
func (o *KeyListUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisplayName adds the displayName to the key list using g e t params
func (o *KeyListUsingGETParams) WithDisplayName(displayName *string) *KeyListUsingGETParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the key list using g e t params
func (o *KeyListUsingGETParams) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WriteToRequest writes these params to a swagger request
func (o *KeyListUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DisplayName != nil {

		// query param displayName
		var qrDisplayName string
		if o.DisplayName != nil {
			qrDisplayName = *o.DisplayName
		}
		qDisplayName := qrDisplayName
		if qDisplayName != "" {
			if err := r.SetQueryParam("displayName", qDisplayName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}