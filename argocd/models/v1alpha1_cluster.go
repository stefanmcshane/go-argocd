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

// V1alpha1Cluster Cluster is the definition of a cluster resource
//
// swagger:model v1alpha1Cluster
type V1alpha1Cluster struct {

	// Annotations for cluster secret metadata
	Annotations map[string]string `json:"annotations,omitempty"`

	// Indicates if cluster level resources should be managed. This setting is used only if cluster is connected in a namespaced mode.
	ClusterResources bool `json:"clusterResources,omitempty"`

	// config
	Config *V1alpha1ClusterConfig `json:"config,omitempty"`

	// connection state
	ConnectionState *V1alpha1ConnectionState `json:"connectionState,omitempty"`

	// info
	Info *V1alpha1ClusterInfo `json:"info,omitempty"`

	// Labels for cluster secret metadata
	Labels map[string]string `json:"labels,omitempty"`

	// Name of the cluster. If omitted, will use the server address
	Name string `json:"name,omitempty"`

	// Holds list of namespaces which are accessible in that cluster. Cluster level resources will be ignored if namespace list is not empty.
	Namespaces []string `json:"namespaces"`

	// Reference between project and cluster that allow you automatically to be added as item inside Destinations project entity
	Project string `json:"project,omitempty"`

	// refresh requested at
	RefreshRequestedAt V1TimeString `json:"refreshRequestedAt,omitempty"`

	// Server is the API server URL of the Kubernetes cluster
	Server string `json:"server,omitempty"`

	// DEPRECATED: use Info.ServerVersion field instead.
	// The server version
	ServerVersion string `json:"serverVersion,omitempty"`

	// Shard contains optional shard number. Calculated on the fly by the application controller if not specified.
	Shard string `json:"shard,omitempty"`
}

// Validate validates this v1alpha1 cluster
func (m *V1alpha1Cluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefreshRequestedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1Cluster) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1Cluster) validateConnectionState(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectionState) { // not required
		return nil
	}

	if m.ConnectionState != nil {
		if err := m.ConnectionState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectionState")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1Cluster) validateInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.Info) { // not required
		return nil
	}

	if m.Info != nil {
		if err := m.Info.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("info")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1Cluster) validateRefreshRequestedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.RefreshRequestedAt) { // not required
		return nil
	}

	if err := m.RefreshRequestedAt.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("refreshRequestedAt")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("refreshRequestedAt")
		}
		return err
	}

	return nil
}

// ContextValidate validate this v1alpha1 cluster based on the context it is used
func (m *V1alpha1Cluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectionState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRefreshRequestedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1Cluster) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {
		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1Cluster) contextValidateConnectionState(ctx context.Context, formats strfmt.Registry) error {

	if m.ConnectionState != nil {
		if err := m.ConnectionState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectionState")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1Cluster) contextValidateInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.Info != nil {
		if err := m.Info.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("info")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1Cluster) contextValidateRefreshRequestedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RefreshRequestedAt.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("refreshRequestedAt")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("refreshRequestedAt")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1Cluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1Cluster) UnmarshalBinary(b []byte) error {
	var res V1alpha1Cluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
