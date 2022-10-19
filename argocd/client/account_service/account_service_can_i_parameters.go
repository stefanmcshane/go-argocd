// Code generated by go-swagger; DO NOT EDIT.

package account_service

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

// NewAccountServiceCanIParams creates a new AccountServiceCanIParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountServiceCanIParams() *AccountServiceCanIParams {
	return &AccountServiceCanIParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountServiceCanIParamsWithTimeout creates a new AccountServiceCanIParams object
// with the ability to set a timeout on a request.
func NewAccountServiceCanIParamsWithTimeout(timeout time.Duration) *AccountServiceCanIParams {
	return &AccountServiceCanIParams{
		timeout: timeout,
	}
}

// NewAccountServiceCanIParamsWithContext creates a new AccountServiceCanIParams object
// with the ability to set a context for a request.
func NewAccountServiceCanIParamsWithContext(ctx context.Context) *AccountServiceCanIParams {
	return &AccountServiceCanIParams{
		Context: ctx,
	}
}

// NewAccountServiceCanIParamsWithHTTPClient creates a new AccountServiceCanIParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountServiceCanIParamsWithHTTPClient(client *http.Client) *AccountServiceCanIParams {
	return &AccountServiceCanIParams{
		HTTPClient: client,
	}
}

/* AccountServiceCanIParams contains all the parameters to send to the API endpoint
   for the account service can i operation.

   Typically these are written to a http.Request.
*/
type AccountServiceCanIParams struct {

	// Action.
	Action string

	// Resource.
	Resource string

	// Subresource.
	Subresource string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the account service can i params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountServiceCanIParams) WithDefaults() *AccountServiceCanIParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the account service can i params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountServiceCanIParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the account service can i params
func (o *AccountServiceCanIParams) WithTimeout(timeout time.Duration) *AccountServiceCanIParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account service can i params
func (o *AccountServiceCanIParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account service can i params
func (o *AccountServiceCanIParams) WithContext(ctx context.Context) *AccountServiceCanIParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account service can i params
func (o *AccountServiceCanIParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account service can i params
func (o *AccountServiceCanIParams) WithHTTPClient(client *http.Client) *AccountServiceCanIParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account service can i params
func (o *AccountServiceCanIParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAction adds the action to the account service can i params
func (o *AccountServiceCanIParams) WithAction(action string) *AccountServiceCanIParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the account service can i params
func (o *AccountServiceCanIParams) SetAction(action string) {
	o.Action = action
}

// WithResource adds the resource to the account service can i params
func (o *AccountServiceCanIParams) WithResource(resource string) *AccountServiceCanIParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the account service can i params
func (o *AccountServiceCanIParams) SetResource(resource string) {
	o.Resource = resource
}

// WithSubresource adds the subresource to the account service can i params
func (o *AccountServiceCanIParams) WithSubresource(subresource string) *AccountServiceCanIParams {
	o.SetSubresource(subresource)
	return o
}

// SetSubresource adds the subresource to the account service can i params
func (o *AccountServiceCanIParams) SetSubresource(subresource string) {
	o.Subresource = subresource
}

// WriteToRequest writes these params to a swagger request
func (o *AccountServiceCanIParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param action
	if err := r.SetPathParam("action", o.Action); err != nil {
		return err
	}

	// path param resource
	if err := r.SetPathParam("resource", o.Resource); err != nil {
		return err
	}

	// path param subresource
	if err := r.SetPathParam("subresource", o.Subresource); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
