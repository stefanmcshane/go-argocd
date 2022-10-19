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
)

// NewApplicationServiceResourceTreeParams creates a new ApplicationServiceResourceTreeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApplicationServiceResourceTreeParams() *ApplicationServiceResourceTreeParams {
	return &ApplicationServiceResourceTreeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationServiceResourceTreeParamsWithTimeout creates a new ApplicationServiceResourceTreeParams object
// with the ability to set a timeout on a request.
func NewApplicationServiceResourceTreeParamsWithTimeout(timeout time.Duration) *ApplicationServiceResourceTreeParams {
	return &ApplicationServiceResourceTreeParams{
		timeout: timeout,
	}
}

// NewApplicationServiceResourceTreeParamsWithContext creates a new ApplicationServiceResourceTreeParams object
// with the ability to set a context for a request.
func NewApplicationServiceResourceTreeParamsWithContext(ctx context.Context) *ApplicationServiceResourceTreeParams {
	return &ApplicationServiceResourceTreeParams{
		Context: ctx,
	}
}

// NewApplicationServiceResourceTreeParamsWithHTTPClient creates a new ApplicationServiceResourceTreeParams object
// with the ability to set a custom HTTPClient for a request.
func NewApplicationServiceResourceTreeParamsWithHTTPClient(client *http.Client) *ApplicationServiceResourceTreeParams {
	return &ApplicationServiceResourceTreeParams{
		HTTPClient: client,
	}
}

/* ApplicationServiceResourceTreeParams contains all the parameters to send to the API endpoint
   for the application service resource tree operation.

   Typically these are written to a http.Request.
*/
type ApplicationServiceResourceTreeParams struct {

	// AppNamespace.
	AppNamespace *string

	// ApplicationName.
	ApplicationName string

	// Group.
	Group *string

	// Kind.
	Kind *string

	// Name.
	Name *string

	// Namespace.
	Namespace *string

	// Version.
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the application service resource tree params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationServiceResourceTreeParams) WithDefaults() *ApplicationServiceResourceTreeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the application service resource tree params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationServiceResourceTreeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the application service resource tree params
func (o *ApplicationServiceResourceTreeParams) WithTimeout(timeout time.Duration) *ApplicationServiceResourceTreeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application service resource tree params
func (o *ApplicationServiceResourceTreeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application service resource tree params
func (o *ApplicationServiceResourceTreeParams) WithContext(ctx context.Context) *ApplicationServiceResourceTreeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application service resource tree params
func (o *ApplicationServiceResourceTreeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application service resource tree params
func (o *ApplicationServiceResourceTreeParams) WithHTTPClient(client *http.Client) *ApplicationServiceResourceTreeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application service resource tree params
func (o *ApplicationServiceResourceTreeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppNamespace adds the appNamespace to the application service resource tree params
func (o *ApplicationServiceResourceTreeParams) WithAppNamespace(appNamespace *string) *ApplicationServiceResourceTreeParams {
	o.SetAppNamespace(appNamespace)
	return o
}

// SetAppNamespace adds the appNamespace to the application service resource tree params
func (o *ApplicationServiceResourceTreeParams) SetAppNamespace(appNamespace *string) {
	o.AppNamespace = appNamespace
}

// WithApplicationName adds the applicationName to the application service resource tree params
func (o *ApplicationServiceResourceTreeParams) WithApplicationName(applicationName string) *ApplicationServiceResourceTreeParams {
	o.SetApplicationName(applicationName)
	return o
}

// SetApplicationName adds the applicationName to the application service resource tree params
func (o *ApplicationServiceResourceTreeParams) SetApplicationName(applicationName string) {
	o.ApplicationName = applicationName
}

// WithGroup adds the group to the application service resource tree params
func (o *ApplicationServiceResourceTreeParams) WithGroup(group *string) *ApplicationServiceResourceTreeParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the application service resource tree params
func (o *ApplicationServiceResourceTreeParams) SetGroup(group *string) {
	o.Group = group
}

// WithKind adds the kind to the application service resource tree params
func (o *ApplicationServiceResourceTreeParams) WithKind(kind *string) *ApplicationServiceResourceTreeParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the application service resource tree params
func (o *ApplicationServiceResourceTreeParams) SetKind(kind *string) {
	o.Kind = kind
}

// WithName adds the name to the application service resource tree params
func (o *ApplicationServiceResourceTreeParams) WithName(name *string) *ApplicationServiceResourceTreeParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the application service resource tree params
func (o *ApplicationServiceResourceTreeParams) SetName(name *string) {
	o.Name = name
}

// WithNamespace adds the namespace to the application service resource tree params
func (o *ApplicationServiceResourceTreeParams) WithNamespace(namespace *string) *ApplicationServiceResourceTreeParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the application service resource tree params
func (o *ApplicationServiceResourceTreeParams) SetNamespace(namespace *string) {
	o.Namespace = namespace
}

// WithVersion adds the version to the application service resource tree params
func (o *ApplicationServiceResourceTreeParams) WithVersion(version *string) *ApplicationServiceResourceTreeParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the application service resource tree params
func (o *ApplicationServiceResourceTreeParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationServiceResourceTreeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param applicationName
	if err := r.SetPathParam("applicationName", o.ApplicationName); err != nil {
		return err
	}

	if o.Group != nil {

		// query param group
		var qrGroup string

		if o.Group != nil {
			qrGroup = *o.Group
		}
		qGroup := qrGroup
		if qGroup != "" {

			if err := r.SetQueryParam("group", qGroup); err != nil {
				return err
			}
		}
	}

	if o.Kind != nil {

		// query param kind
		var qrKind string

		if o.Kind != nil {
			qrKind = *o.Kind
		}
		qKind := qrKind
		if qKind != "" {

			if err := r.SetQueryParam("kind", qKind); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.Namespace != nil {

		// query param namespace
		var qrNamespace string

		if o.Namespace != nil {
			qrNamespace = *o.Namespace
		}
		qNamespace := qrNamespace
		if qNamespace != "" {

			if err := r.SetQueryParam("namespace", qNamespace); err != nil {
				return err
			}
		}
	}

	if o.Version != nil {

		// query param version
		var qrVersion string

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}