// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApplicationApplicationResourceResponse application application resource response
//
// swagger:model applicationApplicationResourceResponse
type ApplicationApplicationResourceResponse struct {

	// manifest
	Manifest string `json:"manifest,omitempty"`
}

// Validate validates this application application resource response
func (m *ApplicationApplicationResourceResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this application application resource response based on context it is used
func (m *ApplicationApplicationResourceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationApplicationResourceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationApplicationResourceResponse) UnmarshalBinary(b []byte) error {
	var res ApplicationApplicationResourceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
