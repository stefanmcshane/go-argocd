// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha1GitFileGeneratorItem v1alpha1 git file generator item
//
// swagger:model v1alpha1GitFileGeneratorItem
type V1alpha1GitFileGeneratorItem struct {

	// path
	Path string `json:"path,omitempty"`
}

// Validate validates this v1alpha1 git file generator item
func (m *V1alpha1GitFileGeneratorItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1alpha1 git file generator item based on context it is used
func (m *V1alpha1GitFileGeneratorItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1GitFileGeneratorItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1GitFileGeneratorItem) UnmarshalBinary(b []byte) error {
	var res V1alpha1GitFileGeneratorItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
