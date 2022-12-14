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

// V1alpha1PullRequestGenerator PullRequestGenerator defines a generator that scrapes a PullRequest API to find candidate pull requests.
//
// swagger:model v1alpha1PullRequestGenerator
type V1alpha1PullRequestGenerator struct {

	// bitbucket server
	BitbucketServer *V1alpha1PullRequestGeneratorBitbucketServer `json:"bitbucketServer,omitempty"`

	// Filters for which pull requests should be considered.
	Filters []*V1alpha1PullRequestGeneratorFilter `json:"filters"`

	// gitea
	Gitea *V1alpha1PullRequestGeneratorGitea `json:"gitea,omitempty"`

	// github
	Github *V1alpha1PullRequestGeneratorGithub `json:"github,omitempty"`

	// gitlab
	Gitlab *V1alpha1PullRequestGeneratorGitLab `json:"gitlab,omitempty"`

	// Standard parameters.
	RequeueAfterSeconds string `json:"requeueAfterSeconds,omitempty"`

	// template
	Template *V1alpha1ApplicationSetTemplate `json:"template,omitempty"`
}

// Validate validates this v1alpha1 pull request generator
func (m *V1alpha1PullRequestGenerator) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBitbucketServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGitea(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGithub(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGitlab(formats); err != nil {
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

func (m *V1alpha1PullRequestGenerator) validateBitbucketServer(formats strfmt.Registry) error {
	if swag.IsZero(m.BitbucketServer) { // not required
		return nil
	}

	if m.BitbucketServer != nil {
		if err := m.BitbucketServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bitbucketServer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bitbucketServer")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1PullRequestGenerator) validateFilters(formats strfmt.Registry) error {
	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	for i := 0; i < len(m.Filters); i++ {
		if swag.IsZero(m.Filters[i]) { // not required
			continue
		}

		if m.Filters[i] != nil {
			if err := m.Filters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha1PullRequestGenerator) validateGitea(formats strfmt.Registry) error {
	if swag.IsZero(m.Gitea) { // not required
		return nil
	}

	if m.Gitea != nil {
		if err := m.Gitea.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gitea")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gitea")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1PullRequestGenerator) validateGithub(formats strfmt.Registry) error {
	if swag.IsZero(m.Github) { // not required
		return nil
	}

	if m.Github != nil {
		if err := m.Github.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("github")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("github")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1PullRequestGenerator) validateGitlab(formats strfmt.Registry) error {
	if swag.IsZero(m.Gitlab) { // not required
		return nil
	}

	if m.Gitlab != nil {
		if err := m.Gitlab.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gitlab")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gitlab")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1PullRequestGenerator) validateTemplate(formats strfmt.Registry) error {
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

// ContextValidate validate this v1alpha1 pull request generator based on the context it is used
func (m *V1alpha1PullRequestGenerator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBitbucketServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGitea(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGithub(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGitlab(ctx, formats); err != nil {
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

func (m *V1alpha1PullRequestGenerator) contextValidateBitbucketServer(ctx context.Context, formats strfmt.Registry) error {

	if m.BitbucketServer != nil {
		if err := m.BitbucketServer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bitbucketServer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bitbucketServer")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1PullRequestGenerator) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Filters); i++ {

		if m.Filters[i] != nil {
			if err := m.Filters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha1PullRequestGenerator) contextValidateGitea(ctx context.Context, formats strfmt.Registry) error {

	if m.Gitea != nil {
		if err := m.Gitea.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gitea")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gitea")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1PullRequestGenerator) contextValidateGithub(ctx context.Context, formats strfmt.Registry) error {

	if m.Github != nil {
		if err := m.Github.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("github")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("github")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1PullRequestGenerator) contextValidateGitlab(ctx context.Context, formats strfmt.Registry) error {

	if m.Gitlab != nil {
		if err := m.Gitlab.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gitlab")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gitlab")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1PullRequestGenerator) contextValidateTemplate(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V1alpha1PullRequestGenerator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1PullRequestGenerator) UnmarshalBinary(b []byte) error {
	var res V1alpha1PullRequestGenerator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
