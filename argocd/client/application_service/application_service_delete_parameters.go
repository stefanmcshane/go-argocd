// Code generated by go-swagger; DO NOT EDIT.

package application_service

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
)

// NewApplicationServiceDeleteParams creates a new ApplicationServiceDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApplicationServiceDeleteParams() *ApplicationServiceDeleteParams {
	return &ApplicationServiceDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationServiceDeleteParamsWithTimeout creates a new ApplicationServiceDeleteParams object
// with the ability to set a timeout on a request.
func NewApplicationServiceDeleteParamsWithTimeout(timeout time.Duration) *ApplicationServiceDeleteParams {
	return &ApplicationServiceDeleteParams{
		timeout: timeout,
	}
}

// NewApplicationServiceDeleteParamsWithContext creates a new ApplicationServiceDeleteParams object
// with the ability to set a context for a request.
func NewApplicationServiceDeleteParamsWithContext(ctx context.Context) *ApplicationServiceDeleteParams {
	return &ApplicationServiceDeleteParams{
		Context: ctx,
	}
}

// NewApplicationServiceDeleteParamsWithHTTPClient creates a new ApplicationServiceDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewApplicationServiceDeleteParamsWithHTTPClient(client *http.Client) *ApplicationServiceDeleteParams {
	return &ApplicationServiceDeleteParams{
		HTTPClient: client,
	}
}

/* ApplicationServiceDeleteParams contains all the parameters to send to the API endpoint
   for the application service delete operation.

   Typically these are written to a http.Request.
*/
type ApplicationServiceDeleteParams struct {

	// AppNamespace.
	AppNamespace *string

	// Cascade.
	Cascade *bool

	// Name.
	Name string

	// PropagationPolicy.
	PropagationPolicy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the application service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationServiceDeleteParams) WithDefaults() *ApplicationServiceDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the application service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationServiceDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the application service delete params
func (o *ApplicationServiceDeleteParams) WithTimeout(timeout time.Duration) *ApplicationServiceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application service delete params
func (o *ApplicationServiceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application service delete params
func (o *ApplicationServiceDeleteParams) WithContext(ctx context.Context) *ApplicationServiceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application service delete params
func (o *ApplicationServiceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application service delete params
func (o *ApplicationServiceDeleteParams) WithHTTPClient(client *http.Client) *ApplicationServiceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application service delete params
func (o *ApplicationServiceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppNamespace adds the appNamespace to the application service delete params
func (o *ApplicationServiceDeleteParams) WithAppNamespace(appNamespace *string) *ApplicationServiceDeleteParams {
	o.SetAppNamespace(appNamespace)
	return o
}

// SetAppNamespace adds the appNamespace to the application service delete params
func (o *ApplicationServiceDeleteParams) SetAppNamespace(appNamespace *string) {
	o.AppNamespace = appNamespace
}

// WithCascade adds the cascade to the application service delete params
func (o *ApplicationServiceDeleteParams) WithCascade(cascade *bool) *ApplicationServiceDeleteParams {
	o.SetCascade(cascade)
	return o
}

// SetCascade adds the cascade to the application service delete params
func (o *ApplicationServiceDeleteParams) SetCascade(cascade *bool) {
	o.Cascade = cascade
}

// WithName adds the name to the application service delete params
func (o *ApplicationServiceDeleteParams) WithName(name string) *ApplicationServiceDeleteParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the application service delete params
func (o *ApplicationServiceDeleteParams) SetName(name string) {
	o.Name = name
}

// WithPropagationPolicy adds the propagationPolicy to the application service delete params
func (o *ApplicationServiceDeleteParams) WithPropagationPolicy(propagationPolicy *string) *ApplicationServiceDeleteParams {
	o.SetPropagationPolicy(propagationPolicy)
	return o
}

// SetPropagationPolicy adds the propagationPolicy to the application service delete params
func (o *ApplicationServiceDeleteParams) SetPropagationPolicy(propagationPolicy *string) {
	o.PropagationPolicy = propagationPolicy
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationServiceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppNamespace != nil {

		// query param appNamespace
		var qrAppNamespace string

		if o.AppNamespace != nil {
			qrAppNamespace = *o.AppNamespace
		}
		qAppNamespace := qrAppNamespace
		if qAppNamespace != "" {

			if err := r.SetQueryParam("appNamespace", qAppNamespace); err != nil {
				return err
			}
		}
	}

	if o.Cascade != nil {

		// query param cascade
		var qrCascade bool

		if o.Cascade != nil {
			qrCascade = *o.Cascade
		}
		qCascade := swag.FormatBool(qrCascade)
		if qCascade != "" {

			if err := r.SetQueryParam("cascade", qCascade); err != nil {
				return err
			}
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.PropagationPolicy != nil {

		// query param propagationPolicy
		var qrPropagationPolicy string

		if o.PropagationPolicy != nil {
			qrPropagationPolicy = *o.PropagationPolicy
		}
		qPropagationPolicy := qrPropagationPolicy
		if qPropagationPolicy != "" {

			if err := r.SetQueryParam("propagationPolicy", qPropagationPolicy); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}