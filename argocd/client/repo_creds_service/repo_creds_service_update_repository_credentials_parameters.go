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

	"github.com/stefanmcshane/go-argocd/argocd/models"
)

// NewRepoCredsServiceUpdateRepositoryCredentialsParams creates a new RepoCredsServiceUpdateRepositoryCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepoCredsServiceUpdateRepositoryCredentialsParams() *RepoCredsServiceUpdateRepositoryCredentialsParams {
	return &RepoCredsServiceUpdateRepositoryCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepoCredsServiceUpdateRepositoryCredentialsParamsWithTimeout creates a new RepoCredsServiceUpdateRepositoryCredentialsParams object
// with the ability to set a timeout on a request.
func NewRepoCredsServiceUpdateRepositoryCredentialsParamsWithTimeout(timeout time.Duration) *RepoCredsServiceUpdateRepositoryCredentialsParams {
	return &RepoCredsServiceUpdateRepositoryCredentialsParams{
		timeout: timeout,
	}
}

// NewRepoCredsServiceUpdateRepositoryCredentialsParamsWithContext creates a new RepoCredsServiceUpdateRepositoryCredentialsParams object
// with the ability to set a context for a request.
func NewRepoCredsServiceUpdateRepositoryCredentialsParamsWithContext(ctx context.Context) *RepoCredsServiceUpdateRepositoryCredentialsParams {
	return &RepoCredsServiceUpdateRepositoryCredentialsParams{
		Context: ctx,
	}
}

// NewRepoCredsServiceUpdateRepositoryCredentialsParamsWithHTTPClient creates a new RepoCredsServiceUpdateRepositoryCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepoCredsServiceUpdateRepositoryCredentialsParamsWithHTTPClient(client *http.Client) *RepoCredsServiceUpdateRepositoryCredentialsParams {
	return &RepoCredsServiceUpdateRepositoryCredentialsParams{
		HTTPClient: client,
	}
}

/* RepoCredsServiceUpdateRepositoryCredentialsParams contains all the parameters to send to the API endpoint
   for the repo creds service update repository credentials operation.

   Typically these are written to a http.Request.
*/
type RepoCredsServiceUpdateRepositoryCredentialsParams struct {

	// Body.
	Body *models.V1alpha1RepoCreds

	/* CredsURL.

	   URL is the URL that this credentials matches to
	*/
	CredsURL string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repo creds service update repository credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoCredsServiceUpdateRepositoryCredentialsParams) WithDefaults() *RepoCredsServiceUpdateRepositoryCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repo creds service update repository credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepoCredsServiceUpdateRepositoryCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repo creds service update repository credentials params
func (o *RepoCredsServiceUpdateRepositoryCredentialsParams) WithTimeout(timeout time.Duration) *RepoCredsServiceUpdateRepositoryCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo creds service update repository credentials params
func (o *RepoCredsServiceUpdateRepositoryCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo creds service update repository credentials params
func (o *RepoCredsServiceUpdateRepositoryCredentialsParams) WithContext(ctx context.Context) *RepoCredsServiceUpdateRepositoryCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo creds service update repository credentials params
func (o *RepoCredsServiceUpdateRepositoryCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo creds service update repository credentials params
func (o *RepoCredsServiceUpdateRepositoryCredentialsParams) WithHTTPClient(client *http.Client) *RepoCredsServiceUpdateRepositoryCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo creds service update repository credentials params
func (o *RepoCredsServiceUpdateRepositoryCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repo creds service update repository credentials params
func (o *RepoCredsServiceUpdateRepositoryCredentialsParams) WithBody(body *models.V1alpha1RepoCreds) *RepoCredsServiceUpdateRepositoryCredentialsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repo creds service update repository credentials params
func (o *RepoCredsServiceUpdateRepositoryCredentialsParams) SetBody(body *models.V1alpha1RepoCreds) {
	o.Body = body
}

// WithCredsURL adds the credsURL to the repo creds service update repository credentials params
func (o *RepoCredsServiceUpdateRepositoryCredentialsParams) WithCredsURL(credsURL string) *RepoCredsServiceUpdateRepositoryCredentialsParams {
	o.SetCredsURL(credsURL)
	return o
}

// SetCredsURL adds the credsUrl to the repo creds service update repository credentials params
func (o *RepoCredsServiceUpdateRepositoryCredentialsParams) SetCredsURL(credsURL string) {
	o.CredsURL = credsURL
}

// WriteToRequest writes these params to a swagger request
func (o *RepoCredsServiceUpdateRepositoryCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param creds.url
	if err := r.SetPathParam("creds.url", o.CredsURL); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
