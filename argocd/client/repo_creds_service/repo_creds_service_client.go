// Code generated by go-swagger; DO NOT EDIT.

package repo_creds_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new repo creds service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for repo creds service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	RepoCredsServiceCreateRepositoryCredentials(params *RepoCredsServiceCreateRepositoryCredentialsParams, opts ...ClientOption) (*RepoCredsServiceCreateRepositoryCredentialsOK, error)

	RepoCredsServiceDeleteRepositoryCredentials(params *RepoCredsServiceDeleteRepositoryCredentialsParams, opts ...ClientOption) (*RepoCredsServiceDeleteRepositoryCredentialsOK, error)

	RepoCredsServiceListRepositoryCredentials(params *RepoCredsServiceListRepositoryCredentialsParams, opts ...ClientOption) (*RepoCredsServiceListRepositoryCredentialsOK, error)

	RepoCredsServiceUpdateRepositoryCredentials(params *RepoCredsServiceUpdateRepositoryCredentialsParams, opts ...ClientOption) (*RepoCredsServiceUpdateRepositoryCredentialsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  RepoCredsServiceCreateRepositoryCredentials creates repository credentials creates a new repository credential set
*/
func (a *Client) RepoCredsServiceCreateRepositoryCredentials(params *RepoCredsServiceCreateRepositoryCredentialsParams, opts ...ClientOption) (*RepoCredsServiceCreateRepositoryCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRepoCredsServiceCreateRepositoryCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RepoCredsService_CreateRepositoryCredentials",
		Method:             "POST",
		PathPattern:        "/api/v1/repocreds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RepoCredsServiceCreateRepositoryCredentialsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RepoCredsServiceCreateRepositoryCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RepoCredsServiceCreateRepositoryCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  RepoCredsServiceDeleteRepositoryCredentials deletes repository credentials deletes a repository credential set from the configuration
*/
func (a *Client) RepoCredsServiceDeleteRepositoryCredentials(params *RepoCredsServiceDeleteRepositoryCredentialsParams, opts ...ClientOption) (*RepoCredsServiceDeleteRepositoryCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRepoCredsServiceDeleteRepositoryCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RepoCredsService_DeleteRepositoryCredentials",
		Method:             "DELETE",
		PathPattern:        "/api/v1/repocreds/{url}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RepoCredsServiceDeleteRepositoryCredentialsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RepoCredsServiceDeleteRepositoryCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RepoCredsServiceDeleteRepositoryCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  RepoCredsServiceListRepositoryCredentials lists repository credentials gets a list of all configured repository credential sets
*/
func (a *Client) RepoCredsServiceListRepositoryCredentials(params *RepoCredsServiceListRepositoryCredentialsParams, opts ...ClientOption) (*RepoCredsServiceListRepositoryCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRepoCredsServiceListRepositoryCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RepoCredsService_ListRepositoryCredentials",
		Method:             "GET",
		PathPattern:        "/api/v1/repocreds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RepoCredsServiceListRepositoryCredentialsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RepoCredsServiceListRepositoryCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RepoCredsServiceListRepositoryCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  RepoCredsServiceUpdateRepositoryCredentials updates repository credentials updates a repository credential set
*/
func (a *Client) RepoCredsServiceUpdateRepositoryCredentials(params *RepoCredsServiceUpdateRepositoryCredentialsParams, opts ...ClientOption) (*RepoCredsServiceUpdateRepositoryCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRepoCredsServiceUpdateRepositoryCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RepoCredsService_UpdateRepositoryCredentials",
		Method:             "PUT",
		PathPattern:        "/api/v1/repocreds/{creds.url}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RepoCredsServiceUpdateRepositoryCredentialsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RepoCredsServiceUpdateRepositoryCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RepoCredsServiceUpdateRepositoryCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
