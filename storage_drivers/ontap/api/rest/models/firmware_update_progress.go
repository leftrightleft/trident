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
	"github.com/go-openapi/validate"
)

// FirmwareUpdateProgress firmware update progress
//
// swagger:model firmware_update_progress
type FirmwareUpdateProgress struct {

	// job
	Job *JobLink `json:"job,omitempty"`

	// update states
	UpdateStates []*FirmwareUpdateProgressState `json:"update_states,omitempty"`

	// zip file name
	// Example: disk_firmware.zip
	// Read Only: true
	ZipFileName string `json:"zip_file_name,omitempty"`
}

// Validate validates this firmware update progress
func (m *FirmwareUpdateProgress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJob(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateStates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FirmwareUpdateProgress) validateJob(formats strfmt.Registry) error {
	if swag.IsZero(m.Job) { // not required
		return nil
	}

	if m.Job != nil {
		if err := m.Job.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("job")
			}
			return err
		}
	}

	return nil
}

func (m *FirmwareUpdateProgress) validateUpdateStates(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateStates) { // not required
		return nil
	}

	for i := 0; i < len(m.UpdateStates); i++ {
		if swag.IsZero(m.UpdateStates[i]) { // not required
			continue
		}

		if m.UpdateStates[i] != nil {
			if err := m.UpdateStates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("update_states" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this firmware update progress based on the context it is used
func (m *FirmwareUpdateProgress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateJob(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdateStates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateZipFileName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FirmwareUpdateProgress) contextValidateJob(ctx context.Context, formats strfmt.Registry) error {

	if m.Job != nil {
		if err := m.Job.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("job")
			}
			return err
		}
	}

	return nil
}

func (m *FirmwareUpdateProgress) contextValidateUpdateStates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UpdateStates); i++ {

		if m.UpdateStates[i] != nil {
			if err := m.UpdateStates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("update_states" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FirmwareUpdateProgress) contextValidateZipFileName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "zip_file_name", "body", string(m.ZipFileName)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FirmwareUpdateProgress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FirmwareUpdateProgress) UnmarshalBinary(b []byte) error {
	var res FirmwareUpdateProgress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
