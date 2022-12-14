// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha1ApplicationSetSpec ApplicationSetSpec represents a class of application set state.
//
// swagger:model v1alpha1ApplicationSetSpec
type V1alpha1ApplicationSetSpec struct {

	// generators
	Generators []*V1alpha1ApplicationSetGenerator `json:"generators"`

	// go template
	GoTemplate bool `json:"goTemplate,omitempty"`

	// sync policy
	SyncPolicy *V1alpha1ApplicationSetSyncPolicy `json:"syncPolicy,omitempty"`

	// template
	Template *V1alpha1ApplicationSetTemplate `json:"template,omitempty"`
}

// Validate validates this v1alpha1 application set spec
func (m *V1alpha1ApplicationSetSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGenerators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyncPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1ApplicationSetSpec) validateGenerators(formats strfmt.Registry) error {
	if swag.IsZero(m.Generators) { // not required
		return nil
	}

	for i := 0; i < len(m.Generators); i++ {
		if swag.IsZero(m.Generators[i]) { // not required
			continue
		}

		if m.Generators[i] != nil {
			if err := m.Generators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("generators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("generators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha1ApplicationSetSpec) validateSyncPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.SyncPolicy) { // not required
		return nil
	}

	if m.SyncPolicy != nil {
		if err := m.SyncPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("syncPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("syncPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1ApplicationSetSpec) validateTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.Template) { // not required
		return nil
	}

	if m.Template != nil {
		if err := m.Template.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1alpha1 application set spec based on the context it is used
func (m *V1alpha1ApplicationSetSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGenerators(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSyncPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1ApplicationSetSpec) contextValidateGenerators(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Generators); i++ {

		if m.Generators[i] != nil {
			if err := m.Generators[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("generators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("generators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha1ApplicationSetSpec) contextValidateSyncPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.SyncPolicy != nil {
		if err := m.SyncPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("syncPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("syncPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1ApplicationSetSpec) contextValidateTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.Template != nil {
		if err := m.Template.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1ApplicationSetSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1ApplicationSetSpec) UnmarshalBinary(b []byte) error {
	var res V1alpha1ApplicationSetSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
