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
)

// NewApplicationSetServiceListParams creates a new ApplicationSetServiceListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApplicationSetServiceListParams() *ApplicationSetServiceListParams {
	return &ApplicationSetServiceListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationSetServiceListParamsWithTimeout creates a new ApplicationSetServiceListParams object
// with the ability to set a timeout on a request.
func NewApplicationSetServiceListParamsWithTimeout(timeout time.Duration) *ApplicationSetServiceListParams {
	return &ApplicationSetServiceListParams{
		timeout: timeout,
	}
}

// NewApplicationSetServiceListParamsWithContext creates a new ApplicationSetServiceListParams object
// with the ability to set a context for a request.
func NewApplicationSetServiceListParamsWithContext(ctx context.Context) *ApplicationSetServiceListParams {
	return &ApplicationSetServiceListParams{
		Context: ctx,
	}
}

// NewApplicationSetServiceListParamsWithHTTPClient creates a new ApplicationSetServiceListParams object
// with the ability to set a custom HTTPClient for a request.
func NewApplicationSetServiceListParamsWithHTTPClient(client *http.Client) *ApplicationSetServiceListParams {
	return &ApplicationSetServiceListParams{
		HTTPClient: client,
	}
}

/* ApplicationSetServiceListParams contains all the parameters to send to the API endpoint
   for the application set service list operation.

   Typically these are written to a http.Request.
*/
type ApplicationSetServiceListParams struct {

	/* Projects.

	   the project names to restrict returned list applicationsets.
	*/
	Projects []string

	/* Selector.

	   the selector to restrict returned list to applications only with matched labels.
	*/
	Selector *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the application set service list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationSetServiceListParams) WithDefaults() *ApplicationSetServiceListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the application set service list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationSetServiceListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the application set service list params
func (o *ApplicationSetServiceListParams) WithTimeout(timeout time.Duration) *ApplicationSetServiceListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application set service list params
func (o *ApplicationSetServiceListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application set service list params
func (o *ApplicationSetServiceListParams) WithContext(ctx context.Context) *ApplicationSetServiceListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application set service list params
func (o *ApplicationSetServiceListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application set service list params
func (o *ApplicationSetServiceListParams) WithHTTPClient(client *http.Client) *ApplicationSetServiceListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application set service list params
func (o *ApplicationSetServiceListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjects adds the projects to the application set service list params
func (o *ApplicationSetServiceListParams) WithProjects(projects []string) *ApplicationSetServiceListParams {
	o.SetProjects(projects)
	return o
}

// SetProjects adds the projects to the application set service list params
func (o *ApplicationSetServiceListParams) SetProjects(projects []string) {
	o.Projects = projects
}

// WithSelector adds the selector to the application set service list params
func (o *ApplicationSetServiceListParams) WithSelector(selector *string) *ApplicationSetServiceListParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the application set service list params
func (o *ApplicationSetServiceListParams) SetSelector(selector *string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationSetServiceListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Projects != nil {

		// binding items for projects
		joinedProjects := o.bindParamProjects(reg)

		// query array param projects
		if err := r.SetQueryParam("projects", joinedProjects...); err != nil {
			return err
		}
	}

	if o.Selector != nil {

		// query param selector
		var qrSelector string

		if o.Selector != nil {
			qrSelector = *o.Selector
		}
		qSelector := qrSelector
		if qSelector != "" {

			if err := r.SetQueryParam("selector", qSelector); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamApplicationSetServiceList binds the parameter projects
func (o *ApplicationSetServiceListParams) bindParamProjects(formats strfmt.Registry) []string {
	projectsIR := o.Projects

	var projectsIC []string
	for _, projectsIIR := range projectsIR { // explode []string

		projectsIIV := projectsIIR // string as string
		projectsIC = append(projectsIC, projectsIIV)
	}

	// items.CollectionFormat: "multi"
	projectsIS := swag.JoinByFormat(projectsIC, "multi")

	return projectsIS
}