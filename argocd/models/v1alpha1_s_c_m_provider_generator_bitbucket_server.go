// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha1SCMProviderGeneratorBitbucketServer SCMProviderGeneratorBitbucketServer defines connection info specific to Bitbucket Server.
//
// swagger:model v1alpha1SCMProviderGeneratorBitbucketServer
type V1alpha1SCMProviderGeneratorBitbucketServer struct {

	// Scan all branches instead of just the default branch.
	AllBranches bool `json:"allBranches,omitempty"`

	// The Bitbucket Server REST API URL to talk to. Required.
	API string `json:"api,omitempty"`

	// basic auth
	BasicAuth *V1alpha1BasicAuthBitbucketServer `json:"basicAuth,omitempty"`

	// Project to scan. Required.
	Project string `json:"project,omitempty"`
}

// Validate validates this v1alpha1 s c m provider generator bitbucket server
func (m *V1alpha1SCMProviderGeneratorBitbucketServer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBasicAuth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1SCMProviderGeneratorBitbucketServer) validateBasicAuth(formats strfmt.Registry) error {
	if swag.IsZero(m.BasicAuth) { // not required
		return nil
	}

	if m.BasicAuth != nil {
		if err := m.BasicAuth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("basicAuth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("basicAuth")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1alpha1 s c m provider generator bitbucket server based on the context it is used
func (m *V1alpha1SCMProviderGeneratorBitbucketServer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBasicAuth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1SCMProviderGeneratorBitbucketServer) contextValidateBasicAuth(ctx context.Context, formats strfmt.Registry) error {

	if m.BasicAuth != nil {
		if err := m.BasicAuth.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("basicAuth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("basicAuth")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1SCMProviderGeneratorBitbucketServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1SCMProviderGeneratorBitbucketServer) UnmarshalBinary(b []byte) error {
	var res V1alpha1SCMProviderGeneratorBitbucketServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
