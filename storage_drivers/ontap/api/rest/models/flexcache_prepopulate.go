// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FlexcachePrepopulate flexcache prepopulate
//
// swagger:model flexcache_prepopulate
type FlexcachePrepopulate struct {

	// dir paths
	DirPaths []string `json:"dir_paths,omitempty"`

	// exclude dir paths
	ExcludeDirPaths []string `json:"exclude_dir_paths,omitempty"`

	// Specifies whether or not the prepopulate action should search through the `dir_paths` recursively. If not set, the default value _true_ is used.
	Recurse *bool `json:"recurse,omitempty"`
}

// Validate validates this flexcache prepopulate
func (m *FlexcachePrepopulate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this flexcache prepopulate based on context it is used
func (m *FlexcachePrepopulate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FlexcachePrepopulate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlexcachePrepopulate) UnmarshalBinary(b []byte) error {
	var res FlexcachePrepopulate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
