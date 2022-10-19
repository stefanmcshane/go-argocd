// Code generated by go-swagger; DO NOT EDIT.

package g_p_g_key_service

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

// NewGPGKeyServiceListParams creates a new GPGKeyServiceListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGPGKeyServiceListParams() *GPGKeyServiceListParams {
	return &GPGKeyServiceListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGPGKeyServiceListParamsWithTimeout creates a new GPGKeyServiceListParams object
// with the ability to set a timeout on a request.
func NewGPGKeyServiceListParamsWithTimeout(timeout time.Duration) *GPGKeyServiceListParams {
	return &GPGKeyServiceListParams{
		timeout: timeout,
	}
}

// NewGPGKeyServiceListParamsWithContext creates a new GPGKeyServiceListParams object
// with the ability to set a context for a request.
func NewGPGKeyServiceListParamsWithContext(ctx context.Context) *GPGKeyServiceListParams {
	return &GPGKeyServiceListParams{
		Context: ctx,
	}
}

// NewGPGKeyServiceListParamsWithHTTPClient creates a new GPGKeyServiceListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGPGKeyServiceListParamsWithHTTPClient(client *http.Client) *GPGKeyServiceListParams {
	return &GPGKeyServiceListParams{
		HTTPClient: client,
	}
}

/* GPGKeyServiceListParams contains all the parameters to send to the API endpoint
   for the g p g key service list operation.

   Typically these are written to a http.Request.
*/
type GPGKeyServiceListParams struct {

	/* KeyID.

	   The GPG key ID to query for.
	*/
	KeyID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the g p g key service list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GPGKeyServiceListParams) WithDefaults() *GPGKeyServiceListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the g p g key service list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GPGKeyServiceListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the g p g key service list params
func (o *GPGKeyServiceListParams) WithTimeout(timeout time.Duration) *GPGKeyServiceListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the g p g key service list params
func (o *GPGKeyServiceListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the g p g key service list params
func (o *GPGKeyServiceListParams) WithContext(ctx context.Context) *GPGKeyServiceListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the g p g key service list params
func (o *GPGKeyServiceListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the g p g key service list params
func (o *GPGKeyServiceListParams) WithHTTPClient(client *http.Client) *GPGKeyServiceListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the g p g key service list params
func (o *GPGKeyServiceListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKeyID adds the keyID to the g p g key service list params
func (o *GPGKeyServiceListParams) WithKeyID(keyID *string) *GPGKeyServiceListParams {
	o.SetKeyID(keyID)
	return o
}

// SetKeyID adds the keyId to the g p g key service list params
func (o *GPGKeyServiceListParams) SetKeyID(keyID *string) {
	o.KeyID = keyID
}

// WriteToRequest writes these params to a swagger request
func (o *GPGKeyServiceListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.KeyID != nil {

		// query param keyID
		var qrKeyID string

		if o.KeyID != nil {
			qrKeyID = *o.KeyID
		}
		qKeyID := qrKeyID
		if qKeyID != "" {

			if err := r.SetQueryParam("keyID", qKeyID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
