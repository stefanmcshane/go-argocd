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

// V1alpha1ApplicationSetCondition ApplicationSetCondition contains details about an applicationset condition, which is usally an error or warning
//
// swagger:model v1alpha1ApplicationSetCondition
type V1alpha1ApplicationSetCondition struct {

	// last transition time
	LastTransitionTime V1TimeString `json:"lastTransitionTime,omitempty"`

	// Message contains human-readable message indicating details about condition
	Message string `json:"message,omitempty"`

	// Single word camelcase representing the reason for the status eg ErrorOccurred
	Reason string `json:"reason,omitempty"`

	// True/False/Unknown
	Status string `json:"status,omitempty"`

	// Type is an applicationset condition type
	Type string `json:"type,omitempty"`
}

// Validate validates this v1alpha1 application set condition
func (m *V1alpha1ApplicationSetCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastTransitionTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1ApplicationSetCondition) validateLastTransitionTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastTransitionTime) { // not required
		return nil
	}

	if err := m.LastTransitionTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastTransitionTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastTransitionTime")
		}
		return err
	}

	return nil
}

// ContextValidate validate this v1alpha1 application set condition based on the context it is used
func (m *V1alpha1ApplicationSetCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastTransitionTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1ApplicationSetCondition) contextValidateLastTransitionTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LastTransitionTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastTransitionTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastTransitionTime")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1ApplicationSetCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1ApplicationSetCondition) UnmarshalBinary(b []byte) error {
	var res V1alpha1ApplicationSetCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
