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

// V1alpha1RevisionMetadata RevisionMetadata contains metadata for a specific revision in a Git repository
//
// swagger:model v1alpha1RevisionMetadata
type V1alpha1RevisionMetadata struct {

	// who authored this revision,
	// typically their name and email, e.g. "John Doe <john_doe@my-company.com>",
	// but might not match this example
	Author string `json:"author,omitempty"`

	// date
	Date *V1Time `json:"date,omitempty"`

	// Message contains the message associated with the revision, most likely the commit message.
	Message string `json:"message,omitempty"`

	// SignatureInfo contains a hint on the signer if the revision was signed with GPG, and signature verification is enabled.
	SignatureInfo string `json:"signatureInfo,omitempty"`

	// Tags specifies any tags currently attached to the revision
	// Floating tags can move from one revision to another
	Tags []string `json:"tags"`
}

// Validate validates this v1alpha1 revision metadata
func (m *V1alpha1RevisionMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1RevisionMetadata) validateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if m.Date != nil {
		if err := m.Date.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("date")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("date")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1alpha1 revision metadata based on the context it is used
func (m *V1alpha1RevisionMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1RevisionMetadata) contextValidateDate(ctx context.Context, formats strfmt.Registry) error {

	if m.Date != nil {
		if err := m.Date.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("date")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("date")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1RevisionMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1RevisionMetadata) UnmarshalBinary(b []byte) error {
	var res V1alpha1RevisionMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
