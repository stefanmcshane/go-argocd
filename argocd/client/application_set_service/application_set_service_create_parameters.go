// Code generated by go-swagger; DO NOT EDIT.

package application_set_service

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

// NewApplicationSetServiceCreateParams creates a new ApplicationSetServiceCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApplicationSetServiceCreateParams() *ApplicationSetServiceCreateParams {
	return &ApplicationSetServiceCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationSetServiceCreateParamsWithTimeout creates a new ApplicationSetServiceCreateParams object
// with the ability to set a timeout on a request.
func NewApplicationSetServiceCreateParamsWithTimeout(timeout time.Duration) *ApplicationSetServiceCreateParams {
	return &ApplicationSetServiceCreateParams{
		timeout: timeout,
	}
}

// NewApplicationSetServiceCreateParamsWithContext creates a new ApplicationSetServiceCreateParams object
// with the ability to set a context for a request.
func NewApplicationSetServiceCreateParamsWithContext(ctx context.Context) *ApplicationSetServiceCreateParams {
	return &ApplicationSetServiceCreateParams{
		Context: ctx,
	}
}

// NewApplicationSetServiceCreateParamsWithHTTPClient creates a new ApplicationSetServiceCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewApplicationSetServiceCreateParamsWithHTTPClient(client *http.Client) *ApplicationSetServiceCreateParams {
	return &ApplicationSetServiceCreateParams{
		HTTPClient: client,
	}
}

/* ApplicationSetServiceCreateParams contains all the parameters to send to the API endpoint
   for the application set service create operation.

   Typically these are written to a http.Request.
*/
type ApplicationSetServiceCreateParams struct {

	// Body.
	Body *models.V1alpha1ApplicationSet

	// Upsert.
	Upsert *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the application set service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationSetServiceCreateParams) WithDefaults() *ApplicationSetServiceCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the application set service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationSetServiceCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the application set service create params
func (o *ApplicationSetServiceCreateParams) WithTimeout(timeout time.Duration) *ApplicationSetServiceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application set service create params
func (o *ApplicationSetServiceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application set service create params
func (o *ApplicationSetServiceCreateParams) WithContext(ctx context.Context) *ApplicationSetServiceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application set service create params
func (o *ApplicationSetServiceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application set service create params
func (o *ApplicationSetServiceCreateParams) WithHTTPClient(client *http.Client) *ApplicationSetServiceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application set service create params
func (o *ApplicationSetServiceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the application set service create params
func (o *ApplicationSetServiceCreateParams) WithBody(body *models.V1alpha1ApplicationSet) *ApplicationSetServiceCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the application set service create params
func (o *ApplicationSetServiceCreateParams) SetBody(body *models.V1alpha1ApplicationSet) {
	o.Body = body
}

// WithUpsert adds the upsert to the application set service create params
func (o *ApplicationSetServiceCreateParams) WithUpsert(upsert *bool) *ApplicationSetServiceCreateParams {
	o.SetUpsert(upsert)
	return o
}

// SetUpsert adds the upsert to the application set service create params
func (o *ApplicationSetServiceCreateParams) SetUpsert(upsert *bool) {
	o.Upsert = upsert
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationSetServiceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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