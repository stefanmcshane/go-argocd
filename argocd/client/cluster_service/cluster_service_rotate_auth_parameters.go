// Code generated by go-swagger; DO NOT EDIT.

package cluster_service

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

// NewClusterServiceRotateAuthParams creates a new ClusterServiceRotateAuthParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterServiceRotateAuthParams() *ClusterServiceRotateAuthParams {
	return &ClusterServiceRotateAuthParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterServiceRotateAuthParamsWithTimeout creates a new ClusterServiceRotateAuthParams object
// with the ability to set a timeout on a request.
func NewClusterServiceRotateAuthParamsWithTimeout(timeout time.Duration) *ClusterServiceRotateAuthParams {
	return &ClusterServiceRotateAuthParams{
		timeout: timeout,
	}
}

// NewClusterServiceRotateAuthParamsWithContext creates a new ClusterServiceRotateAuthParams object
// with the ability to set a context for a request.
func NewClusterServiceRotateAuthParamsWithContext(ctx context.Context) *ClusterServiceRotateAuthParams {
	return &ClusterServiceRotateAuthParams{
		Context: ctx,
	}
}

// NewClusterServiceRotateAuthParamsWithHTTPClient creates a new ClusterServiceRotateAuthParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterServiceRotateAuthParamsWithHTTPClient(client *http.Client) *ClusterServiceRotateAuthParams {
	return &ClusterServiceRotateAuthParams{
		HTTPClient: client,
	}
}

/* ClusterServiceRotateAuthParams contains all the parameters to send to the API endpoint
   for the cluster service rotate auth operation.

   Typically these are written to a http.Request.
*/
type ClusterServiceRotateAuthParams struct {

	/* IDValue.

	   value holds the cluster server URL or cluster name
	*/
	IDValue string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster service rotate auth params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterServiceRotateAuthParams) WithDefaults() *ClusterServiceRotateAuthParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster service rotate auth params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterServiceRotateAuthParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cluster service rotate auth params
func (o *ClusterServiceRotateAuthParams) WithTimeout(timeout time.Duration) *ClusterServiceRotateAuthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster service rotate auth params
func (o *ClusterServiceRotateAuthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster service rotate auth params
func (o *ClusterServiceRotateAuthParams) WithContext(ctx context.Context) *ClusterServiceRotateAuthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster service rotate auth params
func (o *ClusterServiceRotateAuthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster service rotate auth params
func (o *ClusterServiceRotateAuthParams) WithHTTPClient(client *http.Client) *ClusterServiceRotateAuthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster service rotate auth params
func (o *ClusterServiceRotateAuthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIDValue adds the iDValue to the cluster service rotate auth params
func (o *ClusterServiceRotateAuthParams) WithIDValue(iDValue string) *ClusterServiceRotateAuthParams {
	o.SetIDValue(iDValue)
	return o
}

// SetIDValue adds the idValue to the cluster service rotate auth params
func (o *ClusterServiceRotateAuthParams) SetIDValue(iDValue string) {
	o.IDValue = iDValue
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterServiceRotateAuthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id.value
	if err := r.SetPathParam("id.value", o.IDValue); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}