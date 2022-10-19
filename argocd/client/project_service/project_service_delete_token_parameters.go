// Code generated by go-swagger; DO NOT EDIT.

package project_service

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

// NewProjectServiceDeleteTokenParams creates a new ProjectServiceDeleteTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectServiceDeleteTokenParams() *ProjectServiceDeleteTokenParams {
	return &ProjectServiceDeleteTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectServiceDeleteTokenParamsWithTimeout creates a new ProjectServiceDeleteTokenParams object
// with the ability to set a timeout on a request.
func NewProjectServiceDeleteTokenParamsWithTimeout(timeout time.Duration) *ProjectServiceDeleteTokenParams {
	return &ProjectServiceDeleteTokenParams{
		timeout: timeout,
	}
}

// NewProjectServiceDeleteTokenParamsWithContext creates a new ProjectServiceDeleteTokenParams object
// with the ability to set a context for a request.
func NewProjectServiceDeleteTokenParamsWithContext(ctx context.Context) *ProjectServiceDeleteTokenParams {
	return &ProjectServiceDeleteTokenParams{
		Context: ctx,
	}
}

// NewProjectServiceDeleteTokenParamsWithHTTPClient creates a new ProjectServiceDeleteTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectServiceDeleteTokenParamsWithHTTPClient(client *http.Client) *ProjectServiceDeleteTokenParams {
	return &ProjectServiceDeleteTokenParams{
		HTTPClient: client,
	}
}

/* ProjectServiceDeleteTokenParams contains all the parameters to send to the API endpoint
   for the project service delete token operation.

   Typically these are written to a http.Request.
*/
type ProjectServiceDeleteTokenParams struct {

	// Iat.
	//
	// Format: int64
	Iat string

	// ID.
	ID *string

	// Project.
	Project string

	// Role.
	Role string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the project service delete token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectServiceDeleteTokenParams) WithDefaults() *ProjectServiceDeleteTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the project service delete token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectServiceDeleteTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the project service delete token params
func (o *ProjectServiceDeleteTokenParams) WithTimeout(timeout time.Duration) *ProjectServiceDeleteTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project service delete token params
func (o *ProjectServiceDeleteTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project service delete token params
func (o *ProjectServiceDeleteTokenParams) WithContext(ctx context.Context) *ProjectServiceDeleteTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project service delete token params
func (o *ProjectServiceDeleteTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project service delete token params
func (o *ProjectServiceDeleteTokenParams) WithHTTPClient(client *http.Client) *ProjectServiceDeleteTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project service delete token params
func (o *ProjectServiceDeleteTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIat adds the iat to the project service delete token params
func (o *ProjectServiceDeleteTokenParams) WithIat(iat string) *ProjectServiceDeleteTokenParams {
	o.SetIat(iat)
	return o
}

// SetIat adds the iat to the project service delete token params
func (o *ProjectServiceDeleteTokenParams) SetIat(iat string) {
	o.Iat = iat
}

// WithID adds the id to the project service delete token params
func (o *ProjectServiceDeleteTokenParams) WithID(id *string) *ProjectServiceDeleteTokenParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the project service delete token params
func (o *ProjectServiceDeleteTokenParams) SetID(id *string) {
	o.ID = id
}

// WithProject adds the project to the project service delete token params
func (o *ProjectServiceDeleteTokenParams) WithProject(project string) *ProjectServiceDeleteTokenParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the project service delete token params
func (o *ProjectServiceDeleteTokenParams) SetProject(project string) {
	o.Project = project
}

// WithRole adds the role to the project service delete token params
func (o *ProjectServiceDeleteTokenParams) WithRole(role string) *ProjectServiceDeleteTokenParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the project service delete token params
func (o *ProjectServiceDeleteTokenParams) SetRole(role string) {
	o.Role = role
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectServiceDeleteTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param iat
	if err := r.SetPathParam("iat", o.Iat); err != nil {
		return err
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	// path param role
	if err := r.SetPathParam("role", o.Role); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}