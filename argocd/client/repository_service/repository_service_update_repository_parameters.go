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

// NewRepositoryServiceUpdateRepositoryParams creates a new RepositoryServiceUpdateRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepositoryServiceUpdateRepositoryParams() *RepositoryServiceUpdateRepositoryParams {
	return &RepositoryServiceUpdateRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepositoryServiceUpdateRepositoryParamsWithTimeout creates a new RepositoryServiceUpdateRepositoryParams object
// with the ability to set a timeout on a request.
func NewRepositoryServiceUpdateRepositoryParamsWithTimeout(timeout time.Duration) *RepositoryServiceUpdateRepositoryParams {
	return &RepositoryServiceUpdateRepositoryParams{
		timeout: timeout,
	}
}

// NewRepositoryServiceUpdateRepositoryParamsWithContext creates a new RepositoryServiceUpdateRepositoryParams object
// with the ability to set a context for a request.
func NewRepositoryServiceUpdateRepositoryParamsWithContext(ctx context.Context) *RepositoryServiceUpdateRepositoryParams {
	return &RepositoryServiceUpdateRepositoryParams{
		Context: ctx,
	}
}

// NewRepositoryServiceUpdateRepositoryParamsWithHTTPClient creates a new RepositoryServiceUpdateRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepositoryServiceUpdateRepositoryParamsWithHTTPClient(client *http.Client) *RepositoryServiceUpdateRepositoryParams {
	return &RepositoryServiceUpdateRepositoryParams{
		HTTPClient: client,
	}
}

/* RepositoryServiceUpdateRepositoryParams contains all the parameters to send to the API endpoint
   for the repository service update repository operation.

   Typically these are written to a http.Request.
*/
type RepositoryServiceUpdateRepositoryParams struct {

	// Body.
	Body *models.V1alpha1Repository

	/* RepoRepo.

	   Repo contains the URL to the remote repository
	*/
	RepoRepo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repository service update repository params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepositoryServiceUpdateRepositoryParams) WithDefaults() *RepositoryServiceUpdateRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repository service update repository params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepositoryServiceUpdateRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repository service update repository params
func (o *RepositoryServiceUpdateRepositoryParams) WithTimeout(timeout time.Duration) *RepositoryServiceUpdateRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repository service update repository params
func (o *RepositoryServiceUpdateRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repository service update repository params
func (o *RepositoryServiceUpdateRepositoryParams) WithContext(ctx context.Context) *RepositoryServiceUpdateRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repository service update repository params
func (o *RepositoryServiceUpdateRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repository service update repository params
func (o *RepositoryServiceUpdateRepositoryParams) WithHTTPClient(client *http.Client) *RepositoryServiceUpdateRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repository service update repository params
func (o *RepositoryServiceUpdateRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repository service update repository params
func (o *RepositoryServiceUpdateRepositoryParams) WithBody(body *models.V1alpha1Repository) *RepositoryServiceUpdateRepositoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repository service update repository params
func (o *RepositoryServiceUpdateRepositoryParams) SetBody(body *models.V1alpha1Repository) {
	o.Body = body
}

// WithRepoRepo adds the repoRepo to the repository service update repository params
func (o *RepositoryServiceUpdateRepositoryParams) WithRepoRepo(repoRepo string) *RepositoryServiceUpdateRepositoryParams {
	o.SetRepoRepo(repoRepo)
	return o
}

// SetRepoRepo adds the repoRepo to the repository service update repository params
func (o *RepositoryServiceUpdateRepositoryParams) SetRepoRepo(repoRepo string) {
	o.RepoRepo = repoRepo
}

// WriteToRequest writes these params to a swagger request
func (o *RepositoryServiceUpdateRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param repo.repo
	if err := r.SetPathParam("repo.repo", o.RepoRepo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
