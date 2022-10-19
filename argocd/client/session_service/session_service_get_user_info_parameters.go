// Code generated by go-swagger; DO NOT EDIT.

package session_service

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

// NewSessionServiceGetUserInfoParams creates a new SessionServiceGetUserInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSessionServiceGetUserInfoParams() *SessionServiceGetUserInfoParams {
	return &SessionServiceGetUserInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSessionServiceGetUserInfoParamsWithTimeout creates a new SessionServiceGetUserInfoParams object
// with the ability to set a timeout on a request.
func NewSessionServiceGetUserInfoParamsWithTimeout(timeout time.Duration) *SessionServiceGetUserInfoParams {
	return &SessionServiceGetUserInfoParams{
		timeout: timeout,
	}
}

// NewSessionServiceGetUserInfoParamsWithContext creates a new SessionServiceGetUserInfoParams object
// with the ability to set a context for a request.
func NewSessionServiceGetUserInfoParamsWithContext(ctx context.Context) *SessionServiceGetUserInfoParams {
	return &SessionServiceGetUserInfoParams{
		Context: ctx,
	}
}

// NewSessionServiceGetUserInfoParamsWithHTTPClient creates a new SessionServiceGetUserInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewSessionServiceGetUserInfoParamsWithHTTPClient(client *http.Client) *SessionServiceGetUserInfoParams {
	return &SessionServiceGetUserInfoParams{
		HTTPClient: client,
	}
}

/* SessionServiceGetUserInfoParams contains all the parameters to send to the API endpoint
   for the session service get user info operation.

   Typically these are written to a http.Request.
*/
type SessionServiceGetUserInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the session service get user info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SessionServiceGetUserInfoParams) WithDefaults() *SessionServiceGetUserInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the session service get user info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SessionServiceGetUserInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the session service get user info params
func (o *SessionServiceGetUserInfoParams) WithTimeout(timeout time.Duration) *SessionServiceGetUserInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the session service get user info params
func (o *SessionServiceGetUserInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the session service get user info params
func (o *SessionServiceGetUserInfoParams) WithContext(ctx context.Context) *SessionServiceGetUserInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the session service get user info params
func (o *SessionServiceGetUserInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the session service get user info params
func (o *SessionServiceGetUserInfoParams) WithHTTPClient(client *http.Client) *SessionServiceGetUserInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the session service get user info params
func (o *SessionServiceGetUserInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SessionServiceGetUserInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}