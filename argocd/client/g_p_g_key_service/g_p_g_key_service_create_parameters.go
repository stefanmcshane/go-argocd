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
	"github.com/go-openapi/swag"

	"github.com/stefanmcshane/go-argocd/argocd/models"
)

// NewGPGKeyServiceCreateParams creates a new GPGKeyServiceCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGPGKeyServiceCreateParams() *GPGKeyServiceCreateParams {
	return &GPGKeyServiceCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGPGKeyServiceCreateParamsWithTimeout creates a new GPGKeyServiceCreateParams object
// with the ability to set a timeout on a request.
func NewGPGKeyServiceCreateParamsWithTimeout(timeout time.Duration) *GPGKeyServiceCreateParams {
	return &GPGKeyServiceCreateParams{
		timeout: timeout,
	}
}

// NewGPGKeyServiceCreateParamsWithContext creates a new GPGKeyServiceCreateParams object
// with the ability to set a context for a request.
func NewGPGKeyServiceCreateParamsWithContext(ctx context.Context) *GPGKeyServiceCreateParams {
	return &GPGKeyServiceCreateParams{
		Context: ctx,
	}
}

// NewGPGKeyServiceCreateParamsWithHTTPClient creates a new GPGKeyServiceCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGPGKeyServiceCreateParamsWithHTTPClient(client *http.Client) *GPGKeyServiceCreateParams {
	return &GPGKeyServiceCreateParams{
		HTTPClient: client,
	}
}

/* GPGKeyServiceCreateParams contains all the parameters to send to the API endpoint
   for the g p g key service create operation.

   Typically these are written to a http.Request.
*/
type GPGKeyServiceCreateParams struct {

	/* Body.

	   Raw key data of the GPG key(s) to create
	*/
	Body *models.V1alpha1GnuPGPublicKey

	/* Upsert.

	   Whether to upsert already existing public keys.
	*/
	Upsert *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the g p g key service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GPGKeyServiceCreateParams) WithDefaults() *GPGKeyServiceCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the g p g key service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GPGKeyServiceCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the g p g key service create params
func (o *GPGKeyServiceCreateParams) WithTimeout(timeout time.Duration) *GPGKeyServiceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the g p g key service create params
func (o *GPGKeyServiceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the g p g key service create params
func (o *GPGKeyServiceCreateParams) WithContext(ctx context.Context) *GPGKeyServiceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the g p g key service create params
func (o *GPGKeyServiceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the g p g key service create params
func (o *GPGKeyServiceCreateParams) WithHTTPClient(client *http.Client) *GPGKeyServiceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the g p g key service create params
func (o *GPGKeyServiceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the g p g key service create params
func (o *GPGKeyServiceCreateParams) WithBody(body *models.V1alpha1GnuPGPublicKey) *GPGKeyServiceCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the g p g key service create params
func (o *GPGKeyServiceCreateParams) SetBody(body *models.V1alpha1GnuPGPublicKey) {
	o.Body = body
}

// WithUpsert adds the upsert to the g p g key service create params
func (o *GPGKeyServiceCreateParams) WithUpsert(upsert *bool) *GPGKeyServiceCreateParams {
	o.SetUpsert(upsert)
	return o
}

// SetUpsert adds the upsert to the g p g key service create params
func (o *GPGKeyServiceCreateParams) SetUpsert(upsert *bool) {
	o.Upsert = upsert
}

// WriteToRequest writes these params to a swagger request
func (o *GPGKeyServiceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Upsert != nil {

		// query param upsert
		var qrUpsert bool

		if o.Upsert != nil {
			qrUpsert = *o.Upsert
		}
		qUpsert := swag.FormatBool(qrUpsert)
		if qUpsert != "" {

			if err := r.SetQueryParam("upsert", qUpsert); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}