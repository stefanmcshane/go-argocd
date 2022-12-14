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

// NewApplicationServiceDeleteResourceParams creates a new ApplicationServiceDeleteResourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApplicationServiceDeleteResourceParams() *ApplicationServiceDeleteResourceParams {
	return &ApplicationServiceDeleteResourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationServiceDeleteResourceParamsWithTimeout creates a new ApplicationServiceDeleteResourceParams object
// with the ability to set a timeout on a request.
func NewApplicationServiceDeleteResourceParamsWithTimeout(timeout time.Duration) *ApplicationServiceDeleteResourceParams {
	return &ApplicationServiceDeleteResourceParams{
		timeout: timeout,
	}
}

// NewApplicationServiceDeleteResourceParamsWithContext creates a new ApplicationServiceDeleteResourceParams object
// with the ability to set a context for a request.
func NewApplicationServiceDeleteResourceParamsWithContext(ctx context.Context) *ApplicationServiceDeleteResourceParams {
	return &ApplicationServiceDeleteResourceParams{
		Context: ctx,
	}
}

// NewApplicationServiceDeleteResourceParamsWithHTTPClient creates a new ApplicationServiceDeleteResourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewApplicationServiceDeleteResourceParamsWithHTTPClient(client *http.Client) *ApplicationServiceDeleteResourceParams {
	return &ApplicationServiceDeleteResourceParams{
		HTTPClient: client,
	}
}

/* ApplicationServiceDeleteResourceParams contains all the parameters to send to the API endpoint
   for the application service delete resource operation.

   Typically these are written to a http.Request.
*/
type ApplicationServiceDeleteResourceParams struct {

	// AppNamespace.
	AppNamespace *string

	// Force.
	Force *bool

	// Group.
	Group *string

	// Kind.
	Kind *string

	// Name.
	Name string

	// Namespace.
	Namespace *string

	// Orphan.
	Orphan *bool

	// ResourceName.
	ResourceName *string

	// Version.
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the application service delete resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationServiceDeleteResourceParams) WithDefaults() *ApplicationServiceDeleteResourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the application service delete resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationServiceDeleteResourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) WithTimeout(timeout time.Duration) *ApplicationServiceDeleteResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) WithContext(ctx context.Context) *ApplicationServiceDeleteResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) WithHTTPClient(client *http.Client) *ApplicationServiceDeleteResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppNamespace adds the appNamespace to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) WithAppNamespace(appNamespace *string) *ApplicationServiceDeleteResourceParams {
	o.SetAppNamespace(appNamespace)
	return o
}

// SetAppNamespace adds the appNamespace to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) SetAppNamespace(appNamespace *string) {
	o.AppNamespace = appNamespace
}

// WithForce adds the force to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) WithForce(force *bool) *ApplicationServiceDeleteResourceParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) SetForce(force *bool) {
	o.Force = force
}

// WithGroup adds the group to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) WithGroup(group *string) *ApplicationServiceDeleteResourceParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) SetGroup(group *string) {
	o.Group = group
}

// WithKind adds the kind to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) WithKind(kind *string) *ApplicationServiceDeleteResourceParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) SetKind(kind *string) {
	o.Kind = kind
}

// WithName adds the name to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) WithName(name string) *ApplicationServiceDeleteResourceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) WithNamespace(namespace *string) *ApplicationServiceDeleteResourceParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) SetNamespace(namespace *string) {
	o.Namespace = namespace
}

// WithOrphan adds the orphan to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) WithOrphan(orphan *bool) *ApplicationServiceDeleteResourceParams {
	o.SetOrphan(orphan)
	return o
}

// SetOrphan adds the orphan to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) SetOrphan(orphan *bool) {
	o.Orphan = orphan
}

// WithResourceName adds the resourceName to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) WithResourceName(resourceName *string) *ApplicationServiceDeleteResourceParams {
	o.SetResourceName(resourceName)
	return o
}

// SetResourceName adds the resourceName to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) SetResourceName(resourceName *string) {
	o.ResourceName = resourceName
}

// WithVersion adds the version to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) WithVersion(version *string) *ApplicationServiceDeleteResourceParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the application service delete resource params
func (o *ApplicationServiceDeleteResourceParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationServiceDeleteResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Force != nil {

		// query param force
		var qrForce bool

		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}
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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
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

	if o.Orphan != nil {

		// query param orphan
		var qrOrphan bool

		if o.Orphan != nil {
			qrOrphan = *o.Orphan
		}
		qOrphan := swag.FormatBool(qrOrphan)
		if qOrphan != "" {

			if err := r.SetQueryParam("orphan", qOrphan); err != nil {
				return err
			}
		}
	}

	if o.ResourceName != nil {

		// query param resourceName
		var qrResourceName string

		if o.ResourceName != nil {
			qrResourceName = *o.ResourceName
		}
		qResourceName := qrResourceName
		if qResourceName != "" {

			if err := r.SetQueryParam("resourceName", qResourceName); err != nil {
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
