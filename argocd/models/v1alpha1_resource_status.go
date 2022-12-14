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

// V1alpha1ResourceStatus ResourceStatus holds the current sync and health status of a resource
// TODO: describe members of this type
//
// swagger:model v1alpha1ResourceStatus
type V1alpha1ResourceStatus struct {

	// group
	Group string `json:"group,omitempty"`

	// health
	Health *V1alpha1HealthStatus `json:"health,omitempty"`

	// hook
	Hook bool `json:"hook,omitempty"`

	// kind
	Kind string `json:"kind,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// requires pruning
	RequiresPruning bool `json:"requiresPruning,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// sync wave
	SyncWave string `json:"syncWave,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this v1alpha1 resource status
func (m *V1alpha1ResourceStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHealth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1ResourceStatus) validateHealth(formats strfmt.Registry) error {
	if swag.IsZero(m.Health) { // not required
		return nil
	}

	if m.Health != nil {
		if err := m.Health.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1alpha1 resource status based on the context it is used
func (m *V1alpha1ResourceStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHealth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1ResourceStatus) contextValidateHealth(ctx context.Context, formats strfmt.Registry) error {

	if m.Health != nil {
		if err := m.Health.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1ResourceStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1ResourceStatus) UnmarshalBinary(b []byte) error {
	var res V1alpha1ResourceStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
