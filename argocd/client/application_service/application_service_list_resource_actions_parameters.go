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

// NewApplicationServiceListResourceActionsParams creates a new ApplicationServiceListResourceActionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApplicationServiceListResourceActionsParams() *ApplicationServiceListResourceActionsParams {
	return &ApplicationServiceListResourceActionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationServiceListResourceActionsParamsWithTimeout creates a new ApplicationServiceListResourceActionsParams object
// with the ability to set a timeout on a request.
func NewApplicationServiceListResourceActionsParamsWithTimeout(timeout time.Duration) *ApplicationServiceListResourceActionsParams {
	return &ApplicationServiceListResourceActionsParams{
		timeout: timeout,
	}
}

// NewApplicationServiceListResourceActionsParamsWithContext creates a new ApplicationServiceListResourceActionsParams object
// with the ability to set a context for a request.
func NewApplicationServiceListResourceActionsParamsWithContext(ctx context.Context) *ApplicationServiceListResourceActionsParams {
	return &ApplicationServiceListResourceActionsParams{
		Context: ctx,
	}
}

// NewApplicationServiceListResourceActionsParamsWithHTTPClient creates a new ApplicationServiceListResourceActionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewApplicationServiceListResourceActionsParamsWithHTTPClient(client *http.Client) *ApplicationServiceListResourceActionsParams {
	return &ApplicationServiceListResourceActionsParams{
		HTTPClient: client,
	}
}

/* ApplicationServiceListResourceActionsParams contains all the parameters to send to the API endpoint
   for the application service list resource actions operation.

   Typically these are written to a http.Request.
*/
type ApplicationServiceListResourceActionsParams struct {

	// AppNamespace.
	AppNamespace *string

	// Group.
	Group *string

	// Kind.
	Kind *string

	// Name.
	Name string

	// Namespace.
	Namespace *string

	// ResourceName.
	ResourceName *string

	// Version.
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the application service list resource actions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationServiceListResourceActionsParams) WithDefaults() *ApplicationServiceListResourceActionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the application service list resource actions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationServiceListResourceActionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the application service list resource actions params
func (o *ApplicationServiceListResourceActionsParams) WithTimeout(timeout time.Duration) *ApplicationServiceListResourceActionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application service list resource actions params
func (o *ApplicationServiceListResourceActionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application service list resource actions params
func (o *ApplicationServiceListResourceActionsParams) WithContext(ctx context.Context) *ApplicationServiceListResourceActionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application service list resource actions params
func (o *ApplicationServiceListResourceActionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application service list resource actions params
func (o *ApplicationServiceListResourceActionsParams) WithHTTPClient(client *http.Client) *ApplicationServiceListResourceActionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application service list resource actions params
func (o *ApplicationServiceListResourceActionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppNamespace adds the appNamespace to the application service list resource actions params
func (o *ApplicationServiceListResourceActionsParams) WithAppNamespace(appNamespace *string) *ApplicationServiceListResourceActionsParams {
	o.SetAppNamespace(appNamespace)
	return o
}

// SetAppNamespace adds the appNamespace to the application service list resource actions params
func (o *ApplicationServiceListResourceActionsParams) SetAppNamespace(appNamespace *string) {
	o.AppNamespace = appNamespace
}

// WithGroup adds the group to the application service list resource actions params
func (o *ApplicationServiceListResourceActionsParams) WithGroup(group *string) *ApplicationServiceListResourceActionsParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the application service list resource actions params
func (o *ApplicationServiceListResourceActionsParams) SetGroup(group *string) {
	o.Group = group
}

// WithKind adds the kind to the application service list resource actions params
func (o *ApplicationServiceListResourceActionsParams) WithKind(kind *string) *ApplicationServiceListResourceActionsParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the application service list resource actions params
func (o *ApplicationServiceListResourceActionsParams) SetKind(kind *string) {
	o.Kind = kind
}

// WithName adds the name to the application service list resource actions params
func (o *ApplicationServiceListResourceActionsParams) WithName(name string) *ApplicationServiceListResourceActionsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the application service list resource actions params
func (o *ApplicationServiceListResourceActionsParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the application service list resource actions params
func (o *ApplicationServiceListResourceActionsParams) WithNamespace(namespace *string) *ApplicationServiceListResourceActionsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the application service list resource actions params
func (o *ApplicationServiceListResourceActionsParams) SetNamespace(namespace *string) {
	o.Namespace = namespace
}

// WithResourceName adds the resourceName to the application service list resource actions params
func (o *ApplicationServiceListResourceActionsParams) WithResourceName(resourceName *string) *ApplicationServiceListResourceActionsParams {
	o.SetResourceName(resourceName)
	return o
}

// SetResourceName adds the resourceName to the application service list resource actions params
func (o *ApplicationServiceListResourceActionsParams) SetResourceName(resourceName *string) {
	o.ResourceName = resourceName
}

// WithVersion adds the version to the application service list resource actions params
func (o *ApplicationServiceListResourceActionsParams) WithVersion(version *string) *ApplicationServiceListResourceActionsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the application service list resource actions params
func (o *ApplicationServiceListResourceActionsParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationServiceListResourceActionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
