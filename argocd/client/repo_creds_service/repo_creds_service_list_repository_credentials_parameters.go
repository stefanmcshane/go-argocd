// Code generated by go-swagger; DO NOT EDIT.

package repo_creds_service

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

// NewRepoCredsServiceListRepositoryCredentialsParams creates a new RepoCredsServiceListRepositoryCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoCredsServiceListRepositoryCredentialsParams() *RepoCredsServiceListRepositoryCredentialsParams {
	return &RepoCredsServiceListRepositoryCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoCredsServiceListRepositoryCredentialsParamsWithTimeout creates a new RepoCredsServiceListRepositoryCredentialsParams object
// with the ability to set a timeout on a request.
func NewRepoCredsServiceListRepositoryCredentialsParamsWithTimeout(timeout time.Duration) *RepoCredsServiceListRepositoryCredentialsParams {
	return &RepoCredsServiceListRepositoryCredentialsParams{
		timeout: timeout,
	}
}

// NewRepoCredsServiceListRepositoryCredentialsParamsWithContext creates a new RepoCredsServiceListRepositoryCredentialsParams object
// with the ability to set a context for a request.
func NewRepoCredsServiceListRepositoryCredentialsParamsWithContext(ctx context.Context) *RepoCredsServiceListRepositoryCredentialsParams {
	return &RepoCredsServiceListRepositoryCredentialsParams{
		Context: ctx,
	}
}

// NewRepoCredsServiceListRepositoryCredentialsParamsWithHTTPClient creates a new RepoCredsServiceListRepositoryCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoCredsServiceListRepositoryCredentialsParamsWithHTTPClient(client *http.Client) *RepoCredsServiceListRepositoryCredentialsParams {
	return &RepoCredsServiceListRepositoryCredentialsParams{
		HTTPClient: client,
	}
}

/* RepoCredsServiceListRepositoryCredentialsParams contains all the parameters to send to the API endpoint
   for the repo creds service list repository credentials operation.

   Typically these are written to a http.Request.
*/
type RepoCredsServiceListRepositoryCredentialsParams struct {

	/* URL.

	   Repo URL for query.
	*/
	URL *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repo creds service list repository credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoCredsServiceListRepositoryCredentialsParams) WithDefaults() *RepoCredsServiceListRepositoryCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo creds service list repository credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoCredsServiceListRepositoryCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo creds service list repository credentials params
func (o *RepoCredsServiceListRepositoryCredentialsParams) WithTimeout(timeout time.Duration) *RepoCredsServiceListRepositoryCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo creds service list repository credentials params
func (o *RepoCredsServiceListRepositoryCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo creds service list repository credentials params
func (o *RepoCredsServiceListRepositoryCredentialsParams) WithContext(ctx context.Context) *RepoCredsServiceListRepositoryCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo creds service list repository credentials params
func (o *RepoCredsServiceListRepositoryCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo creds service list repository credentials params
func (o *RepoCredsServiceListRepositoryCredentialsParams) WithHTTPClient(client *http.Client) *RepoCredsServiceListRepositoryCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo creds service list repository credentials params
func (o *RepoCredsServiceListRepositoryCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithURL adds the url to the repo creds service list repository credentials params
func (o *RepoCredsServiceListRepositoryCredentialsParams) WithURL(url *string) *RepoCredsServiceListRepositoryCredentialsParams {
	o.SetURL(url)
	return o
}

// SetURL adds the url to the repo creds service list repository credentials params
func (o *RepoCredsServiceListRepositoryCredentialsParams) SetURL(url *string) {
	o.URL = url
}

// WriteToRequest writes these params to a swagger request
func (o *RepoCredsServiceListRepositoryCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.URL != nil {

		// query param url
		var qrURL string

		if o.URL != nil {
			qrURL = *o.URL
		}
		qURL := qrURL
		if qURL != "" {

			if err := r.SetQueryParam("url", qURL); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}