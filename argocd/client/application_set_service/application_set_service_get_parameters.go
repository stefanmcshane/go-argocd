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
)

// NewApplicationSetServiceGetParams creates a new ApplicationSetServiceGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApplicationSetServiceGetParams() *ApplicationSetServiceGetParams {
	return &ApplicationSetServiceGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationSetServiceGetParamsWithTimeout creates a new ApplicationSetServiceGetParams object
// with the ability to set a timeout on a request.
func NewApplicationSetServiceGetParamsWithTimeout(timeout time.Duration) *ApplicationSetServiceGetParams {
	return &ApplicationSetServiceGetParams{
		timeout: timeout,
	}
}

// NewApplicationSetServiceGetParamsWithContext creates a new ApplicationSetServiceGetParams object
// with the ability to set a context for a request.
func NewApplicationSetServiceGetParamsWithContext(ctx context.Context) *ApplicationSetServiceGetParams {
	return &ApplicationSetServiceGetParams{
		Context: ctx,
	}
}

// NewApplicationSetServiceGetParamsWithHTTPClient creates a new ApplicationSetServiceGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewApplicationSetServiceGetParamsWithHTTPClient(client *http.Client) *ApplicationSetServiceGetParams {
	return &ApplicationSetServiceGetParams{
		HTTPClient: client,
	}
}

/* ApplicationSetServiceGetParams contains all the parameters to send to the API endpoint
   for the application set service get operation.

   Typically these are written to a http.Request.
*/
type ApplicationSetServiceGetParams struct {

	/* Name.

	   the applicationsets's name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the application set service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationSetServiceGetParams) WithDefaults() *ApplicationSetServiceGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the application set service get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationSetServiceGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the application set service get params
func (o *ApplicationSetServiceGetParams) WithTimeout(timeout time.Duration) *ApplicationSetServiceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application set service get params
func (o *ApplicationSetServiceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application set service get params
func (o *ApplicationSetServiceGetParams) WithContext(ctx context.Context) *ApplicationSetServiceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application set service get params
func (o *ApplicationSetServiceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application set service get params
func (o *ApplicationSetServiceGetParams) WithHTTPClient(client *http.Client) *ApplicationSetServiceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application set service get params
func (o *ApplicationSetServiceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the application set service get params
func (o *ApplicationSetServiceGetParams) WithName(name string) *ApplicationSetServiceGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the application set service get params
func (o *ApplicationSetServiceGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationSetServiceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
