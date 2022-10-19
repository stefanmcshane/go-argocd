// Code generated by go-swagger; DO NOT EDIT.

package repository_service

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

// NewRepositoryServiceGetAppDetailsParams creates a new RepositoryServiceGetAppDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepositoryServiceGetAppDetailsParams() *RepositoryServiceGetAppDetailsParams {
	return &RepositoryServiceGetAppDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepositoryServiceGetAppDetailsParamsWithTimeout creates a new RepositoryServiceGetAppDetailsParams object
// with the ability to set a timeout on a request.
func NewRepositoryServiceGetAppDetailsParamsWithTimeout(timeout time.Duration) *RepositoryServiceGetAppDetailsParams {
	return &RepositoryServiceGetAppDetailsParams{
		timeout: timeout,
	}
}

// NewRepositoryServiceGetAppDetailsParamsWithContext creates a new RepositoryServiceGetAppDetailsParams object
// with the ability to set a context for a request.
func NewRepositoryServiceGetAppDetailsParamsWithContext(ctx context.Context) *RepositoryServiceGetAppDetailsParams {
	return &RepositoryServiceGetAppDetailsParams{
		Context: ctx,
	}
}

// NewRepositoryServiceGetAppDetailsParamsWithHTTPClient creates a new RepositoryServiceGetAppDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepositoryServiceGetAppDetailsParamsWithHTTPClient(client *http.Client) *RepositoryServiceGetAppDetailsParams {
	return &RepositoryServiceGetAppDetailsParams{
		HTTPClient: client,
	}
}

/* RepositoryServiceGetAppDetailsParams contains all the parameters to send to the API endpoint
   for the repository service get app details operation.

   Typically these are written to a http.Request.
*/
type RepositoryServiceGetAppDetailsParams struct {

	// Body.
	Body *models.RepositoryRepoAppDetailsQuery

	/* SourceRepoURL.

	   RepoURL is the URL to the repository (Git or Helm) that contains the application manifests
	*/
	SourceRepoURL string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repository service get app details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepositoryServiceGetAppDetailsParams) WithDefaults() *RepositoryServiceGetAppDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repository service get app details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepositoryServiceGetAppDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repository service get app details params
func (o *RepositoryServiceGetAppDetailsParams) WithTimeout(timeout time.Duration) *RepositoryServiceGetAppDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repository service get app details params
func (o *RepositoryServiceGetAppDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repository service get app details params
func (o *RepositoryServiceGetAppDetailsParams) WithContext(ctx context.Context) *RepositoryServiceGetAppDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repository service get app details params
func (o *RepositoryServiceGetAppDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repository service get app details params
func (o *RepositoryServiceGetAppDetailsParams) WithHTTPClient(client *http.Client) *RepositoryServiceGetAppDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repository service get app details params
func (o *RepositoryServiceGetAppDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repository service get app details params
func (o *RepositoryServiceGetAppDetailsParams) WithBody(body *models.RepositoryRepoAppDetailsQuery) *RepositoryServiceGetAppDetailsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repository service get app details params
func (o *RepositoryServiceGetAppDetailsParams) SetBody(body *models.RepositoryRepoAppDetailsQuery) {
	o.Body = body
}

// WithSourceRepoURL adds the sourceRepoURL to the repository service get app details params
func (o *RepositoryServiceGetAppDetailsParams) WithSourceRepoURL(sourceRepoURL string) *RepositoryServiceGetAppDetailsParams {
	o.SetSourceRepoURL(sourceRepoURL)
	return o
}

// SetSourceRepoURL adds the sourceRepoUrl to the repository service get app details params
func (o *RepositoryServiceGetAppDetailsParams) SetSourceRepoURL(sourceRepoURL string) {
	o.SourceRepoURL = sourceRepoURL
}

// WriteToRequest writes these params to a swagger request
func (o *RepositoryServiceGetAppDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param source.repoURL
	if err := r.SetPathParam("source.repoURL", o.SourceRepoURL); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
