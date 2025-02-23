// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EmsDestination ems destination
//
// swagger:model ems_destination
type EmsDestination struct {

	// links
	Links *EmsDestinationLinks `json:"_links,omitempty"`

	// certificate
	Certificate *EmsDestinationCertificate `json:"certificate,omitempty"`

	// Event destination
	// Example: administrator@mycompany.com
	Destination string `json:"destination,omitempty"`

	// filters
	Filters []*EmsDestinationFiltersItems0 `json:"filters,omitempty"`

	// Destination name.  Valid in POST.
	// Example: Admin_Email
	Name string `json:"name,omitempty"`

	// Flag indicating system-defined destinations.
	// Example: true
	// Read Only: true
	SystemDefined *bool `json:"system_defined,omitempty"`

	// Type of destination. Valid in POST.
	// Example: email
	// Enum: [snmp email syslog rest_api]
	Type string `json:"type,omitempty"`
}

// Validate validates this ems destination
func (m *EmsDestination) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsDestination) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *EmsDestination) validateCertificate(formats strfmt.Registry) error {
	if swag.IsZero(m.Certificate) { // not required
		return nil
	}

	if m.Certificate != nil {
		if err := m.Certificate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificate")
			}
			return err
		}
	}

	return nil
}

func (m *EmsDestination) validateFilters(formats strfmt.Registry) error {
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
				}
				return err
			}
		}

	}

	return nil
}

var emsDestinationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["snmp","email","syslog","rest_api"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		emsDestinationTypeTypePropEnum = append(emsDestinationTypeTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// ems_destination
	// EmsDestination
	// type
	// Type
	// snmp
	// END DEBUGGING
	// EmsDestinationTypeSnmp captures enum value "snmp"
	EmsDestinationTypeSnmp string = "snmp"

	// BEGIN DEBUGGING
	// ems_destination
	// EmsDestination
	// type
	// Type
	// email
	// END DEBUGGING
	// EmsDestinationTypeEmail captures enum value "email"
	EmsDestinationTypeEmail string = "email"

	// BEGIN DEBUGGING
	// ems_destination
	// EmsDestination
	// type
	// Type
	// syslog
	// END DEBUGGING
	// EmsDestinationTypeSyslog captures enum value "syslog"
	EmsDestinationTypeSyslog string = "syslog"

	// BEGIN DEBUGGING
	// ems_destination
	// EmsDestination
	// type
	// Type
	// rest_api
	// END DEBUGGING
	// EmsDestinationTypeRestAPI captures enum value "rest_api"
	EmsDestinationTypeRestAPI string = "rest_api"
)

// prop value enum
func (m *EmsDestination) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, emsDestinationTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EmsDestination) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ems destination based on the context it is used
func (m *EmsDestination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCertificate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSystemDefined(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsDestination) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *EmsDestination) contextValidateCertificate(ctx context.Context, formats strfmt.Registry) error {

	if m.Certificate != nil {
		if err := m.Certificate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificate")
			}
			return err
		}
	}

	return nil
}

func (m *EmsDestination) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Filters); i++ {

		if m.Filters[i] != nil {
			if err := m.Filters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmsDestination) contextValidateSystemDefined(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "system_defined", "body", m.SystemDefined); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsDestination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsDestination) UnmarshalBinary(b []byte) error {
	var res EmsDestination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsDestinationCertificate Certificate information is valid for the "rest_api" type.
//
// swagger:model EmsDestinationCertificate
type EmsDestinationCertificate struct {

	// Client certificate issuing CA
	// Example: VeriSign
	// Max Length: 256
	// Min Length: 1
	Ca string `json:"ca,omitempty"`

	// Client certificate serial number
	// Example: 1234567890
	// Max Length: 40
	// Min Length: 1
	SerialNumber string `json:"serial_number,omitempty"`
}

// Validate validates this ems destination certificate
func (m *EmsDestinationCertificate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCa(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerialNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsDestinationCertificate) validateCa(formats strfmt.Registry) error {
	if swag.IsZero(m.Ca) { // not required
		return nil
	}

	if err := validate.MinLength("certificate"+"."+"ca", "body", m.Ca, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("certificate"+"."+"ca", "body", m.Ca, 256); err != nil {
		return err
	}

	return nil
}

func (m *EmsDestinationCertificate) validateSerialNumber(formats strfmt.Registry) error {
	if swag.IsZero(m.SerialNumber) { // not required
		return nil
	}

	if err := validate.MinLength("certificate"+"."+"serial_number", "body", m.SerialNumber, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("certificate"+"."+"serial_number", "body", m.SerialNumber, 40); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this ems destination certificate based on context it is used
func (m *EmsDestinationCertificate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmsDestinationCertificate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsDestinationCertificate) UnmarshalBinary(b []byte) error {
	var res EmsDestinationCertificate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsDestinationFiltersItems0 ems destination filters items0
//
// swagger:model EmsDestinationFiltersItems0
type EmsDestinationFiltersItems0 struct {

	// links
	Links *EmsDestinationFiltersItems0Links `json:"_links,omitempty"`

	// name
	// Example: important-events
	Name string `json:"name,omitempty"`
}

// Validate validates this ems destination filters items0
func (m *EmsDestinationFiltersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsDestinationFiltersItems0) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ems destination filters items0 based on the context it is used
func (m *EmsDestinationFiltersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsDestinationFiltersItems0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsDestinationFiltersItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsDestinationFiltersItems0) UnmarshalBinary(b []byte) error {
	var res EmsDestinationFiltersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsDestinationFiltersItems0Links ems destination filters items0 links
//
// swagger:model EmsDestinationFiltersItems0Links
type EmsDestinationFiltersItems0Links struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ems destination filters items0 links
func (m *EmsDestinationFiltersItems0Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsDestinationFiltersItems0Links) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ems destination filters items0 links based on the context it is used
func (m *EmsDestinationFiltersItems0Links) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsDestinationFiltersItems0Links) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsDestinationFiltersItems0Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsDestinationFiltersItems0Links) UnmarshalBinary(b []byte) error {
	var res EmsDestinationFiltersItems0Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsDestinationLinks ems destination links
//
// swagger:model EmsDestinationLinks
type EmsDestinationLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ems destination links
func (m *EmsDestinationLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsDestinationLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ems destination links based on the context it is used
func (m *EmsDestinationLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsDestinationLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsDestinationLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsDestinationLinks) UnmarshalBinary(b []byte) error {
	var res EmsDestinationLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
