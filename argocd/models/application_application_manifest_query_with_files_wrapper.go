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

// ApplicationApplicationManifestQueryWithFilesWrapper application application manifest query with files wrapper
//
// swagger:model applicationApplicationManifestQueryWithFilesWrapper
type ApplicationApplicationManifestQueryWithFilesWrapper struct {

	// chunk
	Chunk *ApplicationFileChunk `json:"chunk,omitempty"`

	// query
	Query *ApplicationApplicationManifestQueryWithFiles `json:"query,omitempty"`
}

// Validate validates this application application manifest query with files wrapper
func (m *ApplicationApplicationManifestQueryWithFilesWrapper) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChunk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationApplicationManifestQueryWithFilesWrapper) validateChunk(formats strfmt.Registry) error {
	if swag.IsZero(m.Chunk) { // not required
		return nil
	}

	if m.Chunk != nil {
		if err := m.Chunk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chunk")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chunk")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationApplicationManifestQueryWithFilesWrapper) validateQuery(formats strfmt.Registry) error {
	if swag.IsZero(m.Query) { // not required
		return nil
	}

	if m.Query != nil {
		if err := m.Query.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("query")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application application manifest query with files wrapper based on the context it is used
func (m *ApplicationApplicationManifestQueryWithFilesWrapper) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChunk(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationApplicationManifestQueryWithFilesWrapper) contextValidateChunk(ctx context.Context, formats strfmt.Registry) error {

	if m.Chunk != nil {
		if err := m.Chunk.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chunk")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chunk")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationApplicationManifestQueryWithFilesWrapper) contextValidateQuery(ctx context.Context, formats strfmt.Registry) error {

	if m.Query != nil {
		if err := m.Query.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("query")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationApplicationManifestQueryWithFilesWrapper) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationApplicationManifestQueryWithFilesWrapper) UnmarshalBinary(b []byte) error {
	var res ApplicationApplicationManifestQueryWithFilesWrapper
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
